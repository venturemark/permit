package action

import "github.com/venturemark/permission/pkg/label"

const (
	Any    label.Label = "action:any"
	Create label.Label = "action:create"
	Delete label.Label = "action:delete"
	Filter label.Label = "action:filter" // filter is a secondary search mechanism
	Search label.Label = "action:search"
	Update label.Label = "action:update"
)
