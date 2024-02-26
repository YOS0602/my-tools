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
