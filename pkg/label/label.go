package label

import "strings"

type Label string

func (l Label) Equal(v Label) bool {
	return string(l) == string(v)
}

func (l Label) Group() string {
	return strings.Split(string(l), ":")[0]
}

func (l Label) IsAny() bool {
	return strings.Split(string(l), ":")[1] == "any"
}
