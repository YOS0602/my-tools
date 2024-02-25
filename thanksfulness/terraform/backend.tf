terraform {
  # See https://developer.hashicorp.com/terraform/language/settings/backends/gcs
  backend "gcs" {
    bucket = "terraform-state-nereid"
    prefix = "thanksfulness/prd"
  }
}
