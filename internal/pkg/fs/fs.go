package pkg

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

func SureLogDir(dir string) (rst bool, err error) {
	excuteDir, err := ExcuteDir()

	if err != nil {
		return false, err
	}

	fullpath := filepath.Join(excuteDir, dir)
	if err = os.MkdirAll(fullpath, 0755); err != nil {
		return false, err
	}

	return true, nil
}
