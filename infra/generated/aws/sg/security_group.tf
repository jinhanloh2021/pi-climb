resource "aws_security_group" "tfer--default_sg-038e78107566ef6e2" {
  description = "default VPC security group"

  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = "0"
    protocol    = "-1"
    self        = "false"
    to_port     = "0"
  }

  ingress {
    from_port = "0"
    protocol  = "-1"
    self      = "true"
    to_port   = "0"
  }

  name   = "default"
  region = "ap-southeast-1"
  vpc_id = "vpc-0ee6b184da0ba05bc"
}

resource "aws_security_group" "tfer--default_sg-06785d05b75b329ef" {
  description = "default VPC security group"

  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = "0"
    protocol    = "-1"
    self        = "false"
    to_port     = "0"
  }

  ingress {
    from_port = "0"
    protocol  = "-1"
    self      = "true"
    to_port   = "0"
  }

  name   = "default"
  region = "ap-southeast-1"
  vpc_id = "vpc-09a8b05adfbc748a0"
}

resource "aws_security_group" "tfer--pi-climb-endpoint-sg_sg-01e41717258f067bc" {
  description = "Allows traffic from ECS Fargate tasks to reach endpoints over HTTPS"

  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = "0"
    protocol    = "-1"
    self        = "false"
    to_port     = "0"
  }

  ingress {
    description     = "Allows traffic from private pi-climb-service to AWS endpoints"
    from_port       = "443"
    protocol        = "tcp"
    security_groups = ["${data.terraform_remote_state.sg.outputs.aws_security_group_tfer--pi-climb-service-sg_sg-0802e5a19b6cc4611_id}"]
    self            = "false"
    to_port         = "443"
  }

  name   = "pi-climb-endpoint-sg"
  region = "ap-southeast-1"
  vpc_id = "vpc-0ee6b184da0ba05bc"
}

resource "aws_security_group" "tfer--pi-climb-lb-sg_sg-02821e5f312595268" {
  description = "Allows public HTTP and HTTPS traffic"

  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = "0"
    protocol    = "-1"
    self        = "false"
    to_port     = "0"
  }

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    description = "Internet"
    from_port   = "443"
    protocol    = "tcp"
    self        = "false"
    to_port     = "443"
  }

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    description = "Internet"
    from_port   = "80"
    protocol    = "tcp"
    self        = "false"
    to_port     = "80"
  }

  name   = "pi-climb-lb-sg"
  region = "ap-southeast-1"
  vpc_id = "vpc-0ee6b184da0ba05bc"
}

resource "aws_security_group" "tfer--pi-climb-service-sg_sg-0802e5a19b6cc4611" {
  description = "Allows traffic only from the load balancer sg"

  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = "0"
    protocol    = "-1"
    self        = "false"
    to_port     = "0"
  }

  ingress {
    description     = "Allow only from load balancer"
    from_port       = "3000"
    protocol        = "tcp"
    security_groups = ["${data.terraform_remote_state.sg.outputs.aws_security_group_tfer--pi-climb-lb-sg_sg-02821e5f312595268_id}"]
    self            = "false"
    to_port         = "3000"
  }

  name   = "pi-climb-service-sg"
  region = "ap-southeast-1"
  vpc_id = "vpc-0ee6b184da0ba05bc"
}
