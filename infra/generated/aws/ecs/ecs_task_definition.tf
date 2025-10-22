data "aws_ssm_parameter" "server_image_uri" {
  name = "/pi-climb/dev/server-image-uri"
}

data "aws_ssm_parameter" "web_image_uri" {
  name = "/pi-climb/dev/web-image-uri"
}

resource "aws_ecs_task_definition" "tfer--task-definition-002F-pi-climb_task" {
  family                   = "pi-climb_task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "1024"
  execution_role_arn       = "arn:aws:iam::842832773369:role/ecsTaskExecutionRole"

  container_definitions = jsonencode([
    {
      name      = "go_server"
      essential = true
      image     = data.aws_ssm_parameter.server_image_uri.value
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
          protocol      = "tcp"
          appProtocol   = "http"
          name          = "go_server-8080-tcp"
        }
      ]
      healthCheck = {
        "command" : [
          "CMD-SHELL",
          "curl -f http://localhost:8080/api/health || exit 1"
        ],
        "interval" : 30,
        "timeout" : 5,
        "retries" : 4,
        "startPeriod" : 90
      }
      secrets = [
        { name = "DATABASE_URL", valueFrom = "arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:DATABASE_URL::" },
        { name = "JWT_SECRET", valueFrom = "arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:JWT_SECRET::" },
        { name = "SUPABASE_AUTH_EXTERNAL_GOOGLE_CLIENT_ID", valueFrom = "arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:SUPABASE_AUTH_EXTERNAL_GOOGLE_CLIENT_ID::" },
        { name = "SUPABASE_AUTH_EXTERNAL_GOOGLE_SECRET", valueFrom = "arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:SUPABASE_AUTH_EXTERNAL_GOOGLE_SECRET::" }
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-group"         = "/ecs/pi-climb_task"
          "awslogs-region"        = "ap-southeast-1"
          "awslogs-stream-prefix" = "ecs"
          "awslogs-create-group"  = "true"
        }
      }
    },
    {
      name      = "nextjs_server"
      essential = true
      image     = data.aws_ssm_parameter.web_image_uri.value
      portMappings = [
        {
          containerPort = 3000
          hostPort      = 3000
          protocol      = "tcp"
          appProtocol   = "http"
          name          = "nextjs_server-3000-tcp"
        }
      ]
      healthCheck = {
        "command" : [
          "CMD-SHELL",
          "curl -f http://localhost:3000/health || exit 1"
        ],
        "interval" : 30,
        "timeout" : 5,
        "retries" : 4,
        "startPeriod" : 90
      }
      secrets = [
        { name = "NEXT_PUBLIC_API_URL", valueFrom = "arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:NEXT_PUBLIC_API_URL::" },
        { name = "NEXT_PUBLIC_SUPABASE_ANON_KEY", valueFrom = "arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:NEXT_PUBLIC_SUPABASE_ANON_KEY::" },
        { name = "NEXT_PUBLIC_SUPABASE_URL", valueFrom = "arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:NEXT_PUBLIC_SUPABASE_URL::" }
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-group"         = "/ecs/pi-climb_task"
          "awslogs-region"        = "ap-southeast-1"
          "awslogs-stream-prefix" = "ecs"
          "awslogs-create-group"  = "true"
        }
      }
    }
  ])

  runtime_platform {
    cpu_architecture        = "X86_64"
    operating_system_family = "LINUX"
  }
}
