# Generated by Terragrunt. Sig: nIlQXj57tbuaRZEa
terraform {
  backend "gcs" {
    bucket = "therentyapp-tfstates"
    prefix = "compute-instance-group/monitoring/terraform.tfstate"
  }
}