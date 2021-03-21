package resolver

import (
	"testing"

	"github.com/venturemark/permission"
)

func Test_Resolver_Interface(t *testing.T) {
	var _ permission.Resolver = &Resolver{}
}
