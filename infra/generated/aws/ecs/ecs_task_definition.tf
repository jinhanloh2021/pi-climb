resource "aws_ecs_task_definition" "tfer--task-definition-002F-pi-climb_task" {
  container_definitions    = "[{\"environment\":[],\"essential\":true,\"image\":\"842832773369.dkr.ecr.ap-southeast-1.amazonaws.com/pi-climb-dev-server:d3584cba4f52d6da9f453109bdd05bc4e5f3429d\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-stream-prefix\":\"ecs\",\"awslogs-group\":\"/ecs/pi-climb_task\",\"awslogs-create-group\":\"true\",\"awslogs-region\":\"ap-southeast-1\"}},\"mountPoints\":[],\"name\":\"go_server\",\"portMappings\":[{\"appProtocol\":\"http\",\"containerPort\":8080,\"hostPort\":8080,\"name\":\"go_server-8080-tcp\",\"protocol\":\"tcp\"}],\"secrets\":[{\"name\":\"DATABASE_URL\",\"valueFrom\":\"arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:DATABASE_URL::\"},{\"name\":\"JWT_SECRET\",\"valueFrom\":\"arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:JWT_SECRET::\"},{\"name\":\"SUPABASE_AUTH_EXTERNAL_GOOGLE_CLIENT_ID\",\"valueFrom\":\"arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:SUPABASE_AUTH_EXTERNAL_GOOGLE_CLIENT_ID::\"},{\"name\":\"SUPABASE_AUTH_EXTERNAL_GOOGLE_SECRET\",\"valueFrom\":\"arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:SUPABASE_AUTH_EXTERNAL_GOOGLE_SECRET::\"}],\"systemControls\":[],\"volumesFrom\":[]},{\"environment\":[],\"essential\":true,\"image\":\"842832773369.dkr.ecr.ap-southeast-1.amazonaws.com/pi-climb-dev-web:d3584cba4f52d6da9f453109bdd05bc4e5f3429d\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-region\":\"ap-southeast-1\",\"awslogs-stream-prefix\":\"ecs\",\"awslogs-group\":\"/ecs/pi-climb_task\",\"awslogs-create-group\":\"true\"}},\"mountPoints\":[],\"name\":\"nextjs_server\",\"portMappings\":[{\"appProtocol\":\"http\",\"containerPort\":3000,\"hostPort\":3000,\"name\":\"nextjs_server-3000-tcp\",\"protocol\":\"tcp\"}],\"secrets\":[{\"name\":\"NEXT_PUBLIC_API_URL\",\"valueFrom\":\"arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:NEXT_PUBLIC_API_URL::\"},{\"name\":\"NEXT_PUBLIC_SUPABASE_ANON_KEY\",\"valueFrom\":\"arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:NEXT_PUBLIC_SUPABASE_ANON_KEY::\"},{\"name\":\"NEXT_PUBLIC_SUPABASE_URL\",\"valueFrom\":\"arn:aws:secretsmanager:ap-southeast-1:842832773369:secret:pi-climb/dev-wXQQeQ:NEXT_PUBLIC_SUPABASE_URL::\"}],\"systemControls\":[],\"volumesFrom\":[]}]"
  cpu                      = "256"
  enable_fault_injection   = "false"
  execution_role_arn       = "arn:aws:iam::842832773369:role/ecsTaskExecutionRole"
  family                   = "pi-climb_task"
  memory                   = "1024"
  network_mode             = "awsvpc"
  region                   = "ap-southeast-1"
  requires_compatibilities = ["FARGATE"]

  runtime_platform {
    cpu_architecture        = "X86_64"
    operating_system_family = "LINUX"
  }

  track_latest = "false"
}
