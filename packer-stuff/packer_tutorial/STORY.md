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

