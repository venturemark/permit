package gateway

import (
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/permission"
)

type Config struct {
	Ingress  permission.Ingress
	Resolver permission.Resolver
}

type Gateway struct {
	ingress  permission.Ingress
	resolver permission.Resolver
}

func New(config Config) (*Gateway, error) {
	if config.Ingress == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Ingress must not be empty", config)
	}
	if config.Resolver == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Resolver must not be empty", config)
	}

	g := &Gateway{
		ingress:  config.Ingress,
		resolver: config.Resolver,
	}

	return g, nil
}

func (g *Gateway) Ingress() permission.Ingress {
	return g.ingress
}

func (g *Gateway) Resolver() permission.Resolver {
	return g.resolver
}
