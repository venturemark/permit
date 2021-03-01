package permission

type Gateway interface {
	Ingress() Ingress
}

type Ingress interface {
	Match(l ...string) bool
}
