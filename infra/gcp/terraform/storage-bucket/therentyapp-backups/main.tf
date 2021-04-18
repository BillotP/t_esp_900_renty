resource "google_storage_bucket" "default" {
  name          = "therentyapp-backups"
  location      = "EU"
  force_destroy = false

  lifecycle_rule {
    condition {
      age = 14
    }
    action {
      type = "Delete"
    }
  }
}
