# Configure the GitHub Provider

# Configure the GitHub Provider

variable "github_token" {}
variable "github_organization" {}

provider "github" {
  token        = "${var.github_token}"
  organization = "${var.github_organization}"
}

# Import the helm-chart module
module "helm-chart" {
  source  = "./module-helm-chart"
  name    = "nats-operator"
}