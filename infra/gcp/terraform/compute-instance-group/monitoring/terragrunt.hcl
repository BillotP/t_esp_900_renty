include {
  path = find_in_parent_folders()
}

dependency "compute_instance_monitoring_001" {
  config_path = "../../compute-instance/monitoring-001"

  mock_outputs = {
    compute_instance_metadata_service = "mock_compute_instance_metadata_service"
    compute_instance_self_link        = "mock_compute_instance_self_link"
  }
}

inputs = {
  compute_instance_metadata_service = dependency.compute_instance_monitoring_001.outputs.compute_instance_metadata_service
  compute_instance_self_link        = dependency.compute_instance_monitoring_001.outputs.compute_instance_self_link
}

generate "variables" {
    path      = "variables.tf"
    if_exists = "overwrite_terragrunt"
    contents  = <<EOF
variable "compute_instance_metadata_service" {}
variable "compute_instance_self_link" {}
EOF
}
