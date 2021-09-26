https://duckduckgo.com/?t=ffab&q=packer+hashicorp&ia=web

https://www.packer.io/

https://learn.hashicorp.com/packer

https://learn.hashicorp.com/collections/packer/aws-get-started

https://duckduckgo.com/?t=ffab&q=hashicorp+configuration+language+vscode+extensions&ia=web

https://www.hashicorp.com/blog/supporting-the-hashicorp-terraform-extension-for-visual-studio-code

https://github.com/hashicorp/vscode-terraform/

https://marketplace.visualstudio.com/items?itemName=HashiCorp.terraform

https://www.packer.io/docs/builders

https://www.packer.io/docs/builders/amazon

https://www.packer.io/docs/builders/amazon/chroot

https://www.packer.io/docs/builders/amazon/instance

https://www.packer.io/docs/builders/amazon/ebs

https://www.packer.io/docs/builders/amazon/ebssurrogate

https://www.packer.io/docs/builders/amazon/ebsvolume

https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ComponentsAMIs.html#storage-for-the-root-device

https://duckduckgo.com/?t=ffab&q=amazon+machine+image&ia=web

https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AMIs.html

https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AMIs.html#creating-an-ami

https://aws.amazon.com/ec2/features/

https://aws.amazon.com/marketplace/pp/prodview-lnik5bngmt4mm

https://aws.amazon.com/marketplace/search/results?searchTerms=amazon+linux+

https://aws.amazon.com/marketplace/search/results?searchTerms=amazon+linux+&AMI_ARCHITECTURE=arm64&filters=AMI_ARCHITECTURE

https://aws.amazon.com/marketplace/pp/prodview-ps24xi3teh36o

https://www.toptal.com/developers/gitignore/api/packer

```bash
packer_tutorial $ packer init .
Installed plugin github.com/hashicorp/amazon v1.0.1 in "/Users/karuppiahn/.config/packer/plugins/github.com/hashicorp/amazon/packer-plugin-amazon_v1.0.1_x5.0_darwin_amd64"
packer_tutorial $
```

```bash
packer_tutorial $ packer fmt .
packer_tutorial $ packer validate .
packer_tutorial $ source .env
packer_tutorial $ packer build aws-ubuntu.pkr.hcl
learn-packer.amazon-ebs.ubuntu: output will be in this color.

==> learn-packer.amazon-ebs.ubuntu: Prevalidating any provided VPC information
==> learn-packer.amazon-ebs.ubuntu: Prevalidating AMI Name: learn-packer-linux-aws
    learn-packer.amazon-ebs.ubuntu: Found Image ID: ami-079e7a3f57cc8e0d0
==> learn-packer.amazon-ebs.ubuntu: Creating temporary keypair: packer_61504688-8e04-859a-9e75-5289f04313b7
==> learn-packer.amazon-ebs.ubuntu: Creating temporary security group for this instance: packer_6150468e-06ae-e5df-5350-96411f3989a5
==> learn-packer.amazon-ebs.ubuntu: Authorizing access to port 22 from [0.0.0.0/0] in the temporary security groups...
==> learn-packer.amazon-ebs.ubuntu: Launching a source AWS instance...
==> learn-packer.amazon-ebs.ubuntu: Adding tags to source instance
    learn-packer.amazon-ebs.ubuntu: Adding tag: "Name": "Packer Builder"
    learn-packer.amazon-ebs.ubuntu: Instance ID: i-066ce10a0b1c4d331
==> learn-packer.amazon-ebs.ubuntu: Waiting for instance (i-066ce10a0b1c4d331) to become ready...
==> learn-packer.amazon-ebs.ubuntu: Using SSH communicator to connect: 34.222.245.109
==> learn-packer.amazon-ebs.ubuntu: Waiting for SSH to become available...
==> learn-packer.amazon-ebs.ubuntu: Connected to SSH!
==> learn-packer.amazon-ebs.ubuntu: Stopping the source instance...
    learn-packer.amazon-ebs.ubuntu: Stopping instance
==> learn-packer.amazon-ebs.ubuntu: Waiting for the instance to stop...
==> learn-packer.amazon-ebs.ubuntu: Creating AMI learn-packer-linux-aws from instance i-066ce10a0b1c4d331
    learn-packer.amazon-ebs.ubuntu: AMI: ami-0d4e27e38e0c366fd
==> learn-packer.amazon-ebs.ubuntu: Waiting for AMI to become ready...
==> learn-packer.amazon-ebs.ubuntu: Skipping Enable AMI deprecation...
==> learn-packer.amazon-ebs.ubuntu: Terminating the source AWS instance...
==> learn-packer.amazon-ebs.ubuntu: Cleaning up any extra volumes...
==> learn-packer.amazon-ebs.ubuntu: No volumes to clean up, skipping
==> learn-packer.amazon-ebs.ubuntu: Deleting temporary security group...
==> learn-packer.amazon-ebs.ubuntu: Deleting temporary keypair...
Build 'learn-packer.amazon-ebs.ubuntu' finished after 4 minutes 15 seconds.

==> Wait completed after 4 minutes 15 seconds

==> Builds finished. The artifacts of successful builds are:
--> learn-packer.amazon-ebs.ubuntu: AMIs were created:
us-west-2: ami-0d4e27e38e0c366fd

packer_tutorial $
```

Deleting AMIs and the corresponding snapshots -
- https://us-west-2.console.aws.amazon.com/ec2/v2/home?region=us-west-2#Images:visibility=owned-by-me;sort=name
- https://us-west-2.console.aws.amazon.com/ec2/v2/home?region=us-west-2#Images:visibility=owned-by-me;search=learn-packer-linux-aws;sort=name
- https://us-west-2.console.aws.amazon.com/ec2/v2/home?region=us-west-2#Snapshots:visibility=owned-by-me;sort=tag:Name

---

https://learn.hashicorp.com/tutorials/packer/aws-get-started-provision?in=packer/aws-get-started

https://www.packer.io/docs/provisioners

https://www.packer.io/docs/provisioners/shell

https://www.packer.io/docs/provisioners/shell-local

https://www.packer.io/docs/provisioners/file

https://learn.hashicorp.com/tutorials/terraform/packer

```bash
packer_tutorial $ packer build aws-ubuntu.pkr.hcl
learn-packer.amazon-ebs.ubuntu: output will be in this color.

==> learn-packer.amazon-ebs.ubuntu: Prevalidating any provided VPC information
==> learn-packer.amazon-ebs.ubuntu: Prevalidating AMI Name: learn-packer-linux-aws
    learn-packer.amazon-ebs.ubuntu: Found Image ID: ami-079e7a3f57cc8e0d0
==> learn-packer.amazon-ebs.ubuntu: Creating temporary keypair: packer_61504adf-32b0-5178-e3c8-4e0c46831202
==> learn-packer.amazon-ebs.ubuntu: Creating temporary security group for this instance: packer_61504ae5-7f3d-6b1b-a184-538b394c6abd
==> learn-packer.amazon-ebs.ubuntu: Authorizing access to port 22 from [0.0.0.0/0] in the temporary security groups...
==> learn-packer.amazon-ebs.ubuntu: Launching a source AWS instance...
==> learn-packer.amazon-ebs.ubuntu: Adding tags to source instance
    learn-packer.amazon-ebs.ubuntu: Adding tag: "Name": "Packer Builder"
    learn-packer.amazon-ebs.ubuntu: Instance ID: i-04dc0273af6fc9c08
==> learn-packer.amazon-ebs.ubuntu: Waiting for instance (i-04dc0273af6fc9c08) to become ready...
==> learn-packer.amazon-ebs.ubuntu: Using SSH communicator to connect: 54.184.166.102
==> learn-packer.amazon-ebs.ubuntu: Waiting for SSH to become available...
==> learn-packer.amazon-ebs.ubuntu: Connected to SSH!
==> learn-packer.amazon-ebs.ubuntu: Provisioning with shell script: /var/folders/4z/09jpfvfj6c19lxl7ch78pzvc0000gn/T/packer-shell142153000
    learn-packer.amazon-ebs.ubuntu: Installing Redis
    learn-packer.amazon-ebs.ubuntu: Hit:1 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial InRelease
    learn-packer.amazon-ebs.ubuntu: Get:2 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    learn-packer.amazon-ebs.ubuntu: Get:3 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    learn-packer.amazon-ebs.ubuntu: Get:4 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial/universe amd64 Packages [7,532 kB]
    learn-packer.amazon-ebs.ubuntu: Get:5 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    learn-packer.amazon-ebs.ubuntu: Get:6 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial/universe Translation-en [4,354 kB]
    learn-packer.amazon-ebs.ubuntu: Get:7 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial/multiverse amd64 Packages [144 kB]
    learn-packer.amazon-ebs.ubuntu: Get:8 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial/multiverse Translation-en [106 kB]
    learn-packer.amazon-ebs.ubuntu: Get:9 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [2,049 kB]
    learn-packer.amazon-ebs.ubuntu: Get:10 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [482 kB]
    learn-packer.amazon-ebs.ubuntu: Get:11 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [1,219 kB]
    learn-packer.amazon-ebs.ubuntu: Get:12 https://esm.ubuntu.com/infra/ubuntu xenial-infra-security InRelease [7,509 B]
    learn-packer.amazon-ebs.ubuntu: Get:13 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [358 kB]
    learn-packer.amazon-ebs.ubuntu: Get:14 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-updates/multiverse amd64 Packages [22.6 kB]
    learn-packer.amazon-ebs.ubuntu: Get:15 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-updates/multiverse Translation-en [8,476 B]
    learn-packer.amazon-ebs.ubuntu: Get:16 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-backports/main amd64 Packages [9,812 B]
    learn-packer.amazon-ebs.ubuntu: Get:17 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-backports/main Translation-en [4,456 B]
    learn-packer.amazon-ebs.ubuntu: Get:18 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-backports/universe amd64 Packages [11.3 kB]
    learn-packer.amazon-ebs.ubuntu: Get:19 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-backports/universe Translation-en [4,476 B]
    learn-packer.amazon-ebs.ubuntu: Get:20 https://esm.ubuntu.com/infra/ubuntu xenial-infra-updates InRelease [7,475 B]
    learn-packer.amazon-ebs.ubuntu: Get:21 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [1,648 kB]
    learn-packer.amazon-ebs.ubuntu: Get:22 https://esm.ubuntu.com/infra/ubuntu xenial-infra-security/main amd64 Packages [182 kB]
    learn-packer.amazon-ebs.ubuntu: Get:23 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [785 kB]
    learn-packer.amazon-ebs.ubuntu: Get:24 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [225 kB]
    learn-packer.amazon-ebs.ubuntu: Get:25 http://security.ubuntu.com/ubuntu xenial-security/multiverse amd64 Packages [7,864 B]
    learn-packer.amazon-ebs.ubuntu: Get:26 http://security.ubuntu.com/ubuntu xenial-security/multiverse Translation-en [2,672 B]
    learn-packer.amazon-ebs.ubuntu: Fetched 19.5 MB in 3s (4,994 kB/s)
    learn-packer.amazon-ebs.ubuntu: Reading package lists...
    learn-packer.amazon-ebs.ubuntu: Reading package lists...
    learn-packer.amazon-ebs.ubuntu: Building dependency tree...
    learn-packer.amazon-ebs.ubuntu: Reading state information...
    learn-packer.amazon-ebs.ubuntu: The following additional packages will be installed:
    learn-packer.amazon-ebs.ubuntu:   libjemalloc1 redis-tools
    learn-packer.amazon-ebs.ubuntu: Suggested packages:
    learn-packer.amazon-ebs.ubuntu:   ruby-redis
    learn-packer.amazon-ebs.ubuntu: The following NEW packages will be installed:
    learn-packer.amazon-ebs.ubuntu:   libjemalloc1 redis-server redis-tools
    learn-packer.amazon-ebs.ubuntu: 0 upgraded, 3 newly installed, 0 to remove and 6 not upgraded.
    learn-packer.amazon-ebs.ubuntu: Need to get 519 kB of archives.
    learn-packer.amazon-ebs.ubuntu: After this operation, 1,507 kB of additional disk space will be used.
    learn-packer.amazon-ebs.ubuntu: Get:1 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial/universe amd64 libjemalloc1 amd64 3.6.0-9ubuntu1 [78.9 kB]
    learn-packer.amazon-ebs.ubuntu: Get:2 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-updates/universe amd64 redis-tools amd64 2:3.0.6-1ubuntu0.4 [95.5 kB]
    learn-packer.amazon-ebs.ubuntu: Get:3 http://us-west-2.ec2.archive.ubuntu.com/ubuntu xenial-updates/universe amd64 redis-server amd64 2:3.0.6-1ubuntu0.4 [344 kB]
==> learn-packer.amazon-ebs.ubuntu: debconf: unable to initialize frontend: Dialog
==> learn-packer.amazon-ebs.ubuntu: debconf: (Dialog frontend will not work on a dumb terminal, an emacs shell buffer, or without a controlling terminal.)
==> learn-packer.amazon-ebs.ubuntu: debconf: falling back to frontend: Readline
==> learn-packer.amazon-ebs.ubuntu: debconf: unable to initialize frontend: Readline
==> learn-packer.amazon-ebs.ubuntu: debconf: (This frontend requires a controlling tty.)
==> learn-packer.amazon-ebs.ubuntu: debconf: falling back to frontend: Teletype
==> learn-packer.amazon-ebs.ubuntu: dpkg-preconfigure: unable to re-open stdin:
    learn-packer.amazon-ebs.ubuntu: Fetched 519 kB in 0s (18.3 MB/s)
    learn-packer.amazon-ebs.ubuntu: Selecting previously unselected package libjemalloc1.
    learn-packer.amazon-ebs.ubuntu: (Reading database ... 51558 files and directories currently installed.)
    learn-packer.amazon-ebs.ubuntu: Preparing to unpack .../libjemalloc1_3.6.0-9ubuntu1_amd64.deb ...
    learn-packer.amazon-ebs.ubuntu: Unpacking libjemalloc1 (3.6.0-9ubuntu1) ...
    learn-packer.amazon-ebs.ubuntu: Selecting previously unselected package redis-tools.
    learn-packer.amazon-ebs.ubuntu: Preparing to unpack .../redis-tools_2%3a3.0.6-1ubuntu0.4_amd64.deb ...
    learn-packer.amazon-ebs.ubuntu: Unpacking redis-tools (2:3.0.6-1ubuntu0.4) ...
    learn-packer.amazon-ebs.ubuntu: Selecting previously unselected package redis-server.
    learn-packer.amazon-ebs.ubuntu: Preparing to unpack .../redis-server_2%3a3.0.6-1ubuntu0.4_amd64.deb ...
    learn-packer.amazon-ebs.ubuntu: Unpacking redis-server (2:3.0.6-1ubuntu0.4) ...
    learn-packer.amazon-ebs.ubuntu: Processing triggers for libc-bin (2.23-0ubuntu11.3) ...
    learn-packer.amazon-ebs.ubuntu: Processing triggers for man-db (2.7.5-1) ...
    learn-packer.amazon-ebs.ubuntu: Processing triggers for ureadahead (0.100.0-19.1) ...
    learn-packer.amazon-ebs.ubuntu: Processing triggers for systemd (229-4ubuntu21.31) ...
    learn-packer.amazon-ebs.ubuntu: Setting up libjemalloc1 (3.6.0-9ubuntu1) ...
    learn-packer.amazon-ebs.ubuntu: Setting up redis-tools (2:3.0.6-1ubuntu0.4) ...
    learn-packer.amazon-ebs.ubuntu: Setting up redis-server (2:3.0.6-1ubuntu0.4) ...
    learn-packer.amazon-ebs.ubuntu: Processing triggers for libc-bin (2.23-0ubuntu11.3) ...
    learn-packer.amazon-ebs.ubuntu: Processing triggers for ureadahead (0.100.0-19.1) ...
    learn-packer.amazon-ebs.ubuntu: Processing triggers for systemd (229-4ubuntu21.31) ...
==> learn-packer.amazon-ebs.ubuntu: Provisioning with shell script: /var/folders/4z/09jpfvfj6c19lxl7ch78pzvc0000gn/T/packer-shell2584371019
    learn-packer.amazon-ebs.ubuntu: This provisioner runs last
==> learn-packer.amazon-ebs.ubuntu: Stopping the source instance...
    learn-packer.amazon-ebs.ubuntu: Stopping instance
==> learn-packer.amazon-ebs.ubuntu: Waiting for the instance to stop...
==> learn-packer.amazon-ebs.ubuntu: Creating AMI learn-packer-linux-aws from instance i-04dc0273af6fc9c08
    learn-packer.amazon-ebs.ubuntu: AMI: ami-01c0179fe8d6cfad8
==> learn-packer.amazon-ebs.ubuntu: Waiting for AMI to become ready...
==> learn-packer.amazon-ebs.ubuntu: Skipping Enable AMI deprecation...
==> learn-packer.amazon-ebs.ubuntu: Terminating the source AWS instance...
==> learn-packer.amazon-ebs.ubuntu: Cleaning up any extra volumes...
==> learn-packer.amazon-ebs.ubuntu: No volumes to clean up, skipping
==> learn-packer.amazon-ebs.ubuntu: Deleting temporary security group...
==> learn-packer.amazon-ebs.ubuntu: Deleting temporary keypair...
Build 'learn-packer.amazon-ebs.ubuntu' finished after 4 minutes 28 seconds.

==> Wait completed after 4 minutes 28 seconds

==> Builds finished. The artifacts of successful builds are:
--> learn-packer.amazon-ebs.ubuntu: AMIs were created:
us-west-2: ami-01c0179fe8d6cfad8

packer_tutorial $ 
```

```bash
packer_tutorial $ ssh -i ~/.ssh/aws-us-west-2.pem ubuntu@52.41.43.83
ssh: connect to host 52.41.43.83 port 22: Connection refused
packer_tutorial $ ssh -v -i ~/.ssh/aws-us-west-2.pem ubuntu@52.41.43.83
OpenSSH_8.1p1, LibreSSL 2.7.3
debug1: Reading configuration data /Users/karuppiahn/.ssh/config
debug1: Reading configuration data /etc/ssh/ssh_config
debug1: /etc/ssh/ssh_config line 47: Applying options for *
debug1: Connecting to 52.41.43.83 [52.41.43.83] port 22.
debug1: Connection established.
debug1: identity file /Users/karuppiahn/.ssh/aws-us-west-2.pem type -1
debug1: identity file /Users/karuppiahn/.ssh/aws-us-west-2.pem-cert type -1
debug1: Local version string SSH-2.0-OpenSSH_8.1
debug1: Remote protocol version 2.0, remote software version OpenSSH_7.2p2 Ubuntu-4ubuntu2.10
debug1: match: OpenSSH_7.2p2 Ubuntu-4ubuntu2.10 pat OpenSSH_7.0*,OpenSSH_7.1*,OpenSSH_7.2*,OpenSSH_7.3*,OpenSSH_7.4*,OpenSSH_7.5*,OpenSSH_7.6*,OpenSSH_7.7* compat 0x04000002
debug1: Authenticating to 52.41.43.83:22 as 'ubuntu'
debug1: SSH2_MSG_KEXINIT sent
debug1: SSH2_MSG_KEXINIT received
debug1: kex: algorithm: curve25519-sha256@libssh.org
debug1: kex: host key algorithm: ecdsa-sha2-nistp256
debug1: kex: server->client cipher: chacha20-poly1305@openssh.com MAC: <implicit> compression: none
debug1: kex: client->server cipher: chacha20-poly1305@openssh.com MAC: <implicit> compression: none
debug1: expecting SSH2_MSG_KEX_ECDH_REPLY
debug1: Server host key: ecdsa-sha2-nistp256 SHA256:9B0kmj5dK5gJUc4CtQAIKubJxYd0Lt/I40UCBkXxaZQ
The authenticity of host '52.41.43.83 (52.41.43.83)' can't be established.
ECDSA key fingerprint is SHA256:9B0kmj5dK5gJUc4CtQAIKubJxYd0Lt/I40UCBkXxaZQ.
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added '52.41.43.83' (ECDSA) to the list of known hosts.
debug1: rekey out after 134217728 blocks
debug1: SSH2_MSG_NEWKEYS sent
debug1: expecting SSH2_MSG_NEWKEYS
debug1: SSH2_MSG_NEWKEYS received
debug1: rekey in after 134217728 blocks
debug1: Will attempt key: karuppiahn@vmware.com ED25519 SHA256:3vku70losGmr1kmlRT52RuAG4e0IEDevQV+uOmCL3NI agent
debug1: Will attempt key: /Users/karuppiahn/.ssh/aws-us-west-2.pem  explicit
debug1: SSH2_MSG_EXT_INFO received
debug1: kex_input_ext_info: server-sig-algs=<rsa-sha2-256,rsa-sha2-512>
debug1: SSH2_MSG_SERVICE_ACCEPT received
debug1: Authentications that can continue: publickey
debug1: Next authentication method: publickey
debug1: Offering public key: karuppiahn@vmware.com ED25519 SHA256:3vku70losGmr1kmlRT52RuAG4e0IEDevQV+uOmCL3NI agent
debug1: Authentications that can continue: publickey
debug1: Trying private key: /Users/karuppiahn/.ssh/aws-us-west-2.pem
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@         WARNING: UNPROTECTED PRIVATE KEY FILE!          @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
Permissions 0644 for '/Users/karuppiahn/.ssh/aws-us-west-2.pem' are too open.
It is required that your private key files are NOT accessible by others.
This private key will be ignored.
Load key "/Users/karuppiahn/.ssh/aws-us-west-2.pem": bad permissions
debug1: No more authentication methods to try.
ubuntu@52.41.43.83: Permission denied (publickey).
packer_tutorial $ chmod ~/.ssh/aws
aws                      aws-us-west-2.pem        aws-windows-machine.pem  aws.pub
packer_tutorial $ chmod ~/.ssh/aws
aws                      aws-us-west-2.pem        aws-windows-machine.pem  aws.pub
packer_tutorial $ ls ~/.ssh/aws-*
/Users/karuppiahn/.ssh/aws-us-west-2.pem	/Users/karuppiahn/.ssh/aws-windows-machine.pem
packer_tutorial $ ls -al ~/.ssh/aws-*
-rw-r--r--@ 1 karuppiahn  staff  1700 Sep 26 16:04 /Users/karuppiahn/.ssh/aws-us-west-2.pem
-rw-r--r--@ 1 karuppiahn  staff  1704 Aug 24 13:54 /Users/karuppiahn/.ssh/aws-windows-machine.pem
packer_tutorial $ chmod 600 ~/.ssh/aws-*
packer_tutorial $ ls -al ~/.ssh/aws-*
-rw-------@ 1 karuppiahn  staff  1700 Sep 26 16:04 /Users/karuppiahn/.ssh/aws-us-west-2.pem
-rw-------@ 1 karuppiahn  staff  1704 Aug 24 13:54 /Users/karuppiahn/.ssh/aws-windows-machine.pem
packer_tutorial $ ssh -v -i ~/.ssh/aws-us-west-2.pem ubuntu@52.41.43.83
OpenSSH_8.1p1, LibreSSL 2.7.3
debug1: Reading configuration data /Users/karuppiahn/.ssh/config
debug1: Reading configuration data /etc/ssh/ssh_config
debug1: /etc/ssh/ssh_config line 47: Applying options for *
debug1: Connecting to 52.41.43.83 [52.41.43.83] port 22.
debug1: Connection established.
debug1: identity file /Users/karuppiahn/.ssh/aws-us-west-2.pem type -1
debug1: identity file /Users/karuppiahn/.ssh/aws-us-west-2.pem-cert type -1
debug1: Local version string SSH-2.0-OpenSSH_8.1
debug1: Remote protocol version 2.0, remote software version OpenSSH_7.2p2 Ubuntu-4ubuntu2.10
debug1: match: OpenSSH_7.2p2 Ubuntu-4ubuntu2.10 pat OpenSSH_7.0*,OpenSSH_7.1*,OpenSSH_7.2*,OpenSSH_7.3*,OpenSSH_7.4*,OpenSSH_7.5*,OpenSSH_7.6*,OpenSSH_7.7* compat 0x04000002
debug1: Authenticating to 52.41.43.83:22 as 'ubuntu'
debug1: SSH2_MSG_KEXINIT sent
debug1: SSH2_MSG_KEXINIT received
debug1: kex: algorithm: curve25519-sha256@libssh.org
debug1: kex: host key algorithm: ecdsa-sha2-nistp256
debug1: kex: server->client cipher: chacha20-poly1305@openssh.com MAC: <implicit> compression: none
debug1: kex: client->server cipher: chacha20-poly1305@openssh.com MAC: <implicit> compression: none
debug1: expecting SSH2_MSG_KEX_ECDH_REPLY
debug1: Server host key: ecdsa-sha2-nistp256 SHA256:9B0kmj5dK5gJUc4CtQAIKubJxYd0Lt/I40UCBkXxaZQ
debug1: Host '52.41.43.83' is known and matches the ECDSA host key.
debug1: Found key in /Users/karuppiahn/.ssh/known_hosts:51
debug1: rekey out after 134217728 blocks
debug1: SSH2_MSG_NEWKEYS sent
debug1: expecting SSH2_MSG_NEWKEYS
debug1: SSH2_MSG_NEWKEYS received
debug1: rekey in after 134217728 blocks
debug1: Will attempt key: karuppiahn@vmware.com ED25519 SHA256:3vku70losGmr1kmlRT52RuAG4e0IEDevQV+uOmCL3NI agent
debug1: Will attempt key: /Users/karuppiahn/.ssh/aws-us-west-2.pem  explicit
debug1: SSH2_MSG_EXT_INFO received
debug1: kex_input_ext_info: server-sig-algs=<rsa-sha2-256,rsa-sha2-512>
debug1: SSH2_MSG_SERVICE_ACCEPT received
debug1: Authentications that can continue: publickey
debug1: Next authentication method: publickey
debug1: Offering public key: karuppiahn@vmware.com ED25519 SHA256:3vku70losGmr1kmlRT52RuAG4e0IEDevQV+uOmCL3NI agent
debug1: Authentications that can continue: publickey
debug1: Trying private key: /Users/karuppiahn/.ssh/aws-us-west-2.pem
debug1: Authentication succeeded (publickey).
Authenticated to 52.41.43.83 ([52.41.43.83]:22).
debug1: channel 0: new [client-session]
debug1: Requesting no-more-sessions@openssh.com
debug1: Entering interactive session.
debug1: pledge: network
debug1: client_input_global_request: rtype hostkeys-00@openssh.com want_reply 0
debug1: Sending environment.
debug1: Sending env LC_CTYPE = UTF-8
Welcome to Ubuntu 16.04.7 LTS (GNU/Linux 4.4.0-1128-aws x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

UA Infra: Extended Security Maintenance (ESM) is not enabled.

6 updates can be applied immediately.
3 of these updates are standard security updates.
To see these additional updates run: apt list --upgradable

40 additional security updates can be applied with UA Infra: ESM
Learn more about enabling UA Infra: ESM service for Ubuntu 16.04 at
https://ubuntu.com/16-04

New release '18.04.6 LTS' available.
Run 'do-release-upgrade' to upgrade to it.


_____________________________________________________________________
WARNING! Your environment specifies an invalid locale.
 The unknown environment variables are:
   LC_CTYPE=UTF-8 LC_ALL=
 This can affect your user experience significantly, including the
 ability to manage packages. You may install the locales by running:

   sudo apt-get install language-pack-UTF-8
     or
   sudo locale-gen UTF-8

To see all available language packs, run:
   apt-cache search "^language-pack-[a-z][a-z]$"
To disable this message for all users, run:
   sudo touch /var/lib/cloud/instance/locale-check.skip
_____________________________________________________________________

ubuntu@ip-172-31-12-173:~$ ls
example.txt
ubuntu@ip-172-31-12-173:~$ redis-server 
1483:C 26 Sep 10:37:05.109 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
1483:M 26 Sep 10:37:05.110 * Increased maximum number of open files to 10032 (it was originally set to 1024).
                _._                                                  
           _.-``__ ''-._                                             
      _.-``    `.  `_.  ''-._           Redis 3.0.6 (00000000/0) 64 bit
  .-`` .-```.  ```\/    _.,_ ''-._                                   
 (    '      ,       .-`  | `,    )     Running in standalone mode
 |`-._`-...-` __...-.``-._|'` _.-'|     Port: 6379
 |    `-._   `._    /     _.-'    |     PID: 1483
  `-._    `-._  `-./  _.-'    _.-'                                   
 |`-._`-._    `-.__.-'    _.-'_.-'|                                  
 |    `-._`-._        _.-'_.-'    |           http://redis.io        
  `-._    `-._`-.__.-'_.-'    _.-'                                   
 |`-._`-._    `-.__.-'    _.-'_.-'|                                  
 |    `-._`-._        _.-'_.-'    |                                  
  `-._    `-._`-.__.-'_.-'    _.-'                                   
      `-._    `-.__.-'    _.-'                                       
          `-._        _.-'                                           
              `-.__.-'                                               

1483:M 26 Sep 10:37:05.111 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
1483:M 26 Sep 10:37:05.111 # Server started, Redis version 3.0.6
1483:M 26 Sep 10:37:05.111 # WARNING overcommit_memory is set to 0! Background save may fail under low memory condition. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
1483:M 26 Sep 10:37:05.111 # WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled.
1483:M 26 Sep 10:37:05.111 * The server is now ready to accept connections on port 6379
^C1483:signal-handler (1632652626) Received SIGINT scheduling shutdown...
1483:M 26 Sep 10:37:06.816 # User requested shutdown...
1483:M 26 Sep 10:37:06.816 * Saving the final RDB snapshot before exiting.
1483:M 26 Sep 10:37:06.819 * DB saved on disk
1483:M 26 Sep 10:37:06.819 # Redis is now ready to exit, bye bye...
ubuntu@ip-172-31-12-173:~$ redis-server &
[1] 1486
ubuntu@ip-172-31-12-173:~$ 1486:C 26 Sep 10:37:16.837 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
1486:M 26 Sep 10:37:16.838 * Increased maximum number of open files to 10032 (it was originally set to 1024).
                _._                                                  
           _.-``__ ''-._                                             
      _.-``    `.  `_.  ''-._           Redis 3.0.6 (00000000/0) 64 bit
  .-`` .-```.  ```\/    _.,_ ''-._                                   
 (    '      ,       .-`  | `,    )     Running in standalone mode
 |`-._`-...-` __...-.``-._|'` _.-'|     Port: 6379
 |    `-._   `._    /     _.-'    |     PID: 1486
  `-._    `-._  `-./  _.-'    _.-'                                   
 |`-._`-._    `-.__.-'    _.-'_.-'|                                  
 |    `-._`-._        _.-'_.-'    |           http://redis.io        
  `-._    `-._`-.__.-'_.-'    _.-'                                   
 |`-._`-._    `-.__.-'    _.-'_.-'|                                  
 |    `-._`-._        _.-'_.-'    |                                  
  `-._    `-._`-.__.-'_.-'    _.-'                                   
      `-._    `-.__.-'    _.-'                                       
          `-._        _.-'                                           
              `-.__.-'                                               

1486:M 26 Sep 10:37:16.839 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
1486:M 26 Sep 10:37:16.839 # Server started, Redis version 3.0.6
1486:M 26 Sep 10:37:16.839 # WARNING overcommit_memory is set to 0! Background save may fail under low memory condition. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
1486:M 26 Sep 10:37:16.839 # WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled.
1486:M 26 Sep 10:37:16.839 * DB loaded from disk: 0.000 seconds
1486:M 26 Sep 10:37:16.839 * The server is now ready to accept connections on port 6379

ubuntu@ip-172-31-12-173:~$ redis-cli
127.0.0.1:6379> info server
# Server
redis_version:3.0.6
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:7785291a3d2152db
redis_mode:standalone
os:Linux 4.4.0-1128-aws x86_64
arch_bits:64
multiplexing_api:epoll
gcc_version:5.4.0
process_id:1240
run_id:8568e201044a9150026b4f14c0fc9373f5895412
tcp_port:6379
uptime_in_seconds:99
uptime_in_days:0
hz:10
lru_clock:5262691
config_file:/etc/redis/redis.conf
127.0.0.1:6379> 
ubuntu@ip-172-31-12-173:~$ cat example.txt 
FOO is hello world
ubuntu@ip-172-31-12-173:~$ echo $FUNCNAME ^C
ubuntu@ip-172-31-12-173:~$ ps aux
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  1.9  0.5  37552  5476 ?        Ss   10:35   0:02 /sbin/init
root         2  0.0  0.0      0     0 ?        S    10:35   0:00 [kthreadd]
root         3  0.0  0.0      0     0 ?        S    10:35   0:00 [ksoftirqd/0]
root         4  0.0  0.0      0     0 ?        S    10:35   0:00 [kworker/0:0]
root         5  0.0  0.0      0     0 ?        S<   10:35   0:00 [kworker/0:0H]
root         6  0.0  0.0      0     0 ?        S    10:35   0:00 [kworker/u30:0]
root         7  0.1  0.0      0     0 ?        S    10:35   0:00 [rcu_sched]
root         8  0.0  0.0      0     0 ?        S    10:35   0:00 [rcu_bh]
root         9  0.0  0.0      0     0 ?        S    10:35   0:00 [migration/0]
root        10  0.0  0.0      0     0 ?        S    10:35   0:00 [watchdog/0]
root        11  0.0  0.0      0     0 ?        S    10:35   0:00 [kdevtmpfs]
root        12  0.0  0.0      0     0 ?        S<   10:35   0:00 [netns]
root        13  0.0  0.0      0     0 ?        S<   10:35   0:00 [perf]
root        14  0.0  0.0      0     0 ?        S    10:35   0:00 [xenwatch]
root        15  0.0  0.0      0     0 ?        S    10:35   0:00 [xenbus]
root        16  0.0  0.0      0     0 ?        S    10:35   0:00 [kworker/0:1]
root        17  0.0  0.0      0     0 ?        S    10:35   0:00 [khungtaskd]
root        18  0.0  0.0      0     0 ?        S<   10:35   0:00 [writeback]
root        19  0.0  0.0      0     0 ?        SN   10:35   0:00 [ksmd]
root        20  0.0  0.0      0     0 ?        SN   10:35   0:00 [khugepaged]
root        21  0.0  0.0      0     0 ?        S<   10:35   0:00 [crypto]
root        22  0.0  0.0      0     0 ?        S<   10:35   0:00 [kintegrityd]
root        23  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        24  0.0  0.0      0     0 ?        S<   10:35   0:00 [kblockd]
root        25  0.0  0.0      0     0 ?        S<   10:35   0:00 [ata_sff]
root        26  0.0  0.0      0     0 ?        S<   10:35   0:00 [md]
root        27  0.0  0.0      0     0 ?        S<   10:35   0:00 [devfreq_wq]
root        28  0.0  0.0      0     0 ?        S    10:35   0:00 [kworker/u30:1]
root        30  0.0  0.0      0     0 ?        S    10:35   0:00 [kswapd0]
root        31  0.0  0.0      0     0 ?        S<   10:35   0:00 [vmstat]
root        32  0.0  0.0      0     0 ?        S    10:35   0:00 [fsnotify_mark]
root        33  0.0  0.0      0     0 ?        S    10:35   0:00 [ecryptfs-kthrea]
root        49  0.0  0.0      0     0 ?        S<   10:35   0:00 [kthrotld]
root        50  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        51  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        52  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        53  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        54  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        55  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        56  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        57  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        58  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        59  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        60  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        61  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        62  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        63  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        64  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        65  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        66  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        67  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        68  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        69  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        70  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        71  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        72  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        73  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        74  0.0  0.0      0     0 ?        S<   10:35   0:00 [nvme]
root        75  0.0  0.0      0     0 ?        S    10:35   0:00 [scsi_eh_0]
root        76  0.0  0.0      0     0 ?        S<   10:35   0:00 [scsi_tmf_0]
root        77  0.0  0.0      0     0 ?        S    10:35   0:00 [scsi_eh_1]
root        78  0.0  0.0      0     0 ?        S<   10:35   0:00 [scsi_tmf_1]
root        79  0.0  0.0      0     0 ?        S    10:35   0:00 [kworker/u30:2]
root        83  0.0  0.0      0     0 ?        S<   10:35   0:00 [ipv6_addrconf]
root        85  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root        97  0.0  0.0      0     0 ?        S<   10:35   0:00 [deferwq]
root       248  0.0  0.0      0     0 ?        S<   10:35   0:00 [raid5wq]
root       275  0.0  0.0      0     0 ?        S<   10:35   0:00 [bioset]
root       305  0.0  0.0      0     0 ?        S    10:35   0:00 [jbd2/xvda1-8]
root       306  0.0  0.0      0     0 ?        S<   10:35   0:00 [ext4-rsv-conver]
root       366  0.0  0.0      0     0 ?        S    10:35   0:00 [kworker/0:2]
root       369  0.0  0.0      0     0 ?        S    10:35   0:00 [kworker/0:3]
root       372  0.0  0.0      0     0 ?        S    10:35   0:00 [kworker/0:4]
root       373  0.0  0.0      0     0 ?        S<   10:35   0:00 [kworker/0:1H]
root       376  0.0  0.0      0     0 ?        S    10:35   0:00 [kworker/0:5]
root       379  0.0  0.0      0     0 ?        S    10:35   0:00 [kworker/0:6]
root       382  0.0  0.0      0     0 ?        S    10:35   0:00 [kworker/0:7]
root       385  0.0  0.0      0     0 ?        S    10:35   0:00 [kworker/0:8]
root       394  0.4  0.2  28348  2700 ?        Ss   10:35   0:00 /lib/systemd/systemd-journald
root       401  0.0  0.0      0     0 ?        S<   10:35   0:00 [iscsi_eh]
root       410  0.0  0.0      0     0 ?        S    10:35   0:00 [kauditd]
root       436  0.0  0.0      0     0 ?        S<   10:35   0:00 [ib_addr]
root       442  0.0  0.1  94768  1564 ?        Ss   10:35   0:00 /sbin/lvmetad -f
root       445  0.0  0.0      0     0 ?        S<   10:35   0:00 [ib_mcast]
root       446  0.0  0.0      0     0 ?        S<   10:35   0:00 [ib_nl_sa_wq]
root       449  0.0  0.0      0     0 ?        S<   10:35   0:00 [ib_cm]
root       451  0.0  0.0      0     0 ?        S<   10:35   0:00 [iw_cm_wq]
root       452  0.0  0.0      0     0 ?        S<   10:35   0:00 [rdma_cm]
root       461  0.0  0.4  43600  4952 ?        Ss   10:35   0:00 /lib/systemd/systemd-udevd
root       497  0.0  0.0      0     0 ?        S<   10:35   0:00 [loop0]
root       515  0.0  0.0      0     0 ?        S<   10:35   0:00 [loop1]
root       522  0.0  0.0      0     0 ?        S<   10:35   0:00 [loop2]
systemd+   547  0.0  0.2 100320  2476 ?        Ssl  10:35   0:00 /lib/systemd/systemd-timesyncd
root       911  0.0  0.0  16120   856 ?        Ss   10:35   0:00 /sbin/dhclient -1 -v -pf /run/dhclient.eth0.pid -lf 
root      1095  0.0  0.3  27724  3044 ?        Ss   10:35   0:00 /usr/sbin/cron -f
root      1106  0.0  0.1   4392  1216 ?        Ss   10:35   0:00 /usr/sbin/acpid
message+  1109  0.0  0.3  42896  3764 ?        Ss   10:35   0:00 /usr/bin/dbus-daemon --system --address=systemd: --n
root      1115  0.0  0.0   5216   152 ?        Ss   10:35   0:00 /sbin/iscsid
root      1116  0.0  0.3   5716  3512 ?        S<Ls 10:35   0:00 /sbin/iscsid
syslog    1121  0.0  0.3 260624  3644 ?        Ssl  10:35   0:00 /usr/sbin/rsyslogd -n
root      1127  0.3  2.0 320240 21248 ?        Ssl  10:35   0:00 /snap/amazon-ssm-agent/3552/amazon-ssm-agent
root      1129  0.0  0.6 274484  6232 ?        Ssl  10:35   0:00 /usr/lib/accountsservice/accounts-daemon
daemon    1137  0.0  0.2  26040  2188 ?        Ss   10:35   0:00 /usr/sbin/atd -f
root      1141  0.0  0.3  28544  3044 ?        Ss   10:35   0:00 /lib/systemd/systemd-logind
root      1152  0.0  0.1  95364  1580 ?        Ssl  10:35   0:00 /usr/bin/lxcfs /var/lib/lxcfs/
root      1155  0.6  2.8 302760 28964 ?        Ssl  10:35   0:00 /usr/lib/snapd/snapd
root      1166  0.0  0.5  65508  5572 ?        Ss   10:35   0:00 /usr/sbin/sshd -D
root      1176  0.0  1.9 175760 19992 ?        Ssl  10:35   0:00 /usr/bin/python3 /usr/share/unattended-upgrades/unat
root      1192  0.0  0.0  13368   156 ?        Ss   10:35   0:00 /sbin/mdadm --monitor --pid-file /run/mdadm/monitor.
redis     1240  0.0  0.3  38852  3080 ?        Ssl  10:35   0:00 /usr/bin/redis-server 127.0.0.1:6379
root      1244  0.0  0.2  14468  2296 ttyS0    Ss+  10:35   0:00 /sbin/agetty --keep-baud 115200 38400 9600 ttyS0 vt2
root      1250  0.0  0.1  14652  1824 tty1     Ss+  10:35   0:00 /sbin/agetty --noclear tty1 linux
root      1272  0.0  0.5 277176  5896 ?        Ssl  10:35   0:00 /usr/lib/policykit-1/polkitd --no-debug
root      1297  0.7  3.8 402396 38864 ?        Sl   10:35   0:00 /snap/amazon-ssm-agent/3552/ssm-agent-worker
root      1402  0.0  0.6  92796  6760 ?        Ss   10:36   0:00 sshd: ubuntu [priv]
ubuntu    1404  0.0  0.4  45008  4440 ?        Ss   10:36   0:00 /lib/systemd/systemd --user
ubuntu    1406  0.0  0.1  61004  1756 ?        S    10:36   0:00 (sd-pam)
ubuntu    1466  0.0  0.3  92796  3364 ?        S    10:36   0:00 sshd: ubuntu@pts/0
ubuntu    1467  0.1  0.5  21384  5328 pts/0    Ss   10:36   0:00 -bash
ubuntu    1486  0.0  0.3  38852  3688 pts/0    Sl   10:37   0:00 redis-server *:6379
ubuntu    1497  0.0  0.3  36052  3260 pts/0    R+   10:37   0:00 ps aux
ubuntu@ip-172-31-12-173:~$ ps aux | grep redis
redis     1240  0.0  0.3  38852  3080 ?        Ssl  10:35   0:00 /usr/bin/redis-server 127.0.0.1:6379
ubuntu    1486  0.0  0.3  38852  3688 pts/0    Sl   10:37   0:00 redis-server *:6379
ubuntu    1499  0.0  0.0  12912   928 pts/0    S+   10:37   0:00 grep --color=auto redis
ubuntu@ip-172-31-12-173:~$ kill 1486
ubuntu@ip-172-31-12-173:~$ 1486:signal-handler (1632652672) Received SIGTERM scheduling shutdown...
1486:M 26 Sep 10:37:52.131 # User requested shutdown...
1486:M 26 Sep 10:37:52.131 * Saving the final RDB snapshot before exiting.
1486:M 26 Sep 10:37:52.133 * DB saved on disk
1486:M 26 Sep 10:37:52.133 # Redis is now ready to exit, bye bye...

[1]+  Done                    redis-server
ubuntu@ip-172-31-12-173:~$ ps aux | grep redis
redis     1240  0.0  0.3  38852  3080 ?        Ssl  10:35   0:00 /usr/bin/redis-server 127.0.0.1:6379
ubuntu    1502  0.0  0.0  12912   960 pts/0    S+   10:37   0:00 grep --color=auto redis
ubuntu@ip-172-31-12-173:~$ kill 1240
-bash: kill: (1240) - Operation not permitted
ubuntu@ip-172-31-12-173:~$ sudo kill 1240
ubuntu@ip-172-31-12-173:~$ ps aux | grep redis
redis     1530  0.0  0.3  38852  3188 ?        Ssl  10:38   0:00 /usr/bin/redis-server 127.0.0.1:6379
ubuntu    1539  0.0  0.0  12912  1012 pts/0    S+   10:38   0:00 grep --color=auto redis
ubuntu@ip-172-31-12-173:~$ system
systemctl                       systemd-delta                   systemd-path
systemd                         systemd-detect-virt             systemd-resolve
systemd-analyze                 systemd-escape                  systemd-run
systemd-ask-password            systemd-hwdb                    systemd-stdio-bridge
systemd-cat                     systemd-inhibit                 systemd-tmpfiles
systemd-cgls                    systemd-machine-id-setup        systemd-tty-ask-password-agent
systemd-cgtop                   systemd-notify                  
ubuntu@ip-172-31-12-173:~$ systemctl status redis
● redis-server.service - Advanced key-value store
   Loaded: loaded (/lib/systemd/system/redis-server.service; enabled; vendor preset: enabled)
   Active: active (running) since Sun 2021-09-26 10:38:16 UTC; 17s ago
     Docs: http://redis.io/documentation,
           man:redis-server(1)
  Process: 1514 ExecStopPost=/bin/run-parts --verbose /etc/redis/redis-server.post-down.d (code=exited, status=0/SUCC
  Process: 1510 ExecStop=/bin/kill -s TERM $MAINPID (code=exited, status=0/SUCCESS)
  Process: 1505 ExecStop=/bin/run-parts --verbose /etc/redis/redis-server.pre-down.d (code=exited, status=0/SUCCESS)
  Process: 1531 ExecStartPost=/bin/run-parts --verbose /etc/redis/redis-server.post-up.d (code=exited, status=0/SUCCE
  Process: 1527 ExecStart=/usr/bin/redis-server /etc/redis/redis.conf (code=exited, status=0/SUCCESS)
  Process: 1523 ExecStartPre=/bin/run-parts --verbose /etc/redis/redis-server.pre-up.d (code=exited, status=0/SUCCESS
 Main PID: 1530 (redis-server)
    Tasks: 3
   Memory: 852.0K
      CPU: 39ms
   CGroup: /system.slice/redis-server.service
           └─1530 /usr/bin/redis-server 127.0.0.1:6379       

Sep 26 10:38:16 ip-172-31-12-173 systemd[1]: redis-server.service: Service hold-off time over, scheduling restart.
Sep 26 10:38:16 ip-172-31-12-173 systemd[1]: Stopped Advanced key-value store.
Sep 26 10:38:16 ip-172-31-12-173 systemd[1]: Starting Advanced key-value store...
Sep 26 10:38:16 ip-172-31-12-173 run-parts[1523]: run-parts: executing /etc/redis/redis-server.pre-up.d/00_example
Sep 26 10:38:16 ip-172-31-12-173 run-parts[1531]: run-parts: executing /etc/redis/redis-server.post-up.d/00_example
Sep 26 10:38:16 ip-172-31-12-173 systemd[1]: Started Advanced key-value store.
ubuntu@ip-172-31-12-173:~$ systemctl stop redis
==== AUTHENTICATING FOR org.freedesktop.systemd1.manage-units ===
Authentication is required to stop 'redis-server.service'.
Authenticating as: Ubuntu (ubuntu)
Password: 
ubuntu@ip-172-31-12-173:~$ sudo systemctl stop redis
ubuntu@ip-172-31-12-173:~$ systemctl status redis
● redis-server.service - Advanced key-value store
   Loaded: loaded (/lib/systemd/system/redis-server.service; enabled; vendor preset: enabled)
   Active: inactive (dead) since Sun 2021-09-26 10:38:49 UTC; 3s ago
     Docs: http://redis.io/documentation,
           man:redis-server(1)
  Process: 1562 ExecStopPost=/bin/run-parts --verbose /etc/redis/redis-server.post-down.d (code=exited, status=0/SUCC
  Process: 1556 ExecStop=/bin/kill -s TERM $MAINPID (code=exited, status=0/SUCCESS)
  Process: 1552 ExecStop=/bin/run-parts --verbose /etc/redis/redis-server.pre-down.d (code=exited, status=0/SUCCESS)
  Process: 1531 ExecStartPost=/bin/run-parts --verbose /etc/redis/redis-server.post-up.d (code=exited, status=0/SUCCE
  Process: 1527 ExecStart=/usr/bin/redis-server /etc/redis/redis.conf (code=exited, status=0/SUCCESS)
  Process: 1523 ExecStartPre=/bin/run-parts --verbose /etc/redis/redis-server.pre-up.d (code=exited, status=0/SUCCESS
 Main PID: 1530 (code=exited, status=0/SUCCESS)

Sep 26 10:38:16 ip-172-31-12-173 systemd[1]: redis-server.service: Service hold-off time over, scheduling restart.
Sep 26 10:38:16 ip-172-31-12-173 systemd[1]: Stopped Advanced key-value store.
Sep 26 10:38:16 ip-172-31-12-173 systemd[1]: Starting Advanced key-value store...
Sep 26 10:38:16 ip-172-31-12-173 run-parts[1523]: run-parts: executing /etc/redis/redis-server.pre-up.d/00_example
Sep 26 10:38:16 ip-172-31-12-173 run-parts[1531]: run-parts: executing /etc/redis/redis-server.post-up.d/00_example
Sep 26 10:38:16 ip-172-31-12-173 systemd[1]: Started Advanced key-value store.
Sep 26 10:38:49 ip-172-31-12-173 systemd[1]: Stopping Advanced key-value store...
Sep 26 10:38:49 ip-172-31-12-173 run-parts[1552]: run-parts: executing /etc/redis/redis-server.pre-down.d/00_example
Sep 26 10:38:49 ip-172-31-12-173 run-parts[1562]: run-parts: executing /etc/redis/redis-server.post-down.d/00_example
Sep 26 10:38:49 ip-172-31-12-173 systemd[1]: Stopped Advanced key-value store.
ubuntu@ip-172-31-12-173:~$ debug1: client_input_channel_req: channel 0 rtype exit-status reply 0
debug1: client_input_channel_req: channel 0 rtype eow@openssh.com reply 0
logout
debug1: channel 0: free: client-session, nchannels 1
Connection to 52.41.43.83 closed.
Transferred: sent 7256, received 44604 bytes, in 122.6 seconds
Bytes per second: sent 59.2, received 363.9
debug1: Exit status 3
packer_tutorial $ 
```
