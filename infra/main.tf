terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.8.0"
    }
  }
}

resource "aws_s3_bucket" "test" {
  bucket = "jeff-${var.bucket_suffix}"
}

variable "bucket_suffix" {
  default = "test"
  type = string
}