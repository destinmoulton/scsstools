package lib

import (
	"scsstools/lib/actions"
)

var Actions = make(map[string]func([]string, string))

func init() {
	Actions["concat"] = actions.Concat
}

func IsAction(a string) bool {
	_, ok := Actions[a]
	return ok
}
