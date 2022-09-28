scenario "dynamic_hosts" {
  terraform_cli = terraform_cli.default
  terraform     = terraform.default
  providers = [
    provider.aws.default,
    provider.enos.default
  ]

  matrix {
    builder = ["local", "crt"]
  }

  locals {
    aws_ssh_private_key_path = abspath(var.aws_ssh_private_key_path)
    boundary_install_dir     = abspath(var.boundary_install_dir)
    local_boundary_dir       = abspath(var.local_boundary_dir)
    build_path = {
      "local" = "/tmp",
      "crt"   = var.crt_bundle_path == null ? null : abspath(var.crt_bundle_path)
    }
  }

  step "find_azs" {
    module = module.az_finder

    variables {
      instance_type = [
        var.worker_instance_type,
        var.controller_instance_type
      ]
    }
  }

  step "create_db_password" {
    module = module.random_stringifier
  }

  step "build_boundary" {
    module = matrix.builder == "crt" ? module.build_crt : module.build_local

    variables {
      path = local.build_path[matrix.builder]
    }
  }

  step "create_base_infra" {
    module = module.infra

    variables {
      availability_zones = step.find_azs.availability_zones
    }
  }

  step "create_boundary_cluster" {
    module = module.boundary
    depends_on = [
      step.create_base_infra,
      step.build_boundary
    ]

    variables {
      boundary_install_dir     = local.boundary_install_dir
      controller_instance_type = var.controller_instance_type
      controller_count         = var.controller_count
      db_pass                  = step.create_db_password.string
      kms_key_arn              = step.create_base_infra.kms_key_arn
      local_artifact_path      = step.build_boundary.artifact_path
      ubuntu_ami_id            = step.create_base_infra.ami_ids["ubuntu"]["amd64"]
      vpc_id                   = step.create_base_infra.vpc_id
      worker_count             = var.worker_count
      worker_instance_type     = var.worker_instance_type
    }
  }

  step "create_vault_cluster" {
    module = module.vault
    depends_on = [
      step.create_base_infra,
    ]

    variables {
      ami_id            = step.create_base_infra.ami_ids["ubuntu"]["amd64"]
      instance_type     = var.vault_instance_type
      instance_count    = 1
      kms_key_arn       = step.create_base_infra.kms_key_arn
      storage_backend   = "raft"
      sg_additional_ips = step.create_boundary_cluster.controller_ips
      unseal_method     = "awskms"
      vault_release = {
        version = "1.11.0"
        edition = "oss"
      }
      vpc_id = step.create_base_infra.vpc_id
    }
  }

  step "create_target_bad" {
    module     = module.target
    depends_on = [step.create_base_infra]

    variables {
      ami_id               = step.create_base_infra.ami_ids["ubuntu"]["amd64"]
      aws_ssh_keypair_name = var.aws_ssh_keypair_name
      enos_user            = var.enos_user
      instance_type        = var.target_instance_type
      target_count         = 1
      vpc_id               = step.create_base_infra.vpc_id
      username             = var.ssh_username
      password             = var.ssh_password
    }
  }

  step "create_target_good" {
    module     = module.target
    depends_on = [step.create_base_infra]

    variables {
      ami_id               = step.create_base_infra.ami_ids["ubuntu"]["amd64"]
      aws_ssh_keypair_name = var.aws_ssh_keypair_name
      enos_user            = var.enos_user
      instance_type        = var.target_instance_type
      target_count         = 1
      vpc_id               = step.create_base_infra.vpc_id
      username             = var.ssh_username
      password             = var.ssh_password
      additional_tags      = { "test_key" : "test_val" }
    }
  }

  step "iam_setup" {
    module     = module.iam_setup
    depends_on = [step.create_base_infra]
  }

  output "boundary_addr" {
    value = step.create_boundary_cluster.alb_boundary_api_addr
  }
  output "auth_method_id" {
    value = step.create_boundary_cluster.auth_method_id
  }
  output "auth_login_name" {
    value = step.create_boundary_cluster.auth_login_name
  }
  output "auth_password" {
    value = step.create_boundary_cluster.auth_password
  }
  output "vault_addr" {
    value = step.create_vault_cluster.instance_public_ips[0]
  }
  output "vault_root_token" {
    value = step.create_vault_cluster.vault_root_token
  }

  output "bad_target" {
    value = step.create_target_bad.target_ips
  }

  output "good_target" {
    value = step.create_target_good.target_ips
  }

  output "access_key_id" {
    value = step.iam_setup.access_key_id
  }

  output "secret_access_key" {
    value     = step.iam_setup.secret_access_key
    sensitive = true
  }
}


