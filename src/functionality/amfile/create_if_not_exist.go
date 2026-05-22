package amfile

import (
	"os"
	"path/filepath"
)

const fileName = ".amf"

func ConfigDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "am"), nil
}

func CreateIfNotExist(dir string) (string, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	filePath := filepath.Join(dir, fileName)
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(filePath)
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	return filePath, nil
}
