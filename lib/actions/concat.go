package actions

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// The concat action
func Concat(action TAction) error {
	sourceBase := action.SourcePath
	sourceBase, err := filepath.Abs(sourceBase)
	if err != nil {
		log.Fatalf("concat: unable to get absolute directory %s", sourceBase)
	}
	fmt.Printf("Concatenating files from sourcepath %s\n", sourceBase)

	var lines []string
	for _, sourceFile := range action.Sources {
		fp := filepath.Join(sourceBase, sourceFile)
		if _, err := os.Stat(fp); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("concat: the file %s does not exist", fp)
		}

		lines = append(lines, concatImports(fp)...)
	}

	file, err := os.OpenFile(action.Destination, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("concat: failed creating file: %s", err)
	}

	defer file.Close()
	datawriter := bufio.NewWriter(file)

	for _, data := range lines {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()

	return nil
}

// Handle @import statements
func concatImports(fpath string) []string {
	fmt.Printf("Concatenating SCSS file: %s\n", fpath)

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatalf("concat: failed to open file %s", fpath)
	}
	defer file.Close()

	basepath := filepath.Dir(fpath)

	scanner := bufio.NewScanner(file)

	impregex := regexp.MustCompile("@import[^'\"]+?['\"](?P<file>.+?)['\"];?")
	lines := []string{
		"// ---",
		"// " + fpath,
		"// ---",
		"",
	}
	for scanner.Scan() {
		line := scanner.Text()
		matches := impregex.FindStringSubmatch(line)
		if len(matches) > 0 {
			fileindex := impregex.SubexpIndex("file")
			subfull := matches[fileindex]
			subpath := filepath.Dir(subfull)
			subfile := filepath.Base(subfull)

			// SCSS imports don't require extension
			ext := filepath.Ext(subfile)
			if ext == "" {
				subfile = fmt.Sprintf("%s.scss", subfile)
			}

			trypath := filepath.Join(basepath, subpath, subfile)
			if _, err := os.Stat(trypath); errors.Is(err, os.ErrNotExist) {
				fmt.Printf("%s does not exist, trying with _ prefix...\n", trypath)
				trypath = filepath.Join(basepath, subpath, "_"+subfile)
				if _, err := os.Stat(trypath); errors.Is(err, os.ErrNotExist) {
					log.Fatalf("concat: %s file does not exist either\n", trypath)
				}
			}
			lines = append(lines, concatImports(trypath)...)
		} else {
			lines = append(lines, line)
		}

	}
	return lines
}
