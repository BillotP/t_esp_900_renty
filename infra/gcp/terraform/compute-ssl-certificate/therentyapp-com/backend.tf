# Generated by Terragrunt. Sig: nIlQXj57tbuaRZEa
terraform {
  backend "gcs" {
    bucket = "therentyapp-tfstates"
    prefix = "compute-ssl-certificate/therentyapp-com/terraform.tfstate"
  }
}