provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "crBucket" {
  count  = 2
  bucket = "723paladin-tf-demo-${count.index}"
}