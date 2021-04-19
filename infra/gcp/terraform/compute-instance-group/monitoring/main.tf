resource "google_compute_instance_group" "default" {
  name      = "ig-${var.compute_instance_metadata_service}"
  zone      = "europe-west1-b"
  instances = [var.compute_instance_self_link]
  
  named_port {
    name = var.compute_instance_metadata_service
    port = "3000"
  }

  lifecycle {
    create_before_destroy = true
  }
}
