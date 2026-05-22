package amfile

import (
	"am/src/functionality/aliases"
	"am/src/types"
	"fmt"
	"os"
	"strings"
)

func WriteAlias(filePath string, entry types.AliasEntry) error {
	content, err := GetFileContent(filePath)
	if err != nil {
		return err
	}

	newLine := aliases.BuildLine(entry)
	lines := strings.Split(strings.TrimSpace(content), "\n")

	for i, line := range lines {
		existing, ok := aliases.ParseLine(line)
		if !ok {
			continue
		}
		if existing.Name == entry.Name {
			lines[i] = newLine
			return os.WriteFile(filePath, []byte(strings.Join(lines, "\n")+"\n"), 0644)
		}
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, newLine)
	return err
}
