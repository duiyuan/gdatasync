package filesystem

import (
	"os"
	"path/filepath"
)

func ExcuteDir() (string, error) {
	excuteFile, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(excuteFile), nil
}

func SureLogDir(dir string) (rst string, err error) {
	excuteDir, err := ExcuteDir()

	if err != nil {
		return "", err
	}

	fullpath := filepath.Join(excuteDir, "logs", dir)
	if err = os.MkdirAll(fullpath, 0755); err != nil {
		return "", err
	}

	return fullpath, nil
}
