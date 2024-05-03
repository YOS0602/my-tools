# See https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloud_run_v2_service
resource "google_cloud_run_v2_service" "default" {
  name     = var.name
  location = var.location
  ingress  = "INGRESS_TRAFFIC_ALL"

  # See https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services#revisiontemplate
  template {
    timeout = "60s"
    scaling {
      max_instance_count = 1
      min_instance_count = 0
    }
    containers {
      image = var.image
      env {
        name  = "NOTION_API_KEY"
        value = var.NOTION_API_KEY
      }
      env {
        name  = "NOTION_THANKSFULNESS_BLOCK_ID"
        value = var.NOTION_THANKSFULNESS_BLOCK_ID
      }
      env {
        name  = "SLACK_BOT_USER_OAUTH_TOKEN"
        value = var.SLACK_BOT_USER_OAUTH_TOKEN
      }
      resources {
        limits = {
          cpu    = "1"
          memory = "512Mi"
        }
        cpu_idle          = true  # Request 処理中のみCPU割り当て
        startup_cpu_boost = false # cold start対策としての起動時CPUブースト
      }
    }
    max_instance_request_concurrency = 1
  }
}

