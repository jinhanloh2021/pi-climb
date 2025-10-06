resource "aws_lb_listener_rule" "tfer--arn-003A-aws-003A-elasticloadbalancing-003A-ap-southeast-1-003A-842832773369-003A-listener-rule-002F-app-002F-pi-climb-alb-002F-e17fb6d75ef46fbc-002F-56457fe2a9b0fc67-002F-6129aabee2a113aa" {
  action {
    forward {
      stickiness {
        duration = "3600"
        enabled  = "false"
      }

      target_group {
        arn    = aws_lb_target_group.tfer--pi-climb-go-tg.arn
        weight = "1"
      }
    }

    order = "1"
    type  = "forward"
  }

  condition {
    path_pattern {
      values = ["/api/*"]
    }
  }

  listener_arn = aws_lb_listener.tfer--arn-003A-aws-003A-elasticloadbalancing-003A-ap-southeast-1-003A-842832773369-003A-listener-002F-app-002F-pi-climb-alb-002F-e17fb6d75ef46fbc-002F-56457fe2a9b0fc67.arn
  priority     = "1"
  region       = "ap-southeast-1"

  tags = {
    Name = "api_forward_go_server"
  }

  tags_all = {
    Name = "api_forward_go_server"
  }
}
