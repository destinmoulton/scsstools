package actions

type TAction struct {
	Name        string
	Description string
	Action      string
	Watch       bool
	SourcePath  string
	Sources     []string
	Destination string
}
