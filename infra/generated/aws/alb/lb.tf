resource "aws_lb" "tfer--pi-climb-alb" {
  client_keep_alive = "3600"

  connection_logs {
    bucket  = ""
    enabled = "false"
  }

  desync_mitigation_mode                      = "defensive"
  drop_invalid_header_fields                  = "false"
  enable_cross_zone_load_balancing            = "true"
  enable_deletion_protection                  = "false"
  enable_http2                                = "true"
  enable_tls_version_and_cipher_suite_headers = "false"
  enable_waf_fail_open                        = "false"
  enable_xff_client_port                      = "false"
  enable_zonal_shift                          = "false"
  idle_timeout                                = "60"
  internal                                    = "false"
  ip_address_type                             = "ipv4"
  load_balancer_type                          = "application"
  name                                        = "pi-climb-alb"
  preserve_host_header                        = "false"
  region                                      = "ap-southeast-1"
  security_groups                             = ["${var.alb_security_group_id}"]

  subnets                    = ["subnet-01e07076b33511ee4", "subnet-05f09df36061ff98f"]
  xff_header_processing_mode = "append"
}
