# See https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#provider-default-values-configuration
terraform {
  required_version = ">=1.7.4"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">=5.17.0"
    }
  }
}

provider "google" {
  project = var.project
  region  = var.region
  zone    = var.zone
  default_labels = {
    p_name = "thanksfulness"
  }
}
