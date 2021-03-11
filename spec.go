package permission

import "github.com/venturemark/permission/pkg/label"

type Gateway interface {
	Ingress() Ingress
	Resource() Resource
}

type Ingress interface {
	Match(l ...label.Label) bool
}

type Resource interface {
	Message() Resolver
	Timeline() Resolver
	Update() Resolver
}

type Resolver interface {
	Role(met map[string]string) (label.Label, error)
	Visibility(met map[string]string) (label.Label, error)
}
