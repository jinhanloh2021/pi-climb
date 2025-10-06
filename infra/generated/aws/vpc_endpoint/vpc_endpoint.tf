resource "aws_vpc_endpoint" "tfer--vpce-04fb7065534a1513f" {
  dns_options {
    dns_record_ip_type                             = "ipv4"
    private_dns_only_for_inbound_resolver_endpoint = "false"
  }

  ip_address_type     = "ipv4"
  policy              = "{\"Statement\":[{\"Action\":\"*\",\"Effect\":\"Allow\",\"Principal\":\"*\",\"Resource\":\"*\"}]}"
  private_dns_enabled = "true"
  region              = "ap-southeast-1"
  security_group_ids  = [var.pi-climb_endpoint_sg]
  service_name        = "com.amazonaws.ap-southeast-1.ecr.dkr"
  service_region      = "ap-southeast-1"

  subnet_configuration {
    ipv4      = "10.0.0.136"
    subnet_id = "subnet-04b6d206e946cb507"
  }

  subnet_ids = ["subnet-04b6d206e946cb507"]

  tags = {
    Name = "ecr-dkr-endpoint"
  }

  tags_all = {
    Name = "ecr-dkr-endpoint"
  }

  vpc_endpoint_type = "Interface"
  vpc_id            = "vpc-0ee6b184da0ba05bc"
}

resource "aws_vpc_endpoint" "tfer--vpce-08f7ba2e8ff3282c6" {
  policy              = "{\"Statement\":[{\"Action\":\"*\",\"Effect\":\"Allow\",\"Principal\":\"*\",\"Resource\":\"*\"}],\"Version\":\"2008-10-17\"}"
  private_dns_enabled = "false"
  region              = "ap-southeast-1"
  route_table_ids     = ["rtb-09b4d703f76e15d12"]
  service_name        = "com.amazonaws.ap-southeast-1.s3"
  service_region      = "ap-southeast-1"

  tags = {
    Description = "Endpoint for private subnet"
    Name        = "pi-climb-vpce-s3"
  }

  tags_all = {
    Description = "Endpoint for private subnet"
    Name        = "pi-climb-vpce-s3"
  }

  vpc_endpoint_type = "Gateway"
  vpc_id            = "vpc-0ee6b184da0ba05bc"
}

resource "aws_vpc_endpoint" "tfer--vpce-0ccce2a48c895cd59" {
  dns_options {
    dns_record_ip_type                             = "ipv4"
    private_dns_only_for_inbound_resolver_endpoint = "false"
  }

  ip_address_type     = "ipv4"
  policy              = "{\"Statement\":[{\"Action\":\"*\",\"Effect\":\"Allow\",\"Principal\":\"*\",\"Resource\":\"*\"}]}"
  private_dns_enabled = "true"
  region              = "ap-southeast-1"
  security_group_ids  = [var.pi-climb_endpoint_sg]
  service_name        = "com.amazonaws.ap-southeast-1.logs"
  service_region      = "ap-southeast-1"

  subnet_configuration {
    ipv4      = "10.0.0.139"
    subnet_id = "subnet-04b6d206e946cb507"
  }

  subnet_ids = ["subnet-04b6d206e946cb507"]

  tags = {
    Name = "logs-endpoint"
  }

  tags_all = {
    Name = "logs-endpoint"
  }

  vpc_endpoint_type = "Interface"
  vpc_id            = "vpc-0ee6b184da0ba05bc"
}

resource "aws_vpc_endpoint" "tfer--vpce-0cd97c188366db4d4" {
  dns_options {
    dns_record_ip_type                             = "ipv4"
    private_dns_only_for_inbound_resolver_endpoint = "false"
  }

  ip_address_type     = "ipv4"
  policy              = "{\"Statement\":[{\"Action\":\"*\",\"Effect\":\"Allow\",\"Principal\":\"*\",\"Resource\":\"*\"}]}"
  private_dns_enabled = "true"
  region              = "ap-southeast-1"
  security_group_ids  = ["sg-01e41717258f067bc"]
  service_name        = "com.amazonaws.ap-southeast-1.secretsmanager"
  service_region      = "ap-southeast-1"

  subnet_configuration {
    ipv4      = "10.0.0.132"
    subnet_id = "subnet-04b6d206e946cb507"
  }

  subnet_ids = ["subnet-04b6d206e946cb507"]

  tags = {
    Name = "secrets-manager-endpoint"
  }

  tags_all = {
    Name = "secrets-manager-endpoint"
  }

  vpc_endpoint_type = "Interface"
  vpc_id            = "vpc-0ee6b184da0ba05bc"
}

resource "aws_vpc_endpoint" "tfer--vpce-0f76d4657a64f2530" {
  dns_options {
    dns_record_ip_type                             = "ipv4"
    private_dns_only_for_inbound_resolver_endpoint = "false"
  }

  ip_address_type     = "ipv4"
  policy              = "{\"Statement\":[{\"Action\":\"*\",\"Effect\":\"Allow\",\"Principal\":\"*\",\"Resource\":\"*\"}]}"
  private_dns_enabled = "true"
  region              = "ap-southeast-1"
  security_group_ids  = ["sg-01e41717258f067bc"]
  service_name        = "com.amazonaws.ap-southeast-1.ecr.api"
  service_region      = "ap-southeast-1"

  subnet_configuration {
    ipv4      = "10.0.0.138"
    subnet_id = "subnet-04b6d206e946cb507"
  }

  subnet_ids = ["subnet-04b6d206e946cb507"]

  tags = {
    Name = "ecr-api-endpoint"
  }

  tags_all = {
    Name = "ecr-api-endpoint"
  }

  vpc_endpoint_type = "Interface"
  vpc_id            = "vpc-0ee6b184da0ba05bc"
}
