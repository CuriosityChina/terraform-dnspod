package dnspod

import (
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["id"],
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["token"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"dnspod_record": resourceDnspodRecord(),
		},
		ConfigureFunc: providerConfigure,
	}
}

var qingcloudMutexKV = mutexkv.NewMutexKV()

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ID:    d.Get("id").(string),
		Token: d.Get("token").(string),
	}
	return config.Client()
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"id":    "个人的ID",
		"token": "账号的 token ",
	}
}
