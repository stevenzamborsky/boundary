variable "vpc_id" {}
variable "ami_id" {}
variable "target_count" {}
variable "environment" {}
variable "project_name" {}
variable "instance_type" {}
variable "aws_ssh_keypair_name" {}
variable "enos_user" {}
variable "additional_tags" {
  default = {}
}

data "aws_subnets" "infra" {
  filter {
    name   = "vpc-id"
    values = [var.vpc_id]
  }
}

resource "aws_security_group" "boundary_target" {
  name_prefix = "boundary-target-sg"
  description = "SSH and boundary Traffic"
  vpc_id      = var.vpc_id

  # SSH
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/8"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    "Name" : "boundary-target-sg"
  }
}

resource "aws_instance" "target" {
  count                  = var.target_count
  ami                    = var.ami_id
  instance_type          = var.instance_type
  vpc_security_group_ids = [aws_security_group.boundary_target.id]
  subnet_id              = tolist(data.aws_subnets.infra.ids)[count.index % length(data.aws_subnets.infra.ids)]
  key_name               = var.aws_ssh_keypair_name
  user_data = templatefile("${path.module}/templates/user_data.sh.tpl", {
    username = var.username,
    password = var.password
  })

  tags = merge(var.additional_tags, {
    "Name" : "boundary-target-${count.index}",
    "Type" : "target",
    "Environment" : var.environment
    "Enos User" : var.enos_user,
  })
}

output "target_ips" {
  value = aws_instance.target.*.private_ip
}
output "target_public_ips" {
  value = aws_instance.target.*.public_ip
}

# Corpsec folks: yes we are enabling password based SSH auth (sorry) to test user/password brokering in Boundary,
# BUT this instance is locked down via security groups to only the IP running enos and itself
variable "username" {
  default = "test"
}

variable "password" {
  default = "Hunter2!"
}
