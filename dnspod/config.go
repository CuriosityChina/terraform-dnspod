package dnspod

import (
	"github.com/CuriosityChina/dnspod-go"
	"github.com/CuriosityChina/dnspod-go/record"
)

type Config struct {
	ID    string
	Token string
}

type DnsPodClient struct {
	record *record.RECORD
}

func (c *Config) Client() (*DnsPodClient, error) {
	clt := dnspod.NewClient(c.ID, c.Token)
	return &DnsPodClient{
		record: record.NewClient(clt),
	}, nil
}
