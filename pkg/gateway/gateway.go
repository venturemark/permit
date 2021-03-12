package gateway

import (
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/permission"
)

type Config struct {
	Ingress  permission.Ingress
	Resource permission.Resource
}

type Gateway struct {
	ingress  permission.Ingress
	resource permission.Resource
}

func New(config Config) (*Gateway, error) {
	if config.Ingress == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Ingress must not be empty", config)
	}
	if config.Resource == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Resource must not be empty", config)
	}

	g := &Gateway{
		ingress:  config.Ingress,
		resource: config.Resource,
	}

	return g, nil
}

func (g *Gateway) Ingress() permission.Ingress {
	return g.ingress
}

func (g *Gateway) Resource() permission.Resource {
	return g.resource
}
