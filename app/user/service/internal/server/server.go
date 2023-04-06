package server

import (
	"log"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"go.etcd.io/etcd/client/v3"

	"github.com/869413421/micro-chat/app/user/service/internal/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar)

// NewRegistrar 使用ETCD注册中心
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{conf.Etcd.Address},
	})
	if err != nil {
		log.Fatal(err)
	}
	c := etcd.New(etcdClient)
	return c
}
