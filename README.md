# Terraform DNSPod Provider [WIP] [![Build Status](https://travis-ci.org/CuriosityChina/terraform-dnspod.svg?branch=master)](https://travis-ci.org/CuriosityChina/terraform-dnspod)

# 使用方式

## 安装 dnspod-provider
```
go install -v github.com/CuriosityChina/terraform-qingcloud/provider-dnspod
```

## 设置 terraform 的插件路径

```
# 启动编辑器
subl ~/.terraformrc

# 修改如下qingcloud 到你本地的路径
providers {
	dnspod = "/Users/YOUR/GO/PATH/bin/provider-dnspod"
}
```

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