package gateway

import (
	"testing"

	"github.com/venturemark/permission"
)

func Test_Gateway_Interface(t *testing.T) {
	var _ permission.Gateway = &Gateway{}
}
