package storev3

import (
	"context"

	"go.etcd.io/etcd/clientv3"
)

func (e *DynamicEtcd) PutKeyValue(key, value string, ttl int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var err error
	if ttl > 0 {
		leaseResp, err := e.cli.Grant(context.TODO(), ttl)
		if err == nil && leaseResp != nil {
			_, err = e.cli.Put(ctx, key, value, clientv3.WithLease(leaseResp.ID))
		}
	} else {
		_, err = e.cli.Put(ctx, key, value)
	}
	return err
}
