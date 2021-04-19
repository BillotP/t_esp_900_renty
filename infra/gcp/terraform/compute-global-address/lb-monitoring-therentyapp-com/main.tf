resource "google_compute_global_address" "default" {
  name        = "lb-monitoring-therentyapp-com-external-ip"
  description = "IP externe du load-balancer de monitoring.therentyapp.com"
}
