provider "aws" {
  region = "us-east-1"
}

# Uses an array of strings and converts it into a "set"
resource "aws_s3_bucket" "crBucket" {
  for_each = toset(["mickey", "mouse"])
  bucket   = "723paladin-tf-demo-${each.key}"
}