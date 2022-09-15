variable "ssh_username" {}
variable "ssh_password" {}

variable "vault_path" {
  default = "/usr/local/bin/vault"
}

variable "vault_ip" {
  type = string
}

variable "vault_token" {
  type      = string
  sensitive = true
}

variable "secret_path" {
  type    = string
  default = "secret/my-app-secret"
}

locals {
  vault_addr = "http://${var.vault_ip}:8200"
}
terraform {
  required_providers {
    enos = {
      source = "app.terraform.io/hashicorp-qti/enos"
    }
  }
}

resource "enos_local_exec" "enable_kv2" {
  environment = {
    VAULT_ADDR  = local.vault_addr
    VAULT_TOKEN = var.vault_token
  }
  inline = ["${var.vault_path} secrets enable -path=${split("/", var.secret_path)[0]} -version=2 kv"]
}

resource "enos_local_exec" "apply_policies" {
  environment = {
    VAULT_ADDR  = local.vault_addr
    VAULT_TOKEN = var.vault_token
  }
  for_each = {
    "boundary_controller" = file("${path.module}/templates/boundary-controller-policy.hcl"),
    "kv-read"             = templatefile("${path.module}/templates/kv-policy.hcl", { secret_path = var.secret_path })
  }
  inline = ["echo '${each.value}' | ${var.vault_path} policy write ${each.key} -"]
}


resource "enos_local_exec" "add_credentials" {
  depends_on = [
    enos_local_exec.apply_policies,
    enos_local_exec.enable_kv2
  ]
  environment = {
    VAULT_ADDR  = local.vault_addr
    VAULT_TOKEN = var.vault_token
  }
  inline = ["${var.vault_path} kv put ${var.secret_path} username=${var.ssh_username} password='${var.ssh_password}'"]
}
