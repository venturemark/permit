package subject

import "github.com/venturemark/permission/pkg/label"

const (
	IsLoggedIn      label.Label = "subject:is-logged-in"
	IsResourceOwner label.Label = "subject:is-resource-owner"
	IsVentureMember label.Label = "subject:is-venture-member"
)
