resource "google_compute_network" "default" {
  name                    = "default"
  auto_create_subnetworks = true
  description             = "Default network for the project"
}
