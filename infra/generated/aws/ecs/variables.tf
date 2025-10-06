variable "pi-climb-service_security_group_id" {
  description = "Security group ID for pi-climb-service"
  type        = string
}

variable "aws_lb_target_group_tfer--pi-climb-nextjs-tg_arn" {
  description = "Target group arn for nextjs container"
  type        = string
}

variable "aws_lb_target_group_tfer--pi-climb-go-tg_arn" {
  description = "Target group arn for Go container"
  type        = string
}
