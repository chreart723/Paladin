provider "aws" {
  region = "us-east-1"
}

# Uses a map for two buckets (These will be different in real world, just for demo)
resource "aws_s3_bucket" "crBucket" {
  for_each = {
    dev  = "Dev"
    prod = "Prod"
  }

  bucket = "723paladin-tf-demo-${each.key}"
  tags = {
    Environment = "${each.value}"
  }
}