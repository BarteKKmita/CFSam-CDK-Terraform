terraform {

  backend "s3" {
    bucket         = "iac-presentation-tf-state"
    key            = "iac-presentation/prod/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform-state-locking"
    encrypt        = true
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}