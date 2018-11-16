package etcdv3

import (
	"google.golang.org/grpc/naming"
	"errors"
	etcd3 "github.com/etcd-io/etcd/clientv3"
	"strings"
	"fmt"
)

type resolver struct {
	serviceNames string
}

func NewResolver(serviceName string) *resolver  {
	return &resolver{serviceNames:serviceName}
}

func (re *resolver)Resolve(target string) (naming.Watcher, error) {
	if re.serviceNames == "" {
		return nil, errors.New("grpclb: no service name provided")
	}

	client, err := etcd3.New(etcd3.Config{
		Endpoints:strings.Split(target, ","),
	})
	if err != nil {
		return nil, fmt.Errorf("grpclb: create etcd3 client failed:%s", err.Error())
	}
	return &watcher{}
}