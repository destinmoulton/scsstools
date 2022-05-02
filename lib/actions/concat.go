package actions

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// The concat action
func Concat(action TAction) error {
	sourceBase := action.SourcePath

	for _, sourceFile := range action.Sources {
		fp := filepath.Join(sourceBase, sourceFile)
		if _, err := os.Stat(fp); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("concat: the file %s does not exist", fp)
		}
	}

	return nil
}
