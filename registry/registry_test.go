package registry

import (
	"testing"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/hashicorp/consul/api"
)

func TestConsul(t *testing.T) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	reg := consul.New(client)

	kratos.Registrar(reg)
}
