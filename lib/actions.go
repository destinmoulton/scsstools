package lib

var Actions = make(map[string]func([]string, string))

func init() {
	Actions["concat"] = actionConcat
}

func IsAction(a string) bool {
	_, ok := Actions[a]
	return ok
}

// The concat action
func actionConcat(sources []string, dest string) {

}
