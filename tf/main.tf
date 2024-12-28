terraform {
  required_version = ">=1.0.0"
  backend "s3" {
    encrypt = true
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.22.0"
    }
  }
}

provider "aws" {
  region = var.region
  default_tags {
    tags = {
      Environment = var.env
      Project     = var.project_name
      Team        = "web"
      ManagedBy   = "terraform"
    }
  }
}

module "identify" {
  source = "git@github.com:vitingr/vitingr-terraform-modules.git//identify?ref=v1.8.9"
}

module "secrets" {
  source       = "git@github.com:vitingr/vitingr-terraform-modules.git//secrets?ref=v1.8.9"
  project_name = var.project_name
  env          = var.env
}

module "cloudwatch" {
  source       = "git@github.com:vitingr/vitingr-terraform-modules.git//cloudwatch?ref=v1.8.9"
  project_name = var.project_name
  env          = var.env
}

module "cluster" {
  source       = "git@github.com:vitingr/vitingr-terraform-modules.git//cluster?ref=v1.8.9"
  project_name = var.project_name
  env          = var.env
}

module "ecr" {
  source       = "git@github.com:vitingr/vitingr-terraform-modules.git//ecr?ref=v1.8.9"
  project_name = var.project_name
  env          = var.env
}

module "data" {
  source = "git@github.com:vitingr/vitingr-terraform-modules.git//data?ref=v1.8.9"
  env    = var.env
  vpc_id = var.vpc_id
}

module "iam" {
  source       = "git@github.com:vitingr/vitingr-portunus.git//roles/main"
  project_name = var.project_name
  account      = module.identify.account
  region       = var.region
  env          = var.env
}

module "lb" {
  source               = "git@github.com:vitingr/vitingr-terraform-modules.git//lb?ref=v1.8.9"
  project_name         = var.project_name
  env                  = var.env
  security_group_lb_id = module.balancerSg.security_group_cloudFlare_id
  certificate_arn      = var.certificate_arn
  private_subnet_id    = module.data.private_subnet_id
  public_subnet_id     = module.data.public_subnet_id
  vpc_id               = var.vpc_id
  internal             = var.internal
}

module "taskWithBalancerSg" {
  source               = "git@github.com:vitingr/vitingr-terraform-modules.git//sg/dynamicSg?ref=v1.8.9"
  project_name         = var.project_name
  env                  = var.env
  vpc_id               = var.vpc_id
  security_groups      = module.balancerSg.security_group_cloudFlare_id
  internal             = var.internal
}

module "balancerSg" {
  source                     = "git@github.com:vitingr/vitingr-terraform-modules.git//sg/cloudFlareSg?ref=v1.8.9"
  project_name               = var.project_name
  env                        = var.env
  vpc_id                     = var.vpc_id
}

module "publicService" {
  source                         = "git@github.com:vitingr/vitingr-terraform-modules.git//service/publicService?ref=v1.8.9"
  project_name                   = var.project_name
  cluster_name                   = module.cluster.cluster_name
  tg_arn                         = module.lb.tg_arn
  task_id                        = module.task.task_id
  grace_period                   = var.grace_period
  desired_count_prod             = var.desired_count_prod
  private_subnet_id              = module.data.private_subnet_id
  public_subnet_id               = module.data.public_subnet_id
  env                            = var.env
  security_group_dynamic_id      = module.taskWithBalancerSg.security_group_dynamic_id
}

module "cloudFlareRecord" {
  source               = "git@github.com:vitingr/vitingr-terraform-modules.git//cloudFlare?ref=v1.8.9"
  env                  = var.env
  project_name         = var.project_name
  dns_name             = var.dns_name_app
  domain_name          = module.lb.dns_name
  cloudflare_api_token = var.cloudflare_api_token
}


module "task" {
  source              = "git@github.com:vitingr/vitingr-terraform-modules.git//task?ref=v1.8.9"
  project_name        = var.project_name
  env                 = var.env
  region              = var.region
  container_image     = var.container_image
  container_memory    = var.container_memory
  container_cpu       = var.container_cpu
  account             = module.identify.account
  awslogs_name        = module.cloudwatch.awslogs_name
  dd_api_key          = var.dd_api_key
  application_version = var.application_version
}

module "datadog" {
  source       = "git@github.com:vitingr/vitingr-terraform-modules.git//datadog?ref=v1.8.9"
  project_name = var.project_name
  team         = var.team
  dd_api_key   = var.dd_api_key
  dd_app_key   = var.dd_app_key
}

module "datadog-monitors-cloud-aws-elb" {
  source      = "git@github.com:vitingr/vitingr-terraform-modules.git//datadog/elb?ref=v1.8.9"

  dd_api_key   = var.dd_api_key
  dd_app_key   = var.dd_app_key

  project_name = var.project_name
  team         = var.team
  elb_name     = module.lb.name
  opsgenie     = var.opsgenie
  tags         = var.tags
  env          = var.env
}
