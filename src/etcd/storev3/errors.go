package storev3

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

func MakeError(err error, prefix string) error {
	if err == nil {
		return nil
	}

	var msg string
	switch err {
	case context.Canceled:
		msg = fmt.Sprintf("ctx is canceled by another routine: %v", err)
	case context.DeadlineExceeded:
		msg = fmt.Sprintf("ctx is attached with a deadline is exceeded: %v", err)
	case rpctypes.ErrEmptyKey:
		msg = fmt.Sprintf("client-side error: %v", err)
	case rpctypes.ErrKeyNotFound:
		msg = fmt.Sprintf("client-side error: %v", err)
	default:
		msg = fmt.Sprintf("bad cluster endpoints, which are not etcd servers: %v", err)
	}

	if len(prefix) > 0 {
		msg = fmt.Sprintf("%s, %s", prefix, msg)
	}
	log.Print(msg)
	return errors.New(msg)
}
