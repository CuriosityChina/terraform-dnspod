package dnspod

import (
	"github.com/CuriosityChina/dnspod-go/client"
	"github.com/CuriosityChina/dnspod-go/service"
)

type Config struct {
	ID    string
	Token string
}

type DnsPodClient struct {
	record *service.RecordService
}

func (c *Config) Client() (*DnsPodClient, error) {
	clt := client.NewClient(c.ID, c.Token)
	return &DnsPodClient{
		record: service.NewRecordService(clt),
	}, nil
}
