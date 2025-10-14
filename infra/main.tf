// export AWS_PROFILE="AdministratorAccess-842832773369"
// terraformer import aws --resources=alb,ecs,eip,logs,nat,sg,vpc_endpoint --connect=true --regions=ap-southeast-1 --profile=AdministratorAccess-842832773369
// possible stale pointer from private-subnet to NAT // Confirmed

terraform {
  required_providers {
    aws = {
      version = "~> 6.15.0"
    }
  }
  backend "s3" {
    bucket       = "pi-climb-tf-state-bucket"
    key          = "tf_state"
    region       = "ap-southeast-1"
    use_lockfile = true
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
  source               = "./generated/aws/vpc_endpoint/"
  pi-climb_endpoint_sg = module.sg.aws_security_group_tfer--pi-climb-endpoint-sg_sg-01e41717258f067bc_id
}

module "route53" {
  source                             = "./generated/aws/route53/"
  aws_lb_tfer--pi-climb-alb_dns_name = module.alb.aws_lb_tfer--pi-climb-alb_dns_name
  aws_lb_tfer--pi-climb-alb_zone_id  = module.alb.aws_lb_tfer--pi-climb-alb_zone_id
}

