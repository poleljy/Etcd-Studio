package storev3

import (
	"context"
	"fmt"

	"go.etcd.io/etcd/clientv3"
)

func (e *DynamicEtcd) DeleteKey(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	dresp, err := e.cli.Delete(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	fmt.Println("Deleted all keys:", dresp.Deleted)
	return nil
}
