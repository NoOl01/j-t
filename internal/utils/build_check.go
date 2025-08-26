package utils

import (
	"os"
)

func BuildCheck() bool {
	file, err := os.Stat("./dist")
	if err != nil {
		return false
	}

	if !file.IsDir() {
		return false
	}

	return true
}
