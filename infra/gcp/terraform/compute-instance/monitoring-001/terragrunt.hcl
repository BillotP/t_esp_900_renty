include {
  path = find_in_parent_folders()
}

dependency "compute_address_therentyapp_monitoring_001" {
  config_path = "../../compute-address/therentyapp-monitoring-001"

  mock_outputs = {
    compute_address_address = "mock_compute_address_address"
  }
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
  compute_address_address = dependency.compute_address_therentyapp_monitoring_001.outputs.compute_address_address
  compute_network_name    = dependency.compute_network_default.outputs.compute_network_name
  service_account_email   = dependency.service_account_backup.outputs.service_account_email
}

generate "variables" {
    path      = "variables.tf"
    if_exists = "overwrite_terragrunt"
    contents  = <<EOF
variable "compute_address_address" {}
variable "compute_network_name" {}
variable "service_account_email" {}
EOF
}
