output "cicd_service_account_id" {
  # an identifier for the resource with format projects/{{project}}/serviceAccounts/{{email}}
  value = google_service_account.cicd.id
}

output "cicd_service_account_email" {
  value = google_service_account.cicd.email
}

output "cicd_service_account_member" {
  # The Identity of the service account in the form serviceAccount:{email}. 
  # This value is often used to refer to the service account in order to grant IAM permissions.
  value = google_service_account.cicd.member
}
