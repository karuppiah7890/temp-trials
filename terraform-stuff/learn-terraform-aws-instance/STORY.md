
https://www.toptal.com/developers/gitignore/

https://www.toptal.com/developers/gitignore/api/terraform

---

https://learn.hashicorp.com/tutorials/terraform/aws-build?in=terraform/aws-get-started

---

https://duckduckgo.com/?t=ffab&q=terraform+providers&ia=web

https://registry.terraform.io/browse/providers

https://registry.terraform.io/providers/hashicorp/aws/latest

https://registry.terraform.io/providers/hashicorp/aws/latest/docs

https://registry.terraform.io/providers/hashicorp/aws/latest/docs#environment-variables

https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance

https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance#associate_public_ip_address

https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance#root_block_device

https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance#vpc_security_group_ids

---

Resources required to create an EC2 instance with SSH access to it or any port access to it from outside from the public Internet from any IP based on my experimentation and experience -

Key Pair
VPC
Subnet
Network interface
EC2 instance
Security Group
Route Table
Internet Gateway

---

```bash
learn-terraform-aws-instance $ terraform init

Initializing the backend...

Initializing provider plugins...
- Finding hashicorp/aws versions matching "~> 3.27"...
- Installing hashicorp/aws v3.60.0...
- Installed hashicorp/aws v3.60.0 (signed by HashiCorp)

Terraform has created a lock file .terraform.lock.hcl to record the provider
selections it made above. Include this file in your version control repository
so that Terraform can guarantee to make the same selections by default when
you run "terraform init" in the future.

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
learn-terraform-aws-instance $ source .env 
learn-terraform-aws-instance $ echo $AWS_
$AWS_ACCESS_KEY_ID      $AWS_SECRET_ACCESS_KEY  
learn-terraform-aws-instance $ terraform plan
╷
│ Error: Invalid AWS Region: us-eas1-1
│ 
│   with provider["registry.terraform.io/hashicorp/aws"],
│   on main.tf line 12, in provider "aws":
│   12: provider "aws" {
│ 
╵
learn-terraform-aws-instance $ terraform plan
╷
│ Error: Invalid AWS Region: us-eas1-1
│ 
│   with provider["registry.terraform.io/hashicorp/aws"],
│   on main.tf line 12, in provider "aws":
│   12: provider "aws" {
│ 
╵
learn-terraform-aws-instance $ ls
STORY.md	main.tf
learn-terraform-aws-instance $ gst
On branch master
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)

Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	new file:   .gitignore
	new file:   STORY.md
	new file:   main.tf

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	modified:   STORY.md
	modified:   main.tf

Untracked files:
  (use "git add <file>..." to include in what will be committed)
	.terraform.lock.hcl

learn-terraform-aws-instance $ ga .
learn-terraform-aws-instance $ terraform plan

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with
the following symbols:
  + create

Terraform will perform the following actions:

  # aws_instance.ubuntu_server will be created
  + resource "aws_instance" "ubuntu_server" {
      + ami                                  = "ami-029c64b3c205e6cce"
      + arn                                  = (known after apply)
      + associate_public_ip_address          = true
      + availability_zone                    = (known after apply)
      + cpu_core_count                       = (known after apply)
      + cpu_threads_per_core                 = (known after apply)
      + disable_api_termination              = (known after apply)
      + ebs_optimized                        = (known after apply)
      + get_password_data                    = false
      + host_id                              = (known after apply)
      + id                                   = (known after apply)
      + instance_initiated_shutdown_behavior = (known after apply)
      + instance_state                       = (known after apply)
      + instance_type                        = "t4g.micro"
      + ipv6_address_count                   = (known after apply)
      + ipv6_addresses                       = (known after apply)
      + key_name                             = "aws"
      + monitoring                           = (known after apply)
      + outpost_arn                          = (known after apply)
      + password_data                        = (known after apply)
      + placement_group                      = (known after apply)
      + primary_network_interface_id         = (known after apply)
      + private_dns                          = (known after apply)
      + private_ip                           = (known after apply)
      + public_dns                           = (known after apply)
      + public_ip                            = (known after apply)
      + secondary_private_ips                = (known after apply)
      + security_groups                      = (known after apply)
      + source_dest_check                    = true
      + subnet_id                            = (known after apply)
      + tags                                 = {
          + "Name" = "Ubuntu Server"
        }
      + tags_all                             = {
          + "Name" = "Ubuntu Server"
        }
      + tenancy                              = (known after apply)
      + user_data                            = (known after apply)
      + user_data_base64                     = (known after apply)
      + vpc_security_group_ids               = (known after apply)

      + capacity_reservation_specification {
          + capacity_reservation_preference = (known after apply)

          + capacity_reservation_target {
              + capacity_reservation_id = (known after apply)
            }
        }

      + ebs_block_device {
          + delete_on_termination = (known after apply)
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + snapshot_id           = (known after apply)
          + tags                  = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = (known after apply)
          + volume_type           = (known after apply)
        }

      + enclave_options {
          + enabled = (known after apply)
        }

      + ephemeral_block_device {
          + device_name  = (known after apply)
          + no_device    = (known after apply)
          + virtual_name = (known after apply)
        }

      + metadata_options {
          + http_endpoint               = (known after apply)
          + http_put_response_hop_limit = (known after apply)
          + http_tokens                 = (known after apply)
        }

      + network_interface {
          + delete_on_termination = (known after apply)
          + device_index          = (known after apply)
          + network_interface_id  = (known after apply)
        }

      + root_block_device {
          + delete_on_termination = true
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = 8
          + volume_type           = (known after apply)
        }
    }

Plan: 1 to add, 0 to change, 0 to destroy.

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Note: You didn't use the -out option to save this plan, so Terraform can't guarantee to take exactly these actions
if you run "terraform apply" now.
learn-terraform-aws-instance $ terraform apply

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with
the following symbols:
  + create

Terraform will perform the following actions:

  # aws_instance.ubuntu_server will be created
  + resource "aws_instance" "ubuntu_server" {
      + ami                                  = "ami-029c64b3c205e6cce"
      + arn                                  = (known after apply)
      + associate_public_ip_address          = true
      + availability_zone                    = (known after apply)
      + cpu_core_count                       = (known after apply)
      + cpu_threads_per_core                 = (known after apply)
      + disable_api_termination              = (known after apply)
      + ebs_optimized                        = (known after apply)
      + get_password_data                    = false
      + host_id                              = (known after apply)
      + id                                   = (known after apply)
      + instance_initiated_shutdown_behavior = (known after apply)
      + instance_state                       = (known after apply)
      + instance_type                        = "t4g.micro"
      + ipv6_address_count                   = (known after apply)
      + ipv6_addresses                       = (known after apply)
      + key_name                             = "aws"
      + monitoring                           = (known after apply)
      + outpost_arn                          = (known after apply)
      + password_data                        = (known after apply)
      + placement_group                      = (known after apply)
      + primary_network_interface_id         = (known after apply)
      + private_dns                          = (known after apply)
      + private_ip                           = (known after apply)
      + public_dns                           = (known after apply)
      + public_ip                            = (known after apply)
      + secondary_private_ips                = (known after apply)
      + security_groups                      = (known after apply)
      + source_dest_check                    = true
      + subnet_id                            = (known after apply)
      + tags                                 = {
          + "Name" = "Ubuntu Server"
        }
      + tags_all                             = {
          + "Name" = "Ubuntu Server"
        }
      + tenancy                              = (known after apply)
      + user_data                            = (known after apply)
      + user_data_base64                     = (known after apply)
      + vpc_security_group_ids               = (known after apply)

      + capacity_reservation_specification {
          + capacity_reservation_preference = (known after apply)

          + capacity_reservation_target {
              + capacity_reservation_id = (known after apply)
            }
        }

      + ebs_block_device {
          + delete_on_termination = (known after apply)
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + snapshot_id           = (known after apply)
          + tags                  = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = (known after apply)
          + volume_type           = (known after apply)
        }

      + enclave_options {
          + enabled = (known after apply)
        }

      + ephemeral_block_device {
          + device_name  = (known after apply)
          + no_device    = (known after apply)
          + virtual_name = (known after apply)
        }

      + metadata_options {
          + http_endpoint               = (known after apply)
          + http_put_response_hop_limit = (known after apply)
          + http_tokens                 = (known after apply)
        }

      + network_interface {
          + delete_on_termination = (known after apply)
          + device_index          = (known after apply)
          + network_interface_id  = (known after apply)
        }

      + root_block_device {
          + delete_on_termination = true
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = 8
          + volume_type           = (known after apply)
        }
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

aws_instance.ubuntu_server: Creating...
aws_instance.ubuntu_server: Still creating... [10s elapsed]
aws_instance.ubuntu_server: Still creating... [20s elapsed]
aws_instance.ubuntu_server: Creation complete after 23s [id=i-0f6d5b2075b78a6ac]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
learn-terraform-aws-instance $ ssh -i ~/.ssh/aws ec2-user@52.90.174.136
^C
learn-terraform-aws-instance $ ssh -i v ~/.ssh/aws ec2-user@52.90.174.136
Warning: Identity file v not accessible: No such file or directory.
ssh: Could not resolve hostname /users/karuppiahn/.ssh/aws: nodename nor servname provided, or not known
learn-terraform-aws-instance $ ssh -i -vv ~/.ssh/aws ec2-user@52.90.174.136
Warning: Identity file -vv not accessible: No such file or directory.
ssh: Could not resolve hostname /users/karuppiahn/.ssh/aws: nodename nor servname provided, or not known
learn-terraform-aws-instance $ ssh -vv -i ~/.ssh/aws ec2-user@52.90.174.136
OpenSSH_8.1p1, LibreSSL 2.7.3
debug1: Reading configuration data /Users/karuppiahn/.ssh/config
debug1: Reading configuration data /etc/ssh/ssh_config
debug1: /etc/ssh/ssh_config line 47: Applying options for *
debug2: resolve_canonicalize: hostname 52.90.174.136 is address
debug2: ssh_connect_direct
debug1: Connecting to 52.90.174.136 [52.90.174.136] port 22.
^C
learn-terraform-aws-instance $ ssh -vv -i ~/.ssh/aws ec2-user@52.90.174.136
OpenSSH_8.1p1, LibreSSL 2.7.3
debug1: Reading configuration data /Users/karuppiahn/.ssh/config
debug1: Reading configuration data /etc/ssh/ssh_config
debug1: /etc/ssh/ssh_config line 47: Applying options for *
debug2: resolve_canonicalize: hostname 52.90.174.136 is address
debug2: ssh_connect_direct
debug1: Connecting to 52.90.174.136 [52.90.174.136] port 22.
^C
learn-terraform-aws-instance $ terraform destroy
aws_instance.ubuntu_server: Refreshing state... [id=i-0f6d5b2075b78a6ac]

Note: Objects have changed outside of Terraform

Terraform detected the following changes made outside of Terraform since the last "terraform apply":

  # aws_instance.ubuntu_server has been changed
  ~ resource "aws_instance" "ubuntu_server" {
        id                                   = "i-0f6d5b2075b78a6ac"
        tags                                 = {
            "Name" = "Ubuntu Server"
        }
        # (29 unchanged attributes hidden)




      ~ root_block_device {
          + tags                  = {}
            # (8 unchanged attributes hidden)
        }
        # (3 unchanged blocks hidden)
    }

Unless you have made equivalent changes to your configuration, or ignored the relevant attributes using
ignore_changes, the following plan may include actions to undo or respond to these changes.

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with
the following symbols:
  - destroy

Terraform will perform the following actions:

  # aws_instance.ubuntu_server will be destroyed
  - resource "aws_instance" "ubuntu_server" {
      - ami                                  = "ami-029c64b3c205e6cce" -> null
      - arn                                  = "arn:aws:ec2:us-east-1:469318448823:instance/i-0f6d5b2075b78a6ac" -> null
      - associate_public_ip_address          = true -> null
      - availability_zone                    = "us-east-1a" -> null
      - cpu_core_count                       = 2 -> null
      - cpu_threads_per_core                 = 1 -> null
      - disable_api_termination              = false -> null
      - ebs_optimized                        = false -> null
      - get_password_data                    = false -> null
      - hibernation                          = false -> null
      - id                                   = "i-0f6d5b2075b78a6ac" -> null
      - instance_initiated_shutdown_behavior = "stop" -> null
      - instance_state                       = "running" -> null
      - instance_type                        = "t4g.micro" -> null
      - ipv6_address_count                   = 0 -> null
      - ipv6_addresses                       = [] -> null
      - key_name                             = "aws" -> null
      - monitoring                           = false -> null
      - primary_network_interface_id         = "eni-033e7741431af5ec3" -> null
      - private_dns                          = "ip-172-31-51-153.ec2.internal" -> null
      - private_ip                           = "172.31.51.153" -> null
      - public_dns                           = "ec2-52-90-174-136.compute-1.amazonaws.com" -> null
      - public_ip                            = "52.90.174.136" -> null
      - secondary_private_ips                = [] -> null
      - security_groups                      = [
          - "default",
        ] -> null
      - source_dest_check                    = true -> null
      - subnet_id                            = "subnet-38bf0e12" -> null
      - tags                                 = {
          - "Name" = "Ubuntu Server"
        } -> null
      - tags_all                             = {
          - "Name" = "Ubuntu Server"
        } -> null
      - tenancy                              = "default" -> null
      - vpc_security_group_ids               = [
          - "sg-bba33fc3",
        ] -> null

      - capacity_reservation_specification {
          - capacity_reservation_preference = "open" -> null
        }

      - enclave_options {
          - enabled = false -> null
        }

      - metadata_options {
          - http_endpoint               = "enabled" -> null
          - http_put_response_hop_limit = 1 -> null
          - http_tokens                 = "optional" -> null
        }

      - root_block_device {
          - delete_on_termination = true -> null
          - device_name           = "/dev/xvda" -> null
          - encrypted             = false -> null
          - iops                  = 100 -> null
          - tags                  = {} -> null
          - throughput            = 0 -> null
          - volume_id             = "vol-0b83a3514d87be694" -> null
          - volume_size           = 8 -> null
          - volume_type           = "gp2" -> null
        }
    }

Plan: 0 to add, 0 to change, 1 to destroy.

Do you really want to destroy all resources?
  Terraform will destroy all your managed infrastructure, as shown above.
  There is no undo. Only 'yes' will be accepted to confirm.

  Enter a value: yes

aws_instance.ubuntu_server: Destroying... [id=i-0f6d5b2075b78a6ac]
aws_instance.ubuntu_server: Still destroying... [id=i-0f6d5b2075b78a6ac, 10s elapsed]
aws_instance.ubuntu_server: Still destroying... [id=i-0f6d5b2075b78a6ac, 20s elapsed]
aws_instance.ubuntu_server: Still destroying... [id=i-0f6d5b2075b78a6ac, 30s elapsed]
aws_instance.ubuntu_server: Still destroying... [id=i-0f6d5b2075b78a6ac, 40s elapsed]
aws_instance.ubuntu_server: Still destroying... [id=i-0f6d5b2075b78a6ac, 50s elapsed]
aws_instance.ubuntu_server: Destruction complete after 58s

Destroy complete! Resources: 1 destroyed.
learn-terraform-aws-instance $ gst
On branch master
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)

Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	new file:   .gitignore
	new file:   .terraform.lock.hcl
	new file:   STORY.md
	new file:   main.tf

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	modified:   main.tf

learn-terraform-aws-instance $ ga .
learn-terraform-aws-instance $ terraform plan

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with
the following symbols:
  + create

Terraform will perform the following actions:

  # aws_instance.ubuntu_server will be created
  + resource "aws_instance" "ubuntu_server" {
      + ami                                  = "ami-029c64b3c205e6cce"
      + arn                                  = (known after apply)
      + associate_public_ip_address          = true
      + availability_zone                    = (known after apply)
      + cpu_core_count                       = (known after apply)
      + cpu_threads_per_core                 = (known after apply)
      + disable_api_termination              = (known after apply)
      + ebs_optimized                        = (known after apply)
      + get_password_data                    = false
      + host_id                              = (known after apply)
      + id                                   = (known after apply)
      + instance_initiated_shutdown_behavior = (known after apply)
      + instance_state                       = (known after apply)
      + instance_type                        = "t4g.micro"
      + ipv6_address_count                   = (known after apply)
      + ipv6_addresses                       = (known after apply)
      + key_name                             = "aws"
      + monitoring                           = (known after apply)
      + outpost_arn                          = (known after apply)
      + password_data                        = (known after apply)
      + placement_group                      = (known after apply)
      + primary_network_interface_id         = (known after apply)
      + private_dns                          = (known after apply)
      + private_ip                           = (known after apply)
      + public_dns                           = (known after apply)
      + public_ip                            = (known after apply)
      + secondary_private_ips                = (known after apply)
      + security_groups                      = (known after apply)
      + source_dest_check                    = true
      + subnet_id                            = (known after apply)
      + tags                                 = {
          + "Name" = "Ubuntu Server"
        }
      + tags_all                             = {
          + "Name" = "Ubuntu Server"
        }
      + tenancy                              = (known after apply)
      + user_data                            = (known after apply)
      + user_data_base64                     = (known after apply)
      + vpc_security_group_ids               = [
          + "sg-0f60af9c0c3331a40",
        ]

      + capacity_reservation_specification {
          + capacity_reservation_preference = (known after apply)

          + capacity_reservation_target {
              + capacity_reservation_id = (known after apply)
            }
        }

      + ebs_block_device {
          + delete_on_termination = (known after apply)
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + snapshot_id           = (known after apply)
          + tags                  = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = (known after apply)
          + volume_type           = (known after apply)
        }

      + enclave_options {
          + enabled = (known after apply)
        }

      + ephemeral_block_device {
          + device_name  = (known after apply)
          + no_device    = (known after apply)
          + virtual_name = (known after apply)
        }

      + metadata_options {
          + http_endpoint               = (known after apply)
          + http_put_response_hop_limit = (known after apply)
          + http_tokens                 = (known after apply)
        }

      + network_interface {
          + delete_on_termination = (known after apply)
          + device_index          = (known after apply)
          + network_interface_id  = (known after apply)
        }

      + root_block_device {
          + delete_on_termination = true
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = 8
          + volume_type           = (known after apply)
        }
    }

Plan: 1 to add, 0 to change, 0 to destroy.

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Note: You didn't use the -out option to save this plan, so Terraform can't guarantee to take exactly these actions
if you run "terraform apply" now.
learn-terraform-aws-instance $ terraform plan -out
╷
│ Error: Failed to parse command-line flags
│ 
│ flag needs an argument: -out
╵

For more help on using this command, run:
  terraform plan -help
learn-terraform-aws-instance $ terraform plan 

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with
the following symbols:
  + create

Terraform will perform the following actions:

  # aws_instance.ubuntu_server will be created
  + resource "aws_instance" "ubuntu_server" {
      + ami                                  = "ami-029c64b3c205e6cce"
      + arn                                  = (known after apply)
      + associate_public_ip_address          = true
      + availability_zone                    = (known after apply)
      + cpu_core_count                       = (known after apply)
      + cpu_threads_per_core                 = (known after apply)
      + disable_api_termination              = (known after apply)
      + ebs_optimized                        = (known after apply)
      + get_password_data                    = false
      + host_id                              = (known after apply)
      + id                                   = (known after apply)
      + instance_initiated_shutdown_behavior = (known after apply)
      + instance_state                       = (known after apply)
      + instance_type                        = "t4g.micro"
      + ipv6_address_count                   = (known after apply)
      + ipv6_addresses                       = (known after apply)
      + key_name                             = "aws"
      + monitoring                           = (known after apply)
      + outpost_arn                          = (known after apply)
      + password_data                        = (known after apply)
      + placement_group                      = (known after apply)
      + primary_network_interface_id         = (known after apply)
      + private_dns                          = (known after apply)
      + private_ip                           = (known after apply)
      + public_dns                           = (known after apply)
      + public_ip                            = (known after apply)
      + secondary_private_ips                = (known after apply)
      + security_groups                      = (known after apply)
      + source_dest_check                    = true
      + subnet_id                            = (known after apply)
      + tags                                 = {
          + "Name" = "Ubuntu Server"
        }
      + tags_all                             = {
          + "Name" = "Ubuntu Server"
        }
      + tenancy                              = (known after apply)
      + user_data                            = (known after apply)
      + user_data_base64                     = (known after apply)
      + vpc_security_group_ids               = [
          + "sg-0f60af9c0c3331a40",
        ]

      + capacity_reservation_specification {
          + capacity_reservation_preference = (known after apply)

          + capacity_reservation_target {
              + capacity_reservation_id = (known after apply)
            }
        }

      + ebs_block_device {
          + delete_on_termination = (known after apply)
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + snapshot_id           = (known after apply)
          + tags                  = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = (known after apply)
          + volume_type           = (known after apply)
        }

      + enclave_options {
          + enabled = (known after apply)
        }

      + ephemeral_block_device {
          + device_name  = (known after apply)
          + no_device    = (known after apply)
          + virtual_name = (known after apply)
        }

      + metadata_options {
          + http_endpoint               = (known after apply)
          + http_put_response_hop_limit = (known after apply)
          + http_tokens                 = (known after apply)
        }

      + network_interface {
          + delete_on_termination = (known after apply)
          + device_index          = (known after apply)
          + network_interface_id  = (known after apply)
        }

      + root_block_device {
          + delete_on_termination = true
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = 8
          + volume_type           = (known after apply)
        }
    }

Plan: 1 to add, 0 to change, 0 to destroy.

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Note: You didn't use the -out option to save this plan, so Terraform can't guarantee to take exactly these actions
if you run "terraform apply" now.
learn-terraform-aws-instance $ terraform plan -out tf.plan

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with
the following symbols:
  + create

Terraform will perform the following actions:

  # aws_instance.ubuntu_server will be created
  + resource "aws_instance" "ubuntu_server" {
      + ami                                  = "ami-029c64b3c205e6cce"
      + arn                                  = (known after apply)
      + associate_public_ip_address          = true
      + availability_zone                    = (known after apply)
      + cpu_core_count                       = (known after apply)
      + cpu_threads_per_core                 = (known after apply)
      + disable_api_termination              = (known after apply)
      + ebs_optimized                        = (known after apply)
      + get_password_data                    = false
      + host_id                              = (known after apply)
      + id                                   = (known after apply)
      + instance_initiated_shutdown_behavior = (known after apply)
      + instance_state                       = (known after apply)
      + instance_type                        = "t4g.micro"
      + ipv6_address_count                   = (known after apply)
      + ipv6_addresses                       = (known after apply)
      + key_name                             = "aws"
      + monitoring                           = (known after apply)
      + outpost_arn                          = (known after apply)
      + password_data                        = (known after apply)
      + placement_group                      = (known after apply)
      + primary_network_interface_id         = (known after apply)
      + private_dns                          = (known after apply)
      + private_ip                           = (known after apply)
      + public_dns                           = (known after apply)
      + public_ip                            = (known after apply)
      + secondary_private_ips                = (known after apply)
      + security_groups                      = (known after apply)
      + source_dest_check                    = true
      + subnet_id                            = (known after apply)
      + tags                                 = {
          + "Name" = "Ubuntu Server"
        }
      + tags_all                             = {
          + "Name" = "Ubuntu Server"
        }
      + tenancy                              = (known after apply)
      + user_data                            = (known after apply)
      + user_data_base64                     = (known after apply)
      + vpc_security_group_ids               = [
          + "sg-0f60af9c0c3331a40",
        ]

      + capacity_reservation_specification {
          + capacity_reservation_preference = (known after apply)

          + capacity_reservation_target {
              + capacity_reservation_id = (known after apply)
            }
        }

      + ebs_block_device {
          + delete_on_termination = (known after apply)
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + snapshot_id           = (known after apply)
          + tags                  = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = (known after apply)
          + volume_type           = (known after apply)
        }

      + enclave_options {
          + enabled = (known after apply)
        }

      + ephemeral_block_device {
          + device_name  = (known after apply)
          + no_device    = (known after apply)
          + virtual_name = (known after apply)
        }

      + metadata_options {
          + http_endpoint               = (known after apply)
          + http_put_response_hop_limit = (known after apply)
          + http_tokens                 = (known after apply)
        }

      + network_interface {
          + delete_on_termination = (known after apply)
          + device_index          = (known after apply)
          + network_interface_id  = (known after apply)
        }

      + root_block_device {
          + delete_on_termination = true
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = 8
          + volume_type           = (known after apply)
        }
    }

Plan: 1 to add, 0 to change, 0 to destroy.

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Saved the plan to: tf.plan

To perform exactly these actions, run the following command to apply:
    terraform apply "tf.plan"
learn-terraform-aws-instance $ terraform apply "tf.plan"
aws_instance.ubuntu_server: Creating...
aws_instance.ubuntu_server: Still creating... [10s elapsed]
aws_instance.ubuntu_server: Still creating... [20s elapsed]
aws_instance.ubuntu_server: Creation complete after 23s [id=i-08be8d99df605a44c]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
learn-terraform-aws-instance $ ssh -vv -i ~/.ssh/aws ec2-user@54.208.139.42
OpenSSH_8.1p1, LibreSSL 2.7.3
debug1: Reading configuration data /Users/karuppiahn/.ssh/config
debug1: Reading configuration data /etc/ssh/ssh_config
debug1: /etc/ssh/ssh_config line 47: Applying options for *
debug2: resolve_canonicalize: hostname 54.208.139.42 is address
debug2: ssh_connect_direct
debug1: Connecting to 54.208.139.42 [54.208.139.42] port 22.
debug1: Connection established.
debug1: identity file /Users/karuppiahn/.ssh/aws type 0
debug1: identity file /Users/karuppiahn/.ssh/aws-cert type -1
debug1: Local version string SSH-2.0-OpenSSH_8.1
debug1: Remote protocol version 2.0, remote software version OpenSSH_7.4
debug1: match: OpenSSH_7.4 pat OpenSSH_7.0*,OpenSSH_7.1*,OpenSSH_7.2*,OpenSSH_7.3*,OpenSSH_7.4*,OpenSSH_7.5*,OpenSSH_7.6*,OpenSSH_7.7* compat 0x04000002
debug2: fd 3 setting O_NONBLOCK
debug1: Authenticating to 54.208.139.42:22 as 'ec2-user'
debug1: SSH2_MSG_KEXINIT sent
debug1: SSH2_MSG_KEXINIT received
debug2: local client KEXINIT proposal
debug2: KEX algorithms: curve25519-sha256,curve25519-sha256@libssh.org,ecdh-sha2-nistp256,ecdh-sha2-nistp384,ecdh-sha2-nistp521,diffie-hellman-group-exchange-sha256,diffie-hellman-group16-sha512,diffie-hellman-group18-sha512,diffie-hellman-group14-sha256,diffie-hellman-group14-sha1,ext-info-c
debug2: host key algorithms: ecdsa-sha2-nistp256-cert-v01@openssh.com,ecdsa-sha2-nistp384-cert-v01@openssh.com,ecdsa-sha2-nistp521-cert-v01@openssh.com,ssh-ed25519-cert-v01@openssh.com,rsa-sha2-512-cert-v01@openssh.com,rsa-sha2-256-cert-v01@openssh.com,ssh-rsa-cert-v01@openssh.com,ecdsa-sha2-nistp256,ecdsa-sha2-nistp384,ecdsa-sha2-nistp521,ssh-ed25519,rsa-sha2-512,rsa-sha2-256,ssh-rsa
debug2: ciphers ctos: chacha20-poly1305@openssh.com,aes128-ctr,aes192-ctr,aes256-ctr,aes128-gcm@openssh.com,aes256-gcm@openssh.com
debug2: ciphers stoc: chacha20-poly1305@openssh.com,aes128-ctr,aes192-ctr,aes256-ctr,aes128-gcm@openssh.com,aes256-gcm@openssh.com
debug2: MACs ctos: umac-64-etm@openssh.com,umac-128-etm@openssh.com,hmac-sha2-256-etm@openssh.com,hmac-sha2-512-etm@openssh.com,hmac-sha1-etm@openssh.com,umac-64@openssh.com,umac-128@openssh.com,hmac-sha2-256,hmac-sha2-512,hmac-sha1
debug2: MACs stoc: umac-64-etm@openssh.com,umac-128-etm@openssh.com,hmac-sha2-256-etm@openssh.com,hmac-sha2-512-etm@openssh.com,hmac-sha1-etm@openssh.com,umac-64@openssh.com,umac-128@openssh.com,hmac-sha2-256,hmac-sha2-512,hmac-sha1
debug2: compression ctos: none,zlib@openssh.com,zlib
debug2: compression stoc: none,zlib@openssh.com,zlib
debug2: languages ctos: 
debug2: languages stoc: 
debug2: first_kex_follows 0 
debug2: reserved 0 
debug2: peer server KEXINIT proposal
debug2: KEX algorithms: curve25519-sha256,curve25519-sha256@libssh.org,ecdh-sha2-nistp256,ecdh-sha2-nistp384,ecdh-sha2-nistp521,diffie-hellman-group-exchange-sha256,diffie-hellman-group16-sha512,diffie-hellman-group18-sha512,diffie-hellman-group-exchange-sha1,diffie-hellman-group14-sha256,diffie-hellman-group14-sha1,diffie-hellman-group1-sha1
debug2: host key algorithms: ssh-rsa,rsa-sha2-512,rsa-sha2-256,ecdsa-sha2-nistp256,ssh-ed25519
debug2: ciphers ctos: chacha20-poly1305@openssh.com,aes128-ctr,aes192-ctr,aes256-ctr,aes128-gcm@openssh.com,aes256-gcm@openssh.com,aes128-cbc,aes192-cbc,aes256-cbc,blowfish-cbc,cast128-cbc,3des-cbc
debug2: ciphers stoc: chacha20-poly1305@openssh.com,aes128-ctr,aes192-ctr,aes256-ctr,aes128-gcm@openssh.com,aes256-gcm@openssh.com,aes128-cbc,aes192-cbc,aes256-cbc,blowfish-cbc,cast128-cbc,3des-cbc
debug2: MACs ctos: umac-64-etm@openssh.com,umac-128-etm@openssh.com,hmac-sha2-256-etm@openssh.com,hmac-sha2-512-etm@openssh.com,hmac-sha1-etm@openssh.com,umac-64@openssh.com,umac-128@openssh.com,hmac-sha2-256,hmac-sha2-512,hmac-sha1
debug2: MACs stoc: umac-64-etm@openssh.com,umac-128-etm@openssh.com,hmac-sha2-256-etm@openssh.com,hmac-sha2-512-etm@openssh.com,hmac-sha1-etm@openssh.com,umac-64@openssh.com,umac-128@openssh.com,hmac-sha2-256,hmac-sha2-512,hmac-sha1
debug2: compression ctos: none,zlib@openssh.com
debug2: compression stoc: none,zlib@openssh.com
debug2: languages ctos: 
debug2: languages stoc: 
debug2: first_kex_follows 0 
debug2: reserved 0 
debug1: kex: algorithm: curve25519-sha256
debug1: kex: host key algorithm: ecdsa-sha2-nistp256
debug1: kex: server->client cipher: chacha20-poly1305@openssh.com MAC: <implicit> compression: none
debug1: kex: client->server cipher: chacha20-poly1305@openssh.com MAC: <implicit> compression: none
debug1: expecting SSH2_MSG_KEX_ECDH_REPLY
debug1: Server host key: ecdsa-sha2-nistp256 SHA256:TFhLrhBvsJVbwMQLH3CiBCPiWsmldYzOkag28U0pzoo
The authenticity of host '54.208.139.42 (54.208.139.42)' can't be established.
ECDSA key fingerprint is SHA256:TFhLrhBvsJVbwMQLH3CiBCPiWsmldYzOkag28U0pzoo.
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added '54.208.139.42' (ECDSA) to the list of known hosts.
debug2: set_newkeys: mode 1
debug1: rekey out after 134217728 blocks
debug1: SSH2_MSG_NEWKEYS sent
debug1: expecting SSH2_MSG_NEWKEYS
debug1: SSH2_MSG_NEWKEYS received
debug2: set_newkeys: mode 0
debug1: rekey in after 134217728 blocks
debug1: Will attempt key: karuppiahn@vmware.com ED25519 SHA256:3vku70losGmr1kmlRT52RuAG4e0IEDevQV+uOmCL3NI agent
debug1: Will attempt key: /Users/karuppiahn/.ssh/aws RSA SHA256:qImJtkIChdPcSAPsZkAhWxOfMBBJEIm4wQ/792Yf6NU explicit
debug2: pubkey_prepare: done
debug1: SSH2_MSG_EXT_INFO received
debug1: kex_input_ext_info: server-sig-algs=<rsa-sha2-256,rsa-sha2-512>
debug2: service_accept: ssh-userauth
debug1: SSH2_MSG_SERVICE_ACCEPT received
debug1: Authentications that can continue: publickey,gssapi-keyex,gssapi-with-mic
debug1: Next authentication method: publickey
debug1: Offering public key: karuppiahn@vmware.com ED25519 SHA256:3vku70losGmr1kmlRT52RuAG4e0IEDevQV+uOmCL3NI agent
debug2: we sent a publickey packet, wait for reply
debug1: Authentications that can continue: publickey,gssapi-keyex,gssapi-with-mic
debug1: Offering public key: /Users/karuppiahn/.ssh/aws RSA SHA256:qImJtkIChdPcSAPsZkAhWxOfMBBJEIm4wQ/792Yf6NU explicit
debug2: we sent a publickey packet, wait for reply
debug1: Server accepts key: /Users/karuppiahn/.ssh/aws RSA SHA256:qImJtkIChdPcSAPsZkAhWxOfMBBJEIm4wQ/792Yf6NU explicit
Enter passphrase for key '/Users/karuppiahn/.ssh/aws': 
debug1: Authentication succeeded (publickey).
Authenticated to 54.208.139.42 ([54.208.139.42]:22).
debug1: channel 0: new [client-session]
debug2: channel 0: send open
debug1: Requesting no-more-sessions@openssh.com
debug1: Entering interactive session.
debug1: pledge: network
debug1: client_input_global_request: rtype hostkeys-00@openssh.com want_reply 0
debug2: channel_input_open_confirmation: channel 0: callback start
debug2: fd 3 setting TCP_NODELAY
debug2: client_session2_setup: id 0
debug2: channel 0: request pty-req confirm 1
debug1: Sending environment.
debug1: Sending env LC_CTYPE = UTF-8
debug2: channel 0: request env confirm 0
debug2: channel 0: request shell confirm 1
debug2: channel_input_open_confirmation: channel 0: callback done
debug2: channel 0: open confirm rwindow 0 rmax 32768
debug2: channel_input_status_confirm: type 99 id 0
debug2: PTY allocation request accepted on channel 0
debug2: channel 0: rcvd adjust 2097152
debug2: channel_input_status_confirm: type 99 id 0
debug2: shell request accepted on channel 0

       __|  __|_  )
       _|  (     /   Amazon Linux 2 AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-2/
11 package(s) needed for security, out of 34 available
Run "sudo yum update" to apply all updates.
-bash: warning: setlocale: LC_CTYPE: cannot change locale (UTF-8): No such file or directory
[ec2-user@ip-172-31-49-123 ~]$ ls
[ec2-user@ip-172-31-49-123 ~]$ logout
debug2: channel 0: rcvd eof
debug2: channel 0: output open -> drain
debug2: channel 0: obuf empty
debug2: channel 0: chan_shutdown_write (i0 o1 sock -1 wfd 5 efd 6 [write])
debug2: channel 0: output drain -> closed
debug1: client_input_channel_req: channel 0 rtype exit-status reply 0
debug1: client_input_channel_req: channel 0 rtype eow@openssh.com reply 0
debug2: channel 0: rcvd eow
debug2: channel 0: chan_shutdown_read (i0 o3 sock -1 wfd 4 efd 6 [write])
debug2: channel 0: input open -> closed
debug2: channel 0: rcvd close
debug2: channel 0: almost dead
debug2: channel 0: gc: notify user
debug2: channel 0: gc: user detached
debug2: channel 0: send close
debug2: channel 0: is dead
debug2: channel 0: garbage collecting
debug1: channel 0: free: client-session, nchannels 1
Connection to 54.208.139.42 closed.
Transferred: sent 3272, received 3480 bytes, in 4.1 seconds
Bytes per second: sent 802.7, received 853.8
debug1: Exit status 0
learn-terraform-aws-instance $ terraform destroy
aws_instance.ubuntu_server: Refreshing state... [id=i-08be8d99df605a44c]

Note: Objects have changed outside of Terraform

Terraform detected the following changes made outside of Terraform since the last "terraform apply":

  # aws_instance.ubuntu_server has been changed
  ~ resource "aws_instance" "ubuntu_server" {
        id                                   = "i-08be8d99df605a44c"
        tags                                 = {
            "Name" = "Ubuntu Server"
        }
        # (29 unchanged attributes hidden)




      ~ root_block_device {
          + tags                  = {}
            # (8 unchanged attributes hidden)
        }
        # (3 unchanged blocks hidden)
    }

Unless you have made equivalent changes to your configuration, or ignored the relevant attributes using
ignore_changes, the following plan may include actions to undo or respond to these changes.

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with
the following symbols:
  - destroy

Terraform will perform the following actions:

  # aws_instance.ubuntu_server will be destroyed
  - resource "aws_instance" "ubuntu_server" {
      - ami                                  = "ami-029c64b3c205e6cce" -> null
      - arn                                  = "arn:aws:ec2:us-east-1:469318448823:instance/i-08be8d99df605a44c" -> null
      - associate_public_ip_address          = true -> null
      - availability_zone                    = "us-east-1a" -> null
      - cpu_core_count                       = 2 -> null
      - cpu_threads_per_core                 = 1 -> null
      - disable_api_termination              = false -> null
      - ebs_optimized                        = false -> null
      - get_password_data                    = false -> null
      - hibernation                          = false -> null
      - id                                   = "i-08be8d99df605a44c" -> null
      - instance_initiated_shutdown_behavior = "stop" -> null
      - instance_state                       = "running" -> null
      - instance_type                        = "t4g.micro" -> null
      - ipv6_address_count                   = 0 -> null
      - ipv6_addresses                       = [] -> null
      - key_name                             = "aws" -> null
      - monitoring                           = false -> null
      - primary_network_interface_id         = "eni-0421872a9e99aaf75" -> null
      - private_dns                          = "ip-172-31-49-123.ec2.internal" -> null
      - private_ip                           = "172.31.49.123" -> null
      - public_dns                           = "ec2-54-208-139-42.compute-1.amazonaws.com" -> null
      - public_ip                            = "54.208.139.42" -> null
      - secondary_private_ips                = [] -> null
      - security_groups                      = [
          - "launch-wizard-1",
        ] -> null
      - source_dest_check                    = true -> null
      - subnet_id                            = "subnet-38bf0e12" -> null
      - tags                                 = {
          - "Name" = "Ubuntu Server"
        } -> null
      - tags_all                             = {
          - "Name" = "Ubuntu Server"
        } -> null
      - tenancy                              = "default" -> null
      - vpc_security_group_ids               = [
          - "sg-0f60af9c0c3331a40",
        ] -> null

      - capacity_reservation_specification {
          - capacity_reservation_preference = "open" -> null
        }

      - enclave_options {
          - enabled = false -> null
        }

      - metadata_options {
          - http_endpoint               = "enabled" -> null
          - http_put_response_hop_limit = 1 -> null
          - http_tokens                 = "optional" -> null
        }

      - root_block_device {
          - delete_on_termination = true -> null
          - device_name           = "/dev/xvda" -> null
          - encrypted             = false -> null
          - iops                  = 100 -> null
          - tags                  = {} -> null
          - throughput            = 0 -> null
          - volume_id             = "vol-07a766297c1e932b6" -> null
          - volume_size           = 8 -> null
          - volume_type           = "gp2" -> null
        }
    }

Plan: 0 to add, 0 to change, 1 to destroy.

Do you really want to destroy all resources?
  Terraform will destroy all your managed infrastructure, as shown above.
  There is no undo. Only 'yes' will be accepted to confirm.

  Enter a value: yes

aws_instance.ubuntu_server: Destroying... [id=i-08be8d99df605a44c]
aws_instance.ubuntu_server: Still destroying... [id=i-08be8d99df605a44c, 10s elapsed]
aws_instance.ubuntu_server: Still destroying... [id=i-08be8d99df605a44c, 20s elapsed]
aws_instance.ubuntu_server: Still destroying... [id=i-08be8d99df605a44c, 30s elapsed]
aws_instance.ubuntu_server: Still destroying... [id=i-08be8d99df605a44c, 40s elapsed]
aws_instance.ubuntu_server: Still destroying... [id=i-08be8d99df605a44c, 50s elapsed]
aws_instance.ubuntu_server: Still destroying... [id=i-08be8d99df605a44c, 1m0s elapsed]
aws_instance.ubuntu_server: Destruction complete after 1m9s

Destroy complete! Resources: 1 destroyed.
learn-terraform-aws-instance $ 
```
