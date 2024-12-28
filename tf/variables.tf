variable "project_name" {}
variable "env" {}
variable "vpc_id" {}
variable "region" {}
variable "dd_api_key" {}
variable "dd_app_key" {}
variable "certificate_arn" {}
variable "application_version" {}
variable "cloudflare_api_token" {}
variable "container_image" {
  type        = string
  description = "container image"
  default     = "dummy"
}
variable "ingress_allowed_cidrs_prod" {
  default     = ["10.0.0.0/16"]
  description = "List of allowed cidrs ranges"
}
variable "container_cpu" {
  type        = number
  description = "Container CPU associated to it."
  default     = 8192
}
variable "container_memory" {
  type        = number
  description = "Container MEMORY associated to it."
  default     = 16384
}
variable "desired_count_prod" {
  type        = number
  description = "Desired number of tasks that must be running in production environments."
  default     = 6
}
variable "team" {
  type        = string
  description = "team name"
  default     = "web"
}
variable "opsgenie" {
  default     = "opsgenie-web"
  description = "Opsgenie channel alarm"
}
variable "tags" {
  type        = list(string)
  default     = ["web", "thrive-api"]
  description = "Tags related to the monitor"
}

variable "grace_period" {
  type = number
  description = "Time in seconds after instance comes into service before checking health."
  default = 300
}

variable "dns_name_app" {
  type        = string
  description = "DNS name"
  default     = "thrive-api"
}

variable "internal" {
  type = string
  description = "App visibility"
  default = "no"
}
