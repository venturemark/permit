package ingress

import (
	"github.com/venturemark/permission/pkg/label"
	"github.com/venturemark/permission/pkg/label/action"
	"github.com/venturemark/permission/pkg/label/resource"
	"github.com/venturemark/permission/pkg/label/subject"
	"github.com/venturemark/permission/pkg/label/visibility"
)

var (
	policy = map[string][]label.Label{
		// Any subject can search any public resource.
		"po001": {action.Search, visibility.Public},

		// Any user being logged in can create a new venture.
		"po002": {subject.IsLoggedIn, action.Create, resource.Venture},

		// A resource owner can execute write operations on their resources.
		"po003": {subject.IsResourceOwner, action.Delete},
		"po004": {subject.IsResourceOwner, action.Update},

		// A venture member can search private resources for that venture.
		"po005": {subject.IsVentureMember, action.Search, visibility.Private},

		// A venture member can create resources for that venture.
		"po006": {subject.IsVentureMember, action.Create, resource.Timeline},
		"po007": {subject.IsVentureMember, action.Create, resource.Message},
		"po008": {subject.IsVentureMember, action.Create, resource.Update},
	}
)
