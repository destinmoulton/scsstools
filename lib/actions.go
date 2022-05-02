package lib

import (
	"fmt"
	"scsstools/lib/actions"
)

var Actions = make(map[string]func(actions.TAction) error)

func init() {
	Actions["concat"] = actions.Concat
}

// Run each action
// These are defined as the actions: in the yaml file
func RunActions(actions []actions.TAction, actionToRun string) error {
	for _, act := range actions {
		if act.Name == actionToRun {
			fmt.Printf("Running %s\n", act.Name)
			fmt.Println("---")
			fmt.Printf("%s\n", act.Description)
			fmt.Println("---")
			return Actions[act.Action](act)
		}
	}
	return nil
}
func IsAction(a string) bool {
	_, ok := Actions[a]
	return ok
}
