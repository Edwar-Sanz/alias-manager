package amfile

import (
	"am/src/functionality/aliases"
	"fmt"
	"os"
	"strings"
)

func DeleteAlias(filePath, name string) error {
	content, err := GetFileContent(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	filtered := lines[:0]
	found := false

	for _, line := range lines {
		entry, ok := aliases.ParseLine(line)
		if ok && entry.Name == name {
			found = true
			continue
		}
		filtered = append(filtered, line)
	}

	if !found {
		return fmt.Errorf("alias %q not found", name)
	}

	output := ""
	if len(filtered) > 0 {
		output = strings.Join(filtered, "\n") + "\n"
	}

	return os.WriteFile(filePath, []byte(output), 0644)
}
