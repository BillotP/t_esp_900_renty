module "gce-lb-http" {
  source               = "GoogleCloudPlatform/lb-http/google"
  version              = "~> 4.4"
  project              = "therentyapp"
  name                 = "lb-monitoring-therentyapp-com"
  address              = var.compute_global_address_self_link
  create_address       = false
  https_redirect       = true
  ssl                  = true
  use_ssl_certificates = true
  ssl_certificates     = [var.compute_ssl_certificate_self_link]
  backends = {
    default = {
      description                     = null
      protocol                        = "HTTP"
      port                            = var.compute_instance_group_named_port_port
      port_name                       = var.compute_instance_group_named_port_name
      timeout_sec                     = 10
      enable_cdn                      = false
      custom_request_headers          = null
      security_policy                 = null

      connection_draining_timeout_sec = null
      session_affinity                = null
      affinity_cookie_ttl_sec         = null

      health_check = {
        check_interval_sec  = null
        timeout_sec         = null
        healthy_threshold   = null
        unhealthy_threshold = null
        request_path        = "/login"
        port                = var.compute_instance_group_named_port_port
        host                = null
        logging             = null
      }

      log_config = {
        enable      = true
        sample_rate = 1.0
      }

      groups = [
        {
          # Each node pool instance group should be added to the backend.
          group                        = var.compute_instance_group_self_link
          balancing_mode               = null
          capacity_scaler              = null
          description                  = null
          max_connections              = null
          max_connections_per_instance = null
          max_connections_per_endpoint = null
          max_rate                     = null
          max_rate_per_instance        = null
          max_rate_per_endpoint        = null
          max_utilization              = null
        },
      ]

      iap_config = {
        enable               = false
        oauth2_client_id     = null
        oauth2_client_secret = null
      }
    }
  }
}
