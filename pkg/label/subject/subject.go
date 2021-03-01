package subject

import "github.com/venturemark/permission/pkg/label"

const (
	IsLoggedIn       label.Label = "subject:is-logged-in"
	IsResourceMember label.Label = "subject:is-resource-member"
	IsResourceOwner  label.Label = "subject:is-resource-owner"
	IsVentureMember  label.Label = "subject:is-venture-member"
)
