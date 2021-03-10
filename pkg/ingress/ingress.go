package ingress

import (
	"github.com/venturemark/permission/pkg/label"
)

type Config struct {
	Policy [][]label.Label
}

type Incoming struct {
	policy [][]label.Label
}

func New(config Config) (*Incoming, error) {
	if len(config.Policy) == 0 {
		config.Policy = policy
	}

	i := &Incoming{
		policy: config.Policy,
	}

	return i, nil
}

func (i *Incoming) Match(l ...label.Label) bool {
	for _, p := range i.policy {
		if equal(p, l) {
			return true
		}
	}

	return false
}
