package ingress

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/venturemark/permission"
	"github.com/venturemark/permission/pkg/label"
	"github.com/venturemark/permission/pkg/label/action"
	"github.com/venturemark/permission/pkg/label/resource"
	"github.com/venturemark/permission/pkg/label/role"
	"github.com/venturemark/permission/pkg/label/visibility"
)

func Test_Ingress_Match_false(t *testing.T) {
	testCases := []struct {
		l []label.Label
	}{
		// Case 0
		{
			l: nil,
		},
		// Case 1
		{
			l: []label.Label{
				action.Search,
			},
		},
		// Case 2
		{
			l: []label.Label{
				action.Search,
				visibility.Private,
			},
		},
		// Case 3
		{
			l: []label.Label{
				visibility.Private,
			},
		},
		// Case 4
		{
			l: []label.Label{
				role.Member,
				action.Delete,
			},
		},
		// Case 5
		{
			l: []label.Label{
				role.Member,
				action.Update,
			},
		},
		// Case 6
		{
			l: []label.Label{
				action.Create,
				resource.Update,
			},
		},
		// Case 7
		{
			l: []label.Label{
				role.Member,
				action.Delete,
				visibility.Private,
				resource.Timeline,
			},
		},
		// Case 8
		{
			l: []label.Label{
				role.Any,
				action.Search,
				visibility.Private,
				resource.Message,
			},
		},
		// Case 9
		{
			l: []label.Label{
				role.Member,
				action.Update,
				visibility.Public,
				resource.Venture,
			},
		},
		// Case 10
		{
			l: []label.Label{
				role.Member,
				role.Member,
				role.Member,
			},
		},
		// Case 11
		{
			l: []label.Label{
				role.Owner,
				role.Owner,
				action.Update,
				action.Update,
				visibility.Public,
				resource.Venture,
			},
		},
		// Case 12
		{
			l: []label.Label{
				"",
				action.Create,
				visibility.Private,
				resource.Message,
			},
		},
		// Case 13
		{
			l: []label.Label{
				role.Owner,
				action.Update,
				visibility.Private,
				"",
			},
		},
		// Case 14
		{
			l: []label.Label{
				resource.Message,
				resource.Update,
				resource.Venture,
				visibility.Any,
				action.Create,
				role.Owner,
			},
		},
		// Case 15
		{
			l: []label.Label{
				resource.Timeline,
				visibility.Private,
				action.Search,
				role.Subject,
			},
		},
		// Case 16
		{
			l: []label.Label{
				resource.Invite,
				visibility.Public,
				action.Create,
				role.Member,
			},
		},
		// Case 17
		{
			l: []label.Label{
				resource.Invite,
				visibility.Private,
				action.Create,
				role.Member,
			},
		},
		// Case 18
		{
			l: []label.Label{
				resource.Message,
				visibility.Private,
				action.Create,
				role.Reader,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var ing permission.Ingress
			{
				c := Config{}

				ing, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			ok := ing.Match(tc.l...)
			if ok != false {
				t.Fatalf("\n\n%s\n", cmp.Diff(false, ok))
			}
		})
	}
}

func Test_Ingress_Match_true(t *testing.T) {
	testCases := []struct {
		l []label.Label
	}{
		// Case 0
		{
			l: []label.Label{
				role.Member,
				action.Create,
				visibility.Public,
				resource.Message,
			},
		},
		// Case 1
		{
			l: []label.Label{
				role.Owner,
				action.Update,
				visibility.Private,
				resource.Timeline,
			},
		},
		// Case 2
		{
			l: []label.Label{
				role.Any,
				action.Search,
				visibility.Public,
				resource.Message,
			},
		},
		// Case 3
		{
			l: []label.Label{
				role.Owner,
				action.Delete,
				visibility.Public,
				resource.Update,
			},
		},
		// Case 4
		{
			l: []label.Label{
				role.Member,
				action.Create,
				visibility.Private,
				resource.Message,
			},
		},
		// Case 5
		{
			l: []label.Label{
				action.Update,
				role.Owner,
				resource.Timeline,
				visibility.Private,
			},
		},
		// Case 6
		{
			l: []label.Label{
				resource.Message,
				visibility.Private,
				action.Create,
				role.Member,
			},
		},
		// Case 7
		{
			l: []label.Label{
				resource.Message,
				visibility.Any,
				action.Create,
				role.Owner,
			},
		},
		// Case 8
		{
			l: []label.Label{
				resource.Update,
				visibility.Private,
				action.Search,
				role.Owner,
			},
		},
		// Case 9
		{
			l: []label.Label{
				resource.Timeline,
				visibility.Private,
				action.Filter,
				role.Subject,
			},
		},
		// Case 10
		{
			l: []label.Label{
				resource.Message,
				visibility.Public,
				action.Filter,
				role.Subject,
			},
		},
		// Case 11
		{
			l: []label.Label{
				resource.Invite,
				visibility.Private,
				action.Create,
				role.Owner,
			},
		},
		// Case 12
		{
			l: []label.Label{
				resource.Message,
				visibility.Private,
				action.Search,
				role.Reader,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var ing permission.Ingress
			{
				c := Config{}

				ing, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			ok := ing.Match(tc.l...)
			if ok != true {
				t.Fatalf("\n\n%s\n", cmp.Diff(true, ok))
			}
		})
	}
}
