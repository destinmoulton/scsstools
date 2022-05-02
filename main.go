// Provides SCSS tools.
// Usage: scsstools <actions.yml> <action>
// Allowed actions are:
//   - concat
//		concatenates `sources` array of scss files into
//      `destination` file
//      follows @import statements insuring full inclusion
//   - compile
//      compiles `sources` array of scss files

package main

import (
	"fmt"
	"os"
	"scsstools/lib"
)

func main() {
	args, err := lib.ParseCLIArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	actions, err := lib.ParseYAMLActions(args["file"])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(actions)
}
