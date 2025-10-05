resource "aws_cloudwatch_log_group" "tfer---002F-ecs-002F-pi-climb_task" {
  log_group_class   = "STANDARD"
  name              = "/ecs/pi-climb_task"
  region            = "ap-southeast-1"
  retention_in_days = "5"
  skip_destroy      = "false"
}
