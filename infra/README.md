# IAC

Import stateless resources that can be destroyed and recreated easily. Stateful services such as IAM and Route53 shouldn't be destroyed.

- alb
- ecs
- eip
- logs
- nat
- route53
- sg
- vpc_endpoint

# Import and manage a new resource

Sign in to aws console to connect terraform to AWS

```bash
aws sso login --profile [PROFILE]`
export AWS_PROFILE=[PROFILE]
```

Create module for service, such as `/generated/aws/route53/route53.tf` and create module in `main.tf`. Write resources block and whatever data block required to match the existing infra on AWS. See AWS provider [docs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs).

Import from AWS infra into module. E.g.

```bash
terraform import module.route53.aws_route53_record.www_dev_piclimb_com Z0081576APTH2MKBFDT2_www.dev.piclimb.com_A
```

Use `terraform plan` to ensure that local config matches infra. Make changes till it matches. Reference variables from other resources using `variables.tf` and `outputs.tf`, and link via modules in `main.tf`.
