resource "google_compute_instance" "default" {
  name         = "therentyapp-monitoring-001"
  machine_type = "n1-standard-1"
  zone         = "europe-west1-b"

  tags = ["monitoring"]

  boot_disk {
    auto_delete = false
    
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = var.compute_network_name

    access_config {

    }
  }

  metadata = {
    env     = "prod",
    service = "monitoring"
  }

  service_account {
    email  = var.service_account_email
    scopes = ["cloud-platform"]
  }
}