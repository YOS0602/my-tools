variable "name" {
  type        = string
  nullable    = false
  description = "Cloud Run Service Name"
}

variable "location" {
  type = string
}

variable "image" {
  type        = string
  description = "URL of the Container image in Google Container Registry or Google Artifact Registry"
}

variable "NOTION_API_KEY" {
  type      = string
  sensitive = true
  nullable  = false
}

variable "NOTION_THANKSFULNESS_BLOCK_ID" {
  type      = string
  sensitive = true
  nullable  = false
}

variable "SLACK_BOT_USER_OAUTH_TOKEN" {
  type      = string
  sensitive = true
  nullable  = false
}
