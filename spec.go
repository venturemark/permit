package permission

import "github.com/venturemark/permission/pkg/label"

type Gateway interface {
	Ingress() Ingress
	Resource() Resource
}

type Ingress interface {
	Match(l ...label.Label) bool
}

type Resolver interface {
	Invite() Resource
	Message() Resource
	Timeline() Resource
	Update() Resource
	User() Resource
	Venture() Resource
}

type Resource interface {
	Role(met map[string]string) (string, error)
	Visibility(met map[string]string) (string, error)
}
