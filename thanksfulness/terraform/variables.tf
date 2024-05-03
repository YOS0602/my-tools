variable "project" {
  type     = string
  nullable = false
  # プロジェクトIDの要件について See https://cloud.google.com/resource-manager/docs/creating-managing-projects?hl=ja
  validation {
    condition     = length(var.project) >= 6 && length(var.project) <= 30
    error_message = "6～30 文字にする必要があります"
  }
  validation {
    condition     = can(regex("^[a-z][a-z0-9|-]{4,28}[a-z0-9]$", var.project))
    error_message = "小文字、数字、ハイフンのみを含めることができ、先頭は英文字でなければなりません。末尾にハイフンは使用できません。"
  }
  description = "GCP Project"
}

variable "region" {
  type        = string
  nullable    = true
  default     = "asia-northeast1"
  description = "GCP リソースを作成するデフォルトリージョン default: 東京"
}

variable "zone" {
  type        = string
  nullable    = true
  default     = "asia-northeast1-c"
  description = "GCP リソースを作成するデフォルトゾーン"
}

variable "run_image_url" {
  type = string
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
