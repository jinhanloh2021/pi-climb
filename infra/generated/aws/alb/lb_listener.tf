resource "aws_lb_listener" "tfer--arn-003A-aws-003A-elasticloadbalancing-003A-ap-southeast-1-003A-842832773369-003A-listener-002F-app-002F-pi-climb-alb-002F-e17fb6d75ef46fbc-002F-56457fe2a9b0fc67" {
  certificate_arn = "arn:aws:acm:ap-southeast-1:842832773369:certificate/488ad81e-3c84-4ee0-9eb5-23681cb270eb"

  default_action {
    forward {
      stickiness {
        duration = "3600"
        enabled  = "false"
      }

      target_group {
        arn    = "arn:aws:elasticloadbalancing:ap-southeast-1:842832773369:targetgroup/pi-climb-nextjs-tg/f3d117d78cd69bce"
        weight = "1"
      }
    }

    order            = "1"
    target_group_arn = "arn:aws:elasticloadbalancing:ap-southeast-1:842832773369:targetgroup/pi-climb-nextjs-tg/f3d117d78cd69bce"
    type             = "forward"
  }

  load_balancer_arn = aws_lb.tfer--pi-climb-alb.id

  port                                 = "443"
  protocol                             = "HTTPS"
  region                               = "ap-southeast-1"
  routing_http_response_server_enabled = "true"
  ssl_policy                           = "ELBSecurityPolicy-TLS13-1-2-Res-2021-06"
}

resource "aws_lb_listener" "tfer--arn-003A-aws-003A-elasticloadbalancing-003A-ap-southeast-1-003A-842832773369-003A-listener-002F-app-002F-pi-climb-alb-002F-e17fb6d75ef46fbc-002F-e73dee8ecf4b8bfa" {
  default_action {
    order = "1"

    redirect {
      host        = "#{host}"
      path        = "/#{path}"
      port        = "443"
      protocol    = "HTTPS"
      query       = "#{query}"
      status_code = "HTTP_301"
    }

    type = "redirect"
  }

  load_balancer_arn                    = aws_lb.tfer--pi-climb-alb.id
  port                                 = "80"
  protocol                             = "HTTP"
  region                               = "ap-southeast-1"
  routing_http_response_server_enabled = "true"
}
