resource "aws_ecs_cluster" "tfer--pi-climb-bzz9ub" {
  configuration {
    execute_command_configuration {
      logging = "DEFAULT"
    }
  }

  name   = "pi-climb-bzz9ub"
  region = "ap-southeast-1"

  setting {
    name  = "containerInsights"
    value = "disabled"
  }
}
