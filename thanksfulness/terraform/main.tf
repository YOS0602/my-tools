# See: https://registry.terraform.io/providers/hashicorp/google/latest/docs

module "thanksfulness_run" {
  source                        = "./modules/run"
  name                          = "thanksfulness"
  location                      = var.region
  image                         = var.run_image_url
  NOTION_API_KEY                = var.NOTION_API_KEY
  NOTION_THANKSFULNESS_BLOCK_ID = var.NOTION_THANKSFULNESS_BLOCK_ID
  SLACK_BOT_USER_OAUTH_TOKEN    = var.SLACK_BOT_USER_OAUTH_TOKEN
}

module "thanksfulness_ar" {
  source   = "./modules/artifact-registry"
  location = var.region
}

module "thanksfulness_cicd_sa" {
  source  = "./modules/service-accounts/"
  project = var.project
}
