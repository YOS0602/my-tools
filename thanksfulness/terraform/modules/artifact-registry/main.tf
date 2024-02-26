# See https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/artifact_registry_repository
resource "google_artifact_registry_repository" "default" {
  location      = var.location
  repository_id = "thanksfulness"
  format        = "DOCKER"
  description   = "thanksfulness repository"

  docker_config {
    # immutable_tags = true
  }
  # See https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories#CleanupPolicy
  cleanup_policies {
    id     = "clean-untagged"
    action = "DELETE"
    condition {
      tag_state = "UNTAGGED"
    }
  }
}
