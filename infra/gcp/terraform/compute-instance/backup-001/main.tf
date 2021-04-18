resource "google_compute_instance" "default" {
  name         = "therentyapp-backup-001"
  machine_type = "e2-micro"
  zone         = "europe-west1-b"

  tags = ["backup"]

  boot_disk {
    auto_delete = false
    
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = var.compute_network_name
  }

  metadata = {
    env     = "prod",
    service = "backup"
  }

  service_account {
    email  = var.service_account_email
    scopes = ["cloud-platform"]
  }
}
