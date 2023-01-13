package app

import (
	"os"
)

func getLocalFiles(path string) (lf []string, err error) {

	f, err := os.Open(path)
	if err != nil {
		return
	}

	files, err := f.Readdir(0)
	if err != nil {
		return
	}

	// Get all folders
	for _, v := range files {
		lf = append(lf, v.Name())
	}

	return
}
