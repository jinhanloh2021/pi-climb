# Read only existing infra
data "aws_route53_zone" "piclimb_com" {
  name         = "piclimb.com."
  private_zone = false
}

resource "aws_route53_record" "dev_piclimb_com" {
  name    = "dev.piclimb.com"
  zone_id = data.aws_route53_zone.piclimb_com.zone_id 
  type    = "A"
  
  alias {
    name                   = "dualstack.${var.aws_lb_tfer--pi-climb-alb_dns_name}"
    zone_id                = var.aws_lb_tfer--pi-climb-alb_zone_id
    evaluate_target_health = true
  }
}

resource "aws_route53_record" "www_dev_piclimb_com" {
  name    = "www.dev.piclimb.com"
  zone_id = data.aws_route53_zone.piclimb_com.zone_id 
  type    = "A"
  
  alias {
    name                   = "dualstack.${var.aws_lb_tfer--pi-climb-alb_dns_name}"
    zone_id                = var.aws_lb_tfer--pi-climb-alb_zone_id
    evaluate_target_health = true
  }
}

# Manage resource for prod www.piclimb.com
