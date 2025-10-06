resource "aws_nat_gateway" "tfer--nat-0c87118685dfc7982" {
  allocation_id     = "eipalloc-02ef648c1ac2176f5"
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
