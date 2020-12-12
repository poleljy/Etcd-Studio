package storev3

import (
	"context"
	"fmt"

	"go.etcd.io/etcd/clientv3"
)

func (e *DynamicEtcd) GetKeysOnly(prefix string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var keys []string
	resp, err := e.cli.Get(ctx, prefix, clientv3.WithKeysOnly(), clientv3.WithPrefix())
	if err != nil {
		return keys, MakeError(err, fmt.Sprintf("failed to get keys with prefix '%s'", prefix))
	}

	keys = make([]string, 0, len(resp.Kvs))
	for _, ev := range resp.Kvs {
		keys = append(keys, string(ev.Key))
	}
	return keys, nil
}

func (e *DynamicEtcd) GetKeyValue(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	resp, err := e.cli.Get(ctx, key)
	if err != nil {
		return "", MakeError(err, fmt.Sprintf("failed to get value: '%s'", key))
	}
	if len(resp.Kvs) == 0 {
		return "", nil
	}
	return string(resp.Kvs[0].Value), nil
}
