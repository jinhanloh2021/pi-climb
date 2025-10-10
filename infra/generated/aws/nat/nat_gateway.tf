resource "aws_nat_gateway" "tfer--nat-0c87118685dfc7982" {
  allocation_id     = var.aws_eip_tfer--eipalloc_allocation_id
  connectivity_type = "public"
  private_ip        = "10.0.0.10"
  region            = "ap-southeast-1"
  subnet_id         = "subnet-01e07076b33511ee4"

  tags = {
    Name = "pi-climb-public-nat"
  }

  tags_all = {
    Name = "pi-climb-public-nat"
  }
}

data "aws_route_table" "rtb-09b4d703f76e15d12" {
  route_table_id = "rtb-09b4d703f76e15d12"
}

resource "aws_route" "route" {
  route_table_id         = data.aws_route_table.rtb-09b4d703f76e15d12.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.tfer--nat-0c87118685dfc7982.id
  depends_on             = [aws_nat_gateway.tfer--nat-0c87118685dfc7982]
}
