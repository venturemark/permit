package label

import "strings"

type Label string

func (l Label) Group() string {
	return strings.Split(string(l), ":")[0]
}

func (l Label) IsAny() bool {
	return l.Label() == "any"
}

func (l Label) Label() string {
	return strings.Split(string(l), ":")[1]
}
