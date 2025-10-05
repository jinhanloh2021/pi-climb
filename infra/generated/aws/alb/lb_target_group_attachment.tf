resource "aws_lb_target_group_attachment" "tfer--arn-003A-aws-003A-elasticloadbalancing-003A-ap-southeast-1-003A-842832773369-003A-targetgroup-002F-pi-climb-go-tg-002F-0c97f04d14eb3e7f-10-002E-0-002E-0-002E-134" {
  region           = "ap-southeast-1"
  target_group_arn = "arn:aws:elasticloadbalancing:ap-southeast-1:842832773369:targetgroup/pi-climb-go-tg/0c97f04d14eb3e7f"
  target_id        = "10.0.0.134"
}

resource "aws_lb_target_group_attachment" "tfer--arn-003A-aws-003A-elasticloadbalancing-003A-ap-southeast-1-003A-842832773369-003A-targetgroup-002F-pi-climb-nextjs-tg-002F-f3d117d78cd69bce-10-002E-0-002E-0-002E-134" {
  region           = "ap-southeast-1"
  target_group_arn = "arn:aws:elasticloadbalancing:ap-southeast-1:842832773369:targetgroup/pi-climb-nextjs-tg/f3d117d78cd69bce"
  target_id        = "10.0.0.134"
}
