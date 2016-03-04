# Terraform DNSPod Provider [WIP]

## Sample
如下是一个使用案例
```
provider "dnspod"{
	id = "xxx"
	token = "xxxx"
}

variable "domain_id" {
	default = 123123123
}

# 设置以一个域名
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
# 设置第二个域名
resource "dnspod_record" "second"{
	domain_id = "${var.domain_id}"
	sub_domain = "b"
	value = "12.12.112.11"
}
```