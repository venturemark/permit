package permission

import "github.com/venturemark/permission/pkg/label"

type Gateway interface {
	Ingress() Ingress
}

type Ingress interface {
	Match(l ...label.Label) bool
}
