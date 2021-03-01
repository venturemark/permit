package ingress

import "github.com/venturemark/permission/pkg/label"

func equal(a []label.Label, b []label.Label) bool {
	if a != nil && b == nil {
		return false
	}

	if a == nil && b != nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for _, x := range a {
		var f bool

		for _, y := range b {
			if x == y {
				f = true
				break
			}
		}

		if !f {
			return false
		}
	}

	return true
}
