variable "region" { }
variable "secret_key" { }
variable "access_key" { }

provider "aws" {
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
  region = "${var.region}"
}

resource "aws_instance" "web" {
    ami = "ami-b5620ac2"
    instance_type = "m1.small"
    count = 1
}
