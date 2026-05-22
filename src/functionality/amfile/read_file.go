package amfile

import (
	"os"
)

func GetFileContent(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	stringContent := string(content)
	return stringContent, nil
}
