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
    count = "1"
    security_groups = ["master_group"]
    key_name = "trifork-pub"
}

resource "aws_security_group" "master_security" {
  name = "master_group"
  description = "Allow all inbound traffic"

  tags {
    Name = "security_master"
  }
}

resource "aws_security_group_rule" "allow_ssh" {
    type = "ingress"
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]

    security_group_id = "${aws_security_group.master_security.id}"
}


resource "aws_security_group_rule" "allow_icmp" {
    type = "ingress"
    from_port=-1
    to_port=-1
    protocol = "icmp"
    cidr_blocks=["0.0.0.0/0"]

    security_group_id = "${aws_security_group.master_security.id}"
}


resource "aws_key_pair" "deployer" {
  key_name = "trifork-pub"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC4/eg7kv1hWYO9qKI1kPjDyk4mb4HZAqoxHlTQufbmkkT0QD7vrLH4vlb8nV+QjBgl/tMbiUj+7ZK4OaAvMro4oypMP+exAYZNVAHF0Dc/rLLlZcSFo3EBhCNKmhH5rQhOzcTRMfAzWA704PXwqfYBTmDnaDjrtLZ6IpHJrXEc6Kll7wdWwmLiNvN8OeqQOpHLg70ERo/xPhfUy9BJccLUW1bdrt1YcIxwPewTXCwesTjmWj6QtcMLhgyjcEKeTWu9hbC9g01Td1ENAiEEVjSt9G7esu4qHUwoqryldjbG8mhVvsG4QdE/I8wPxdSIU5dSviYRRzCQ9OBCaF6nL6EJ stockholm@trifork.com"
}

