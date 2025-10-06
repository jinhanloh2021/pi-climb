# For importing sg for pi-climb-service
# data "terraform_remote_state" "sg" {
#   backend = "local"
#
#   config = {
#     path = "../../../generated/aws/sg/terraform.tfstate"
#   }
# }

variable "pi-climb-service_security_group_id" {
  description = "Security group ID for pi-climb-service"
  type        = string
  default     = ""
}

