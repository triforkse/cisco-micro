variable "region" { }
variable "secret_key" { }
variable "access_key" { }
variable "deployment_id" { }

provider "aws" {
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
  region = "${var.region}"
}

resource "aws_instance" "web" {
    # TODO: Use the deployment_id to avoid name conflicts
    ami = "ami-b5620ac2"
    instance_type = "m1.small"
    count = 1
}
