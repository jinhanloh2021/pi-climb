// todo: Module orchestration with terraform
// todo: Terraform import to generate state files. S3 storage
// Use SSM to get image tag of latest image for terraform
// Inline task-definition.json. Find online example on github, copy

provider "aws" {
  region = "ap-southeast-1"
}

terraform {
  required_providers {
    aws = {
      version = "~> 6.15.0"
    }
  }
}
