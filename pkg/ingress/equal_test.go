package ingress

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/permission/pkg/label"
)

func Test_ingress_empty_false(t *testing.T) {
	testCases := []struct {
		a []label.Label
		b []label.Label
	}{
		// Case 0
		{
			a: nil,
			b: []label.Label{},
		},
		// Case 1
		{
			a: []label.Label{},
			b: nil,
		},
		// Case 2
		{
			a: []label.Label{"a"},
			b: []label.Label{},
		},
		// Case 3
		{
			a: []label.Label{"a"},
			b: []label.Label{"b"},
		},
		// Case 4
		{
			a: []label.Label{"a", "1", ""},
			b: []label.Label{"b"},
		},
		// Case 5
		{
			a: []label.Label{"a", "1", ""},
			b: []label.Label{"", "b"},
		},
		// Case 6
		{
			a: []label.Label{"", "", ""},
			b: []label.Label{"", ""},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ok := equal(tc.a, tc.b)
			if ok != false {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, false))
			}
		})
	}
}

func Test_ingress_empty_true(t *testing.T) {
	testCases := []struct {
		a []label.Label
		b []label.Label
	}{
		// Case 0
		{
			a: nil,
			b: nil,
		},
		// Case 1
		{
			a: []label.Label{},
			b: []label.Label{},
		},
		// Case 2
		{
			a: []label.Label{"a"},
			b: []label.Label{"a"},
		},
		// Case 3
		{
			a: []label.Label{"a", "b"},
			b: []label.Label{"a", "b"},
		},
		// Case 4
		{
			a: []label.Label{"b", "a"},
			b: []label.Label{"a", "b"},
		},
		// Case 5
		{
			a: []label.Label{"1", "foo", ""},
			b: []label.Label{"foo", "1", ""},
		},
		// Case 6
		{
			a: []label.Label{"", "", ""},
			b: []label.Label{"", "", ""},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ok := equal(tc.a, tc.b)
			if ok != true {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, true))
			}
		})
	}
}
