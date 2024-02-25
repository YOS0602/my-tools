variable "project" {
  type     = string
  nullable = false
  validation {
    condition     = length(var.project) >= 6 && length(var.project) <= 30 && can(regex("^[a-z][a-z0-9]*[a-z0-9]$", var.project))
    error_message = "GCPプロジェクトIDはこちらに記載された要件を満たす必要がある。\n https://cloud.google.com/resource-manager/docs/creating-managing-projects?hl=ja"
  }
  description = "GCP Project"
}

variable "region" {
  type        = string
  nullable    = true
  default     = "asia-northeast1"
  description = "GCP リソースを作成するデフォルトリージョン"
}

variable "zone" {
  type        = string
  nullable    = true
  default     = "asia-northeast1-c"
  description = "GCP リソースを作成するデフォルトゾーン"
}
