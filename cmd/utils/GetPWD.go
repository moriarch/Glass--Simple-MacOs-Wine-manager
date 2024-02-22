package utils

import (
	"os"
	"path/filepath"
)

func GetPWD() string {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path := filepath.Dir(ex)
	return path

}
