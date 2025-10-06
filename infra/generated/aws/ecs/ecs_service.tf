resource "aws_ecs_service" "tfer--pi-climb-bzz9ub_pi-climb_service-dtvx78q8" {
  availability_zone_rebalancing = "ENABLED"

  capacity_provider_strategy {
    base              = "0"
    capacity_provider = "FARGATE"
    weight            = "1"
  }

  cluster = "pi-climb-bzz9ub"

  deployment_circuit_breaker {
    enable   = "true"
    rollback = "true"
  }

  deployment_configuration {
    bake_time_in_minutes = "0"
    strategy             = "ROLLING"
  }

  deployment_controller {
    type = "ECS"
  }

  deployment_maximum_percent         = "200"
  deployment_minimum_healthy_percent = "100"
  desired_count                      = "1"
  enable_ecs_managed_tags            = "true"
  enable_execute_command             = "false"
  health_check_grace_period_seconds  = "0"

  load_balancer {
    container_name   = "go_server"
    container_port   = "8080"
    target_group_arn = "arn:aws:elasticloadbalancing:ap-southeast-1:842832773369:targetgroup/pi-climb-go-tg/0c97f04d14eb3e7f"
  }

  load_balancer {
    container_name   = "nextjs_server"
    container_port   = "3000"
    target_group_arn = "arn:aws:elasticloadbalancing:ap-southeast-1:842832773369:targetgroup/pi-climb-nextjs-tg/f3d117d78cd69bce"
  }

  name = "pi-climb_service-dtvx78q8"

  network_configuration {
    assign_public_ip = "false"
    security_groups  = ["${var.pi-climb-service_security_group_id}"]
    subnets          = ["subnet-04b6d206e946cb507"]
  }

  platform_version    = "LATEST"
  region              = "ap-southeast-1"
  scheduling_strategy = "REPLICA"
  task_definition     = "arn:aws:ecs:ap-southeast-1:842832773369:task-definition/pi-climb_task:16"
}
