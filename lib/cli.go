package lib

import (
	"errors"
	"fmt"
	"os"
)

// Parse the cli args and perform basic validation
func ParseCLIArgs() (map[string]string, error) {
	p := make(map[string]string)
	if len(os.Args) < 3 {
		return nil, fmt.Errorf("usage: scsstools <actions.yml> <action>")
	}

	if _, err := os.Stat(os.Args[1]); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("cli: the file %s does not exist", os.Args[1])
	}
	p["file"] = os.Args[1]

	p["action"] = os.Args[2]

	return p, nil
}
