resource "aws_lb" "tfer--pi-climb-alb" {
  client_keep_alive = "3600"

  connection_logs {
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
  security_groups                             = ["${data.terraform_remote_state.sg.outputs.aws_security_group_tfer--pi-climb-lb-sg_sg-02821e5f312595268_id}"]

  subnet_mapping {
    subnet_id = "subnet-01e07076b33511ee4"
  }

  subnet_mapping {
    subnet_id = "subnet-05f09df36061ff98f"
  }

  subnets                    = ["subnet-01e07076b33511ee4", "subnet-05f09df36061ff98f"]
  xff_header_processing_mode = "append"
}
