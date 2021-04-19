output "compute_instance_self_link" {
  value = google_compute_instance.default.self_link
}

output "compute_instance_metadata_service" {
  value = google_compute_instance.default.metadata.service
}
