terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }

  required_version = ">= 0.14.9"
}

provider "aws" {
  profile = "default"
  region  = "us-east-1"
}

resource "aws_instance" "ubuntu_server" {
  ami                         = "ami-029c64b3c205e6cce"
  instance_type               = "t4g.micro"
  associate_public_ip_address = true
  vpc_security_group_ids = [ "sg-0f60af9c0c3331a40" ]

  root_block_device {
    volume_size = 8
  }

  key_name = "aws"

  tags = {
    Name = "Ubuntu Server"
  }
}
