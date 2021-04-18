include {
  path = find_in_parent_folders()
}

dependency "compute_network_default" {
  config_path = "../../compute-network/default"

  mock_outputs = {
    compute_network_name = "mock_compute_network_name"
  }
}

dependency "service_account_backup" {
  config_path = "../../service-account/backup"

  mock_outputs = {
    service_account_email = "mock_service_account_email"
  }
}

inputs = {
  compute_network_name  = dependency.compute_network_default.outputs.compute_network_name
  service_account_email = dependency.service_account_backup.outputs.service_account_email
}

generate "variables" {
    path      = "variables.tf"
    if_exists = "overwrite_terragrunt"
    contents  = <<EOF
variable "compute_network_name" {}
variable "service_account_email" {}
EOF
}
