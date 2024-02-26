# See: https://registry.terraform.io/providers/hashicorp/google/latest/docs

module "thanksfulness_run" {
  source   = "./modules/run"
  name     = "thanksfulness"
  location = var.region
  image    = var.run_image_url
}

module "thanksfulness_ar" {
  source   = "./modules/artifact-registry"
  location = var.region
}
