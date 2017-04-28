package dnspod

import (
	"github.com/CuriosityChina/dnspod-go/client"
	"github.com/CuriosityChina/dnspod-go/service"
)

// Config dnspod provider configration
type Config struct {
	ID    string
	Token string
}

// DNSPodClient dnspod resource client
type DNSPodClient struct {
	record *service.RecordService
}

// Client dnspod provider client
func (c *Config) Client() (*DNSPodClient, error) {
	clt := client.NewClient(c.ID, c.Token)
	return &DNSPodClient{
		record: service.NewRecordService(clt),
	}, nil
}
