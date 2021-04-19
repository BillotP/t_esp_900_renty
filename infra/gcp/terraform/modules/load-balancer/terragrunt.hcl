include {
  path = find_in_parent_folders()
}

dependency "compute_global_address_lb_monitoring_therentyapp_com" {
  config_path = "../../compute-global-address/lb-monitoring-therentyapp-com"

  mock_outputs = {
    compute_global_address_self_link = "mock_compute_global_address_self_link"
  }
}

dependency "compute_ssl_certificate_therentyapp_com" {
  config_path = "../../compute-ssl-certificate/therentyapp-com"

  mock_outputs = {
    compute_ssl_certificate_self_link = "mock_compute_ssl_certificate_self_link"
  }
}

dependency "compute_instance_group_monitoring" {
  config_path = "../../compute-instance-group/monitoring"

  mock_outputs = {
    compute_instance_group_named_port_name = "mock_compute_instance_group_named_port_name"
    compute_instance_group_named_port_port = "3000"
    compute_instance_group_self_link       = "mock_compute_instance_group_name"
  }
}

inputs = {
  compute_global_address_self_link       = dependency.compute_global_address_lb_monitoring_therentyapp_com.outputs.compute_global_address_self_link
  compute_instance_group_named_port_name = dependency.compute_instance_group_monitoring.outputs.compute_instance_group_named_port_name
  compute_instance_group_named_port_port = dependency.compute_instance_group_monitoring.outputs.compute_instance_group_named_port_port
  compute_instance_group_self_link       = dependency.compute_instance_group_monitoring.outputs.compute_instance_group_self_link
  compute_ssl_certificate_self_link      = dependency.compute_ssl_certificate_therentyapp_com.outputs.compute_ssl_certificate_self_link
}

generate "variables" {
    path      = "variables.tf"
    if_exists = "overwrite_terragrunt"
    contents  = <<EOF
variable "compute_global_address_self_link" {}
variable "compute_instance_group_named_port_name" {}
variable "compute_instance_group_named_port_port" {}
variable "compute_instance_group_self_link" {}
variable "compute_ssl_certificate_self_link" {}
EOF
}
