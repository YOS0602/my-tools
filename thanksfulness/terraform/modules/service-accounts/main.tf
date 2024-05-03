# See https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_service_account
resource "google_service_account" "cicd" {
  account_id                   = "cicd-service-account"
  display_name                 = "Service Account for GitHub Actions"
  create_ignore_already_exists = true
}

# See https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project_iam
resource "google_project_iam_member" "cicd-account-iam-run-developer" {
  project = var.project
  # https://cloud.google.com/run/docs/deploying?hl=ja#permissions_required_to_deploy
  role   = "roles/run.developer"
  member = google_service_account.cicd.member
}
resource "google_project_iam_member" "cicd-account-iam-iam-serviceAccountUser" {
  project = var.project
  # https://cloud.google.com/run/docs/reference/iam/roles#permissions-required-for-the-deploying-services-revisions
  role   = "roles/iam.serviceAccountUser"
  member = google_service_account.cicd.member
}

resource "google_project_iam_member" "cicd-account-iam-artifactregistry-writer" {
  project = var.project
  # https://cloud.google.com/artifact-registry/docs/access-control?hl=ja#grant
  role   = "roles/artifactregistry.writer"
  member = google_service_account.cicd.member
}
