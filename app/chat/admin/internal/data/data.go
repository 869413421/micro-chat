package data

import (
	"context"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"

	v1 "github.com/869413421/micro-chat/api/user/service/v1"
	"github.com/869413421/micro-chat/app/chat/admin/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDiscovery, NewRegistrar, NewUserServiceClient, NewUserRepo, NewRoleRepo, NewPermissionRepo)

// Data .
type Data struct {
	log *log.Helper
	uc  v1.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc v1.UserClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, uc: uc}, cleanup, nil
}

// NewDiscovery 服务发现
func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

// NewRegistrar 服务注册
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

// NewUserServiceClient .
func NewUserServiceClient(r registry.Discovery) v1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///micro-chat.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := v1.NewUserClient(conn)
	return c
}
