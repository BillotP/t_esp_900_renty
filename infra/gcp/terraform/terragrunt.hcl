generate "provider" {
    path      = "provider.tf"
    if_exists = "overwrite_terragrunt"
    contents  = <<EOF
provider "google" {
   project = "therentyapp"
   region  = "europe-west1-b"
}
EOF
}

remote_state {
    backend = "gcs"

    generate = {
        path      = "backend.tf"
        if_exists = "overwrite_terragrunt"
    }

    config = {
      project  = "therentyapp"
      bucket   = "therentyapp-tfstates"
      prefix   = "${path_relative_to_include()}/terraform.tfstate"
      location = "eu"
    }
}
