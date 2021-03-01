package gateway

import (
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/permission"
)

type Config struct {
	Ingress permission.Ingress
}

type Gateway struct {
	ingress permission.Ingress
}

func New(config Config) (*Gateway, error) {
	if config.Ingress == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Ingress must not be empty", config)
	}

	g := &Gateway{
		ingress: config.Ingress,
	}

	return g, nil
}

func (g *Gateway) Ingress() permission.Ingress {
	return g.ingress
}
