resource "aws_eip" "tfer--eipalloc-02ef648c1ac2176f5" {
  domain               = "vpc"
  network_border_group = "ap-southeast-1"
  public_ipv4_pool     = "amazon"
  region               = "ap-southeast-1"
}
