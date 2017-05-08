package dnspod

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CuriosityChina/dnspod-go/service"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDnspodRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnspodRecordCreate,
		Read:   resourceDnspodRecordRead,
		Update: resourceDnspodRecordUpdate,
		Delete: resourceDnspodRecordDelete,
		Schema: map[string]*schema.Schema{
			"domain_id": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "域名ID",
			},
			"sub_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "A",
				ValidateFunc: func(v interface{}, k string) (ws []string, es []error) {
					value := v.(string)
					opts := map[string]bool{
						"A":     true,
						"CNAME": true,
						"MX":    true,
						"TXT":   true,
						"NS":    true,
						"AAAA":  true,
						"SRV":   true,
						"显性URL": true,
						"隐性URL": true,
					}
					if !opts[value] {
						es = append(es, fmt.Errorf(
							"类型不正确 %q", k))
					}
					return
				},
			},
			"line": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "默认",
				ValidateFunc: func(v interface{}, k string) (ws []string, es []error) {
					value := v.(string)
					opts := []string{"默认", "国内", "国外", "电信", "联通", "教育网", "移动", "百度", "谷歌", "搜搜", "有道", "必应", "搜狗", "奇虎", "搜索引擎"}
					var ok bool
					for i := range opts {
						if opts[i] == value {
							return
						}
					}
					if !ok {
						es = append(es, fmt.Errorf(
							"类型不正确 %q", k))
					}
					return
				},
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mx": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, es []error) {
					value, _ := strconv.Atoi(v.(string))
					if 1 <= value && value <= 20 {
						return
					}
					es = append(es, fmt.Errorf(
						"范围1-20 %q", k))
					return
				},
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "600",
				ValidateFunc: func(v interface{}, k string) (ws []string, es []error) {
					value, _ := strconv.Atoi(v.(string))
					if 1 <= value && value <= 604800 {
						return
					}
					es = append(es, fmt.Errorf(
						"范围1-604800，不同等级域名最小值不同 %q", k))
					return
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "enable",
				ValidateFunc: func(v interface{}, k string) (ws []string, es []error) {
					value := v.(string)
					if value != "enable" && value != "disable" {
						es = append(es, fmt.Errorf(
							"范围1-604800，不同等级域名最小值不同 %q", k))
					}
					return
				},
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, es []error) {
					value, _ := strconv.Atoi(v.(string))
					if 1 <= value && value <= 100 {
						return
					}
					es = append(es, fmt.Errorf(
						"0到100的整数，可选。仅企业 VIP 域名可用 %q", k))
					return
				},
			},
		},
	}
}

func resourceDnspodRecordCreate(d *schema.ResourceData, meta interface{}) error {
	clt := meta.(*DNSPodClient).record
	params := service.RecordCreateRequest{
		DomainID:   d.Get("domain_id").(int),
		SubDomain:  d.Get("sub_domain").(string),
		RecordType: d.Get("type").(string),
		RecordLine: d.Get("line").(string),
		Value:      d.Get("value").(string),
		MX:         d.Get("mx").(string),
		TTL:        d.Get("ttl").(string),
		Status:     d.Get("status").(string),
		Weight:     d.Get("weight").(string),
	}
	resp, err := clt.RecordCreate(params)
	if err != nil {
		return err
	}
	d.SetId(resp.Record.ID)
	return nil
}

func resourceDnspodRecordRead(d *schema.ResourceData, meta interface{}) error {
	clt := meta.(*DNSPodClient).record
	params := service.RecordInfoRequest{
		RecordID: d.Id(),
		DomainID: d.Get("domain_id").(int),
	}
	resp, err := clt.RecordInfo(params)
	if err != nil {
		return err
	}
	d.Set("domain_id", resp.Domain.ID)
	d.Set("sub_domain", resp.Record.SubDomain)
	d.Set("value", resp.Record.Value)
	d.Set("ttl", resp.Record.TTL)
	d.Set("weight", resp.Record.Weight)
	return nil
}

func resourceDnspodRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	clt := meta.(*DNSPodClient).record
	params := service.RecordModifyRequest{
		RecordID:   d.Id(),
		DomainID:   d.Get("domain_id").(int),
		SubDomain:  d.Get("sub_domain").(string),
		RecordType: d.Get("type").(string),
		RecordLine: d.Get("line").(string),
		Value:      d.Get("value").(string),
		MX:         d.Get("mx").(string),
		TTL:        d.Get("ttl").(string),
		Status:     d.Get("status").(string),
		Weight:     d.Get("weight").(string),
	}
	resp, err := clt.RecordModify(params)
	if err != nil {
		return err
	}
	log.Printf("After change :%s", resp.Record.ID)
	d.SetId(strconv.Itoa(resp.Record.ID))
	return nil
}

func resourceDnspodRecordDelete(d *schema.ResourceData, meta interface{}) error {
	clt := meta.(*DNSPodClient).record
	params := service.RecordRemoveRequest{
		DomainID: d.Get("domain_id").(int),
		RecordID: d.Id(),
	}
	_, err := clt.RecordRemove(params)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
