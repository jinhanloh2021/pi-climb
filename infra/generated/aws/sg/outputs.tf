output "aws_security_group_tfer--pi-climb-endpoint-sg_sg-01e41717258f067bc_id" {
  description = "Security Group ID of pi-climb endpoint, used for internal AWS api calls"
  value       = aws_security_group.tfer--pi-climb-endpoint-sg_sg-01e41717258f067bc.id
}

output "aws_security_group_tfer--pi-climb-lb-sg_sg-02821e5f312595268_id" {
  description = "Security Group ID of load balancer"
  value       = aws_security_group.tfer--pi-climb-lb-sg_sg-02821e5f312595268.id
}

output "aws_security_group_tfer--pi-climb-service-sg_sg-0802e5a19b6cc4611_id" {
  description = "Security Group ID of pi-climb service"
  value       = aws_security_group.tfer--pi-climb-service-sg_sg-0802e5a19b6cc4611.id
}
