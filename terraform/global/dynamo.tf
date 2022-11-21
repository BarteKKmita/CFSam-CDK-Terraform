terraform {

  backend "s3" {
    bucket         = "iac-presentation-tf-state"
    key            = "iac-presentation/global/terraform.tfstate"
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

provider "aws" {
  region = "us-east-1"
}

resource "aws_dynamodb_table" "IaCPresentationDatabase" {
  name = "IaCPresentationTable"
  billing_mode = "PAY_PER_REQUEST"
  attribute {
    name = "id"
    type = "S"
  }
  attribute {
    name = "username"
    type = "S"
  }
  hash_key = "id"
  range_key = "username"
}