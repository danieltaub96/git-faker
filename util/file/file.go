package util

import (
	"os"
	"path/filepath"
)

func AbsPath(path string) string {
	var pathAbs, _ = filepath.Abs(path)

	return pathAbs
}

func IsFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}