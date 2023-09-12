package depend

import (
	"github.com/Rascal0814/boot/config"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

// NewConsulRegistrar consul 服务注册
func NewConsulRegistrar(conf *config.Data) registry.Registrar {
	if conf.Consul == nil {
		return nil
	}
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
