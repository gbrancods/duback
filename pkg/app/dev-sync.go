package app

import (
	"bytes"
	"duback/pkg/drive"
	"fmt"
	"os"
)

func DevSync(rf string) (err error) {

	// Get dev-env folders
	lf, err := getLocalFiles(DevF)
	if err != nil {
		return
	}

	//Get google drive files
	df, _ := drive.GDriveFiles()

	//Walks local files
	for _, f := range lf {

		//Define status
		s := false

		// Verify if file exist
		for _, d := range df {
			if f == d {
				s = true
			}
		}

		// Not found?
		if !s {

			// Create folder
			df, err := drive.CreateFile(nil, f, drive.FType, rf)
			if err != nil {
				fmt.Println(err)
			}

			// Walks names files concats and upload
			for _, fc := range devFilesConcats(f) {

				//Get the content file
				cf, err := os.ReadFile(fc)
				if err != nil {
					fmt.Println(err)
				}

				// convert byte slice to io.Reader
				r := bytes.NewReader(cf)

				//Create files
				_, err = drive.CreateFile(r, f, drive.TType, df.Id)
				if err != nil {
					fmt.Println(err)
				}
			}

		}
	}

	return nil
}
