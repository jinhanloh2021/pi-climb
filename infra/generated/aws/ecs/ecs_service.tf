resource "aws_ecs_service" "tfer--pi-climb-bzz9ub_pi-climb_service-dtvx78q8" {
  availability_zone_rebalancing = "ENABLED"

  capacity_provider_strategy {
    base              = "0"
    capacity_provider = "FARGATE"
    weight            = "1"
  }

  cluster = aws_ecs_cluster.tfer--pi-climb-bzz9ub.arn

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
  health_check_grace_period_seconds  = "60"

  load_balancer {
    container_name   = "go_server"
    container_port   = "8080"
    target_group_arn = var.aws_lb_target_group_tfer--pi-climb-go-tg_arn
  }

  load_balancer {
    container_name   = "nextjs_server"
    container_port   = "3000"
    target_group_arn = var.aws_lb_target_group_tfer--pi-climb-nextjs-tg_arn
  }

  name = "pi-climb_service-dtvx78q8"

  network_configuration {
    assign_public_ip = "false"
    security_groups  = ["${var.pi-climb-service_security_group_id}"]
    subnets          = ["subnet-04b6d206e946cb507"] # subnet hardcoded
  }

  platform_version      = "LATEST"
  region                = "ap-southeast-1"
  scheduling_strategy   = "REPLICA"
  wait_for_steady_state = false
  task_definition       = "${aws_ecs_task_definition.tfer--task-definition-002F-pi-climb_task.id}:${aws_ecs_task_definition.tfer--task-definition-002F-pi-climb_task.revision}" # // pi-climb_task:16
  tags                  = {}
  tags_all              = {}
}
