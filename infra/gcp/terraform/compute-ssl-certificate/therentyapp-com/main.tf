resource "google_compute_ssl_certificate" "default" {
  name        = "cert-therentyapp-com"
  description = "Certificat pour le domaine therentyapp.com"
  private_key = file("./secrets/therentyapp.com.key")
  certificate = file("./secrets/therentyapp.com.crt")

  lifecycle {
    create_before_destroy = true
  }
}
