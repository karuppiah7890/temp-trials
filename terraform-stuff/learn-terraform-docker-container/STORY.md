
https://www.terraform.io

https://learn.hashicorp.com/terraform

https://learn.hashicorp.com/collections/terraform/aws-get-started

https://learn.hashicorp.com/tutorials/terraform/infrastructure-as-code?in=terraform/aws-get-started

https://learn.hashicorp.com/tutorials/terraform/install-cli?in=terraform/aws-get-started

https://learn.hashicorp.com/tutorials/terraform/aws-build?in=terraform/aws-get-started

---

https://spacelift.io/atlantis-alternative

https://spacelift.io/terraform-cloud-alternative

---

https://www.runatlantis.io

https://www.runatlantis.io/guide/

https://www.runatlantis.io/guide/#getting-started

---

```bash
learn-terraform-docker-container $ cat main.tf 
terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.13.0"
    }
  }
}

provider "docker" {}

resource "docker_image" "nginx" {
  name         = "nginx:latest"
  keep_locally = false
}

resource "docker_container" "nginx" {
  image = docker_image.nginx.latest
  name  = "tutorial"
  ports {
    internal = 80
    external = 8000
  }
}

learn-terraform-docker-container $ terraform init

Initializing the backend...

Initializing provider plugins...
- Finding kreuzwerker/docker versions matching "~> 2.13.0"...
- Installing kreuzwerker/docker v2.13.0...
- Installed kreuzwerker/docker v2.13.0 (self-signed, key ID 24E54F214569A8A5)

Partner and community providers are signed by their developers.
If you'd like to know more about provider signing, you can read about it here:
https://www.terraform.io/docs/cli/plugins/signing.html

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
learn-terraform-docker-container $ ls
main.tf
learn-terraform-docker-container $ ls -al
total 16
drwxr-xr-x  5 karuppiahn  staff   160 Sep 24 19:13 .
drwxr-xr-x  3 karuppiahn  staff    96 Sep 24 19:13 ..
drwxr-xr-x  3 karuppiahn  staff    96 Sep 24 19:13 .terraform
-rw-r--r--  1 karuppiahn  staff  1264 Sep 24 19:13 .terraform.lock.hcl
-rw-r--r--  1 karuppiahn  staff   392 Sep 24 19:11 main.tf
learn-terraform-docker-container $ cat .terraform
cat: .terraform: Is a directory
learn-terraform-docker-container $ cat .terraform/providers/registry.terraform.io/kreuzwerker/docker/2.13.0/darwin_amd64/
CHANGELOG.md                       README.md                          
LICENSE                            terraform-provider-docker_v2.13.0  
learn-terraform-docker-container $ cat .terraform/providers/registry.terraform.io/kreuzwerker/docker/2.13.0/darwin_amd64/
CHANGELOG.md                       README.md                          
LICENSE                            terraform-provider-docker_v2.13.0  
learn-terraform-docker-container $ cat .terraform/providers/registry.terraform.io/kreuzwerker/docker/2.13.0/darwin_amd64/
learn-terraform-docker-container $ ./.terraform/providers/registry.terraform.io/kreuzwerker/docker/2.13.0/darwin_amd64/terraform-provider-docker_v2.13.0 
This binary is a plugin. These are not meant to be executed directly.
Please execute the program that consumes these plugins, which will
load any plugins automatically
learn-terraform-docker-container $ terraform plan
╷
│ Error: Error pinging Docker server: Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
│ 
│   with provider["registry.terraform.io/kreuzwerker/docker"],
│   on main.tf line 10, in provider "docker":
│   10: provider "docker" {}
│ 
╵
learn-terraform-docker-container $ docker ps
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
learn-terraform-docker-container $ docker ps
Error response from daemon: dial unix docker.raw.sock: connect: connection refused
learn-terraform-docker-container $ docker ps
CONTAINER ID   IMAGE        COMMAND                  CREATED      STATUS         PORTS                                       NAMES
68d7cc6bf458   registry:2   "/entrypoint.sh /etc…"   4 days ago   Up 3 seconds   0.0.0.0:5000->5000/tcp, :::5000->5000/tcp   kind-registry
learn-terraform-docker-container $ terraform plan

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with
the following symbols:
  + create

Terraform will perform the following actions:

  # docker_container.nginx will be created
  + resource "docker_container" "nginx" {
      + attach           = false
      + bridge           = (known after apply)
      + command          = (known after apply)
      + container_logs   = (known after apply)
      + entrypoint       = (known after apply)
      + env              = (known after apply)
      + exit_code        = (known after apply)
      + gateway          = (known after apply)
      + hostname         = (known after apply)
      + id               = (known after apply)
      + image            = (known after apply)
      + init             = (known after apply)
      + ip_address       = (known after apply)
      + ip_prefix_length = (known after apply)
      + ipc_mode         = (known after apply)
      + log_driver       = "json-file"
      + logs             = false
      + must_run         = true
      + name             = "tutorial"
      + network_data     = (known after apply)
      + read_only        = false
      + remove_volumes   = true
      + restart          = "no"
      + rm               = false
      + security_opts    = (known after apply)
      + shm_size         = (known after apply)
      + start            = true
      + stdin_open       = false
      + tty              = false

      + healthcheck {
          + interval     = (known after apply)
          + retries      = (known after apply)
          + start_period = (known after apply)
          + test         = (known after apply)
          + timeout      = (known after apply)
        }

      + labels {
          + label = (known after apply)
          + value = (known after apply)
        }

      + ports {
          + external = 8000
          + internal = 80
          + ip       = "0.0.0.0"
          + protocol = "tcp"
        }
    }

  # docker_image.nginx will be created
  + resource "docker_image" "nginx" {
      + id           = (known after apply)
      + keep_locally = false
      + latest       = (known after apply)
      + name         = "nginx:latest"
      + output       = (known after apply)
      + repo_digest  = (known after apply)
    }

Plan: 2 to add, 0 to change, 0 to destroy.

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Note: You didn't use the -out option to save this plan, so Terraform can't guarantee to take exactly these actions
if you run "terraform apply" now.
learn-terraform-docker-container $ terraform apply

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with
the following symbols:
  + create

Terraform will perform the following actions:

  # docker_container.nginx will be created
  + resource "docker_container" "nginx" {
      + attach           = false
      + bridge           = (known after apply)
      + command          = (known after apply)
      + container_logs   = (known after apply)
      + entrypoint       = (known after apply)
      + env              = (known after apply)
      + exit_code        = (known after apply)
      + gateway          = (known after apply)
      + hostname         = (known after apply)
      + id               = (known after apply)
      + image            = (known after apply)
      + init             = (known after apply)
      + ip_address       = (known after apply)
      + ip_prefix_length = (known after apply)
      + ipc_mode         = (known after apply)
      + log_driver       = "json-file"
      + logs             = false
      + must_run         = true
      + name             = "tutorial"
      + network_data     = (known after apply)
      + read_only        = false
      + remove_volumes   = true
      + restart          = "no"
      + rm               = false
      + security_opts    = (known after apply)
      + shm_size         = (known after apply)
      + start            = true
      + stdin_open       = false
      + tty              = false

      + healthcheck {
          + interval     = (known after apply)
          + retries      = (known after apply)
          + start_period = (known after apply)
          + test         = (known after apply)
          + timeout      = (known after apply)
        }

      + labels {
          + label = (known after apply)
          + value = (known after apply)
        }

      + ports {
          + external = 8000
          + internal = 80
          + ip       = "0.0.0.0"
          + protocol = "tcp"
        }
    }

  # docker_image.nginx will be created
  + resource "docker_image" "nginx" {
      + id           = (known after apply)
      + keep_locally = false
      + latest       = (known after apply)
      + name         = "nginx:latest"
      + output       = (known after apply)
      + repo_digest  = (known after apply)
    }

Plan: 2 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

docker_image.nginx: Creating...
docker_image.nginx: Still creating... [10s elapsed]
docker_image.nginx: Creation complete after 12s [id=sha256:ad4c705f24d392b982b2f0747704b1c5162e45674294d5640cca7076eba2865dnginx:latest]
docker_container.nginx: Creating...
docker_container.nginx: Creation complete after 1s [id=3c61919a1975f37732eac15b8bd1ad7a8597b79de53f78b7c7235e57dac9df10]

Apply complete! Resources: 2 added, 0 changed, 0 destroyed.
learn-terraform-docker-container $ terraform plan
docker_image.nginx: Refreshing state... [id=sha256:ad4c705f24d392b982b2f0747704b1c5162e45674294d5640cca7076eba2865dnginx:latest]
docker_container.nginx: Refreshing state... [id=3c61919a1975f37732eac15b8bd1ad7a8597b79de53f78b7c7235e57dac9df10]

Note: Objects have changed outside of Terraform

Terraform detected the following changes made outside of Terraform since the last "terraform apply":

  # docker_container.nginx has been changed
  ~ resource "docker_container" "nginx" {
      + dns               = []
      + dns_opts          = []
      + dns_search        = []
      + group_add         = []
        id                = "3c61919a1975f37732eac15b8bd1ad7a8597b79de53f78b7c7235e57dac9df10"
      + links             = []
      + log_opts          = {}
        name              = "tutorial"
      + sysctls           = {}
      + tmpfs             = {}
        # (31 unchanged attributes hidden)

        # (1 unchanged block hidden)
    }

Unless you have made equivalent changes to your configuration, or ignored the relevant attributes using
ignore_changes, the following plan may include actions to undo or respond to these changes.

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

No changes. Your infrastructure matches the configuration.

Your configuration already matches the changes detected above. If you'd like to update the Terraform state to match,
create and apply a refresh-only plan:
  terraform apply -refresh-only
learn-terraform-docker-container $ terraform plan
docker_image.nginx: Refreshing state... [id=sha256:ad4c705f24d392b982b2f0747704b1c5162e45674294d5640cca7076eba2865dnginx:latest]
docker_container.nginx: Refreshing state... [id=3c61919a1975f37732eac15b8bd1ad7a8597b79de53f78b7c7235e57dac9df10]

Note: Objects have changed outside of Terraform

Terraform detected the following changes made outside of Terraform since the last "terraform apply":

  # docker_container.nginx has been changed
  ~ resource "docker_container" "nginx" {
      + dns               = []
      + dns_opts          = []
      + dns_search        = []
      + group_add         = []
        id                = "3c61919a1975f37732eac15b8bd1ad7a8597b79de53f78b7c7235e57dac9df10"
      + links             = []
      + log_opts          = {}
        name              = "tutorial"
      + sysctls           = {}
      + tmpfs             = {}
        # (31 unchanged attributes hidden)

        # (1 unchanged block hidden)
    }

Unless you have made equivalent changes to your configuration, or ignored the relevant attributes using
ignore_changes, the following plan may include actions to undo or respond to these changes.

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

No changes. Your infrastructure matches the configuration.

Your configuration already matches the changes detected above. If you'd like to update the Terraform state to match,
create and apply a refresh-only plan:
  terraform apply -refresh-only
learn-terraform-docker-container $ docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED              STATUS              PORTS                                       NAMES
3c61919a1975   ad4c705f24d3   "/docker-entrypoint.…"   About a minute ago   Up About a minute   0.0.0.0:8000->80/tcp                        tutorial
68d7cc6bf458   registry:2     "/entrypoint.sh /etc…"   4 days ago           Up 2 minutes        0.0.0.0:5000->5000/tcp, :::5000->5000/tcp   kind-registry
learn-terraform-docker-container $ curl localhost:8000
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
learn-terraform-docker-container $ terraform destroy
docker_image.nginx: Refreshing state... [id=sha256:ad4c705f24d392b982b2f0747704b1c5162e45674294d5640cca7076eba2865dnginx:latest]
docker_container.nginx: Refreshing state... [id=3c61919a1975f37732eac15b8bd1ad7a8597b79de53f78b7c7235e57dac9df10]

Note: Objects have changed outside of Terraform

Terraform detected the following changes made outside of Terraform since the last "terraform apply":

  # docker_container.nginx has been changed
  ~ resource "docker_container" "nginx" {
      + dns               = []
      + dns_opts          = []
      + dns_search        = []
      + group_add         = []
        id                = "3c61919a1975f37732eac15b8bd1ad7a8597b79de53f78b7c7235e57dac9df10"
      + links             = []
      + log_opts          = {}
        name              = "tutorial"
      + sysctls           = {}
      + tmpfs             = {}
        # (31 unchanged attributes hidden)

        # (1 unchanged block hidden)
    }

Unless you have made equivalent changes to your configuration, or ignored the relevant attributes using
ignore_changes, the following plan may include actions to undo or respond to these changes.

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with
the following symbols:
  - destroy

Terraform will perform the following actions:

  # docker_container.nginx will be destroyed
  - resource "docker_container" "nginx" {
      - attach            = false -> null
      - command           = [
          - "nginx",
          - "-g",
          - "daemon off;",
        ] -> null
      - cpu_shares        = 0 -> null
      - dns               = [] -> null
      - dns_opts          = [] -> null
      - dns_search        = [] -> null
      - entrypoint        = [
          - "/docker-entrypoint.sh",
        ] -> null
      - env               = [] -> null
      - gateway           = "172.17.0.1" -> null
      - group_add         = [] -> null
      - hostname          = "3c61919a1975" -> null
      - id                = "3c61919a1975f37732eac15b8bd1ad7a8597b79de53f78b7c7235e57dac9df10" -> null
      - image             = "sha256:ad4c705f24d392b982b2f0747704b1c5162e45674294d5640cca7076eba2865d" -> null
      - init              = false -> null
      - ip_address        = "172.17.0.3" -> null
      - ip_prefix_length  = 16 -> null
      - ipc_mode          = "private" -> null
      - links             = [] -> null
      - log_driver        = "json-file" -> null
      - log_opts          = {} -> null
      - logs              = false -> null
      - max_retry_count   = 0 -> null
      - memory            = 0 -> null
      - memory_swap       = 0 -> null
      - must_run          = true -> null
      - name              = "tutorial" -> null
      - network_data      = [
          - {
              - gateway                   = "172.17.0.1"
              - global_ipv6_address       = ""
              - global_ipv6_prefix_length = 0
              - ip_address                = "172.17.0.3"
              - ip_prefix_length          = 16
              - ipv6_gateway              = ""
              - network_name              = "bridge"
            },
        ] -> null
      - network_mode      = "default" -> null
      - privileged        = false -> null
      - publish_all_ports = false -> null
      - read_only         = false -> null
      - remove_volumes    = true -> null
      - restart           = "no" -> null
      - rm                = false -> null
      - security_opts     = [] -> null
      - shm_size          = 64 -> null
      - start             = true -> null
      - stdin_open        = false -> null
      - sysctls           = {} -> null
      - tmpfs             = {} -> null
      - tty               = false -> null

      - ports {
          - external = 8000 -> null
          - internal = 80 -> null
          - ip       = "0.0.0.0" -> null
          - protocol = "tcp" -> null
        }
    }

  # docker_image.nginx will be destroyed
  - resource "docker_image" "nginx" {
      - id           = "sha256:ad4c705f24d392b982b2f0747704b1c5162e45674294d5640cca7076eba2865dnginx:latest" -> null
      - keep_locally = false -> null
      - latest       = "sha256:ad4c705f24d392b982b2f0747704b1c5162e45674294d5640cca7076eba2865d" -> null
      - name         = "nginx:latest" -> null
      - repo_digest  = "nginx@sha256:853b221d3341add7aaadf5f81dd088ea943ab9c918766e295321294b035f3f3e" -> null
    }

Plan: 0 to add, 0 to change, 2 to destroy.

Do you really want to destroy all resources?
  Terraform will destroy all your managed infrastructure, as shown above.
  There is no undo. Only 'yes' will be accepted to confirm.

  Enter a value: yes

docker_container.nginx: Destroying... [id=3c61919a1975f37732eac15b8bd1ad7a8597b79de53f78b7c7235e57dac9df10]
docker_container.nginx: Destruction complete after 0s
docker_image.nginx: Destroying... [id=sha256:ad4c705f24d392b982b2f0747704b1c5162e45674294d5640cca7076eba2865dnginx:latest]
docker_image.nginx: Destruction complete after 1s

Destroy complete! Resources: 2 destroyed.
learn-terraform-docker-container $ docker ps 
CONTAINER ID   IMAGE        COMMAND                  CREATED      STATUS         PORTS                                       NAMES
68d7cc6bf458   registry:2   "/entrypoint.sh /etc…"   4 days ago   Up 2 minutes   0.0.0.0:5000->5000/tcp, :::5000->5000/tcp   kind-registry
learn-terraform-docker-container $ docker stop kind-registry
kind-registry
learn-terraform-docker-container $ docker ps 
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
learn-terraform-docker-container $ code .
learn-terraform-docker-container $ 
```



