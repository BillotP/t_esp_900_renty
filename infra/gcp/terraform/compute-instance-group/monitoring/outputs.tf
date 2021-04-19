output "compute_instance_group_self_link" {
  value = google_compute_instance_group.default.self_link
}

output "compute_instance_group_named_port_name" {
  value = google_compute_instance_group.default.named_port[0].name
}

output "compute_instance_group_named_port_port" {
  value = google_compute_instance_group.default.named_port[0].port
}
