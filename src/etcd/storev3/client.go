package storev3

import (
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

const (
	dialTimeOut    = 15 * time.Second
	requestTimeout = 15 * time.Second
)

type DynamicEtcd struct {
	Endpoints []string
	cli       *clientv3.Client
}

func (e *DynamicEtcd) Connect() error {
	if e.cli != nil {
		return nil
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   e.Endpoints,
		DialTimeout: dialTimeOut,
	})
	if err != nil {
		log.Print("failed to create endpoint client(V3): ", err)
		return err
	}

	e.cli = cli
	return nil
}

func (e *DynamicEtcd) IsConnect() bool {
	return e.cli == nil
}

func (e *DynamicEtcd) DisConnect() error {
	return e.cli.Close()
}
