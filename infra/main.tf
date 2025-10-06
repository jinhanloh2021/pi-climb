// todo: Terraform import to generate state files. S3 storage
// Use SSM to get image tag of latest image for terraform
// Inline task-definition.json. Find online example on github, copy
// export AWS_PROFILE="AdministratorAccess-842832773369"
// terraformer import aws --resources=alb,ecs,eip,logs,nat,sg,vpc_endpoint --connect=true --regions=ap-southeast-1 --profile=AdministratorAccess-842832773369

terraform {
  required_providers {
    aws = {
      version = "~> 6.15.0"
    }
  }
}

module "alb" {
  source                = "./generated/aws/alb/"
  alb_security_group_id = module.sg.aws_security_group_tfer--pi-climb-lb-sg_sg-02821e5f312595268_id
}

module "nat" {
  source                               = "./generated/aws/nat/"
  aws_eip_tfer--eipalloc_allocation_id = module.eip.aws_eip_tfer--eipalloc_allocation_id
}

module "ecs" {
  source                             = "./generated/aws/ecs/"
  pi-climb-service_security_group_id = module.sg.aws_security_group_tfer--pi-climb-service-sg_sg-0802e5a19b6cc4611_id
  // task definition latest image tag
  aws_lb_target_group_tfer--pi-climb-nextjs-tg_arn = module.alb.aws_lb_target_group_tfer--pi-climb-nextjs-tg_arn
  aws_lb_target_group_tfer--pi-climb-go-tg_arn     = module.alb.aws_lb_target_group_tfer--pi-climb-go-tg_arn
}

module "eip" {
  source = "./generated/aws/eip/"
}

module "logs" {
  source = "./generated/aws/logs/"
}


module "sg" {
  source = "./generated/aws/sg/"
}

module "vpc_endpoint" {
  source               = "./generated/aws/vpc_endpoint"
  pi-climb_endpoint_sg = module.sg.aws_security_group_tfer--pi-climb-endpoint-sg_sg-01e41717258f067bc_id
}

