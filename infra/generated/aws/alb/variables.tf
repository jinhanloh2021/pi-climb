# data "terraform_remote_state" "alb" {
#   backend = "local"
#
#   config = {
#     path = "../../../generated/aws/alb/terraform.tfstate"
#   }
# }

# Originally for security group ID
# data "terraform_remote_state" "sg" {
#   backend = "local"
#
#   config = {
#     path = "../../../generated/aws/sg/terraform.tfstate"
#   }
# }

variable "alb_security_group_id" {
  description = "Security group ID for load balancer"
  type        = string
  default     = ""
}

