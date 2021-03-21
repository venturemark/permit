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
	Invite() Resolver
	Message() Resolver
	Timeline() Resolver
	Update() Resolver
	User() Resolver
	Venture() Resolver
}

type Resolver interface {
	Role(met map[string]string) (string, error)
	Visibility(met map[string]string) (string, error)
}
