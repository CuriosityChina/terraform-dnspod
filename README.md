# Terraform DNSPod Provider [WIP] [![Build Status](https://travis-ci.org/CuriosityChina/terraform-dnspod.svg?branch=master)](https://travis-ci.org/CuriosityChina/terraform-dnspod)

# Usage

## Installation dnspod-provider
```
go install -v go install -v github.com/CuriosityChina/terraform-dnspod/provider-dnspod
```

## Configuring the terraform plug in path

```
# edit
subl ~/.terraformrc

# Modify the following path to your local installation
providers {
	dnspod = "/Users/YOUR/GO/PATH/bin/provider-dnspod"
}
```

## Sample
Example
```
provider "dnspod"{
	id = "xxx"
	token = "xxxx"
}

variable "domain_id" {
	default = 123123123
}

# Domain name
resource "dnspod_record" "first"{
	domain_id = "${var.domain_id}"
	sub_domain = "a"
	type = "A"
	line = "默认"
	value = "12.12.112.11"
	mx = "2"
	ttl = "12312"
	status = "enable"
}
# Second domain name
resource "dnspod_record" "second"{
	domain_id = "${var.domain_id}"
	sub_domain = "b"
	value = "12.12.112.11"
}
```
