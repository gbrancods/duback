package app

import (
	"duback/pkg/drive"
	"fmt"
	"io"
	"io/ioutil"
)

func App() (err error) {

	// Backup folder id
	fId := "1BuuDMi4Z9YrelWffjAqvAPOTSZisCsA3"
	// MimeType folder
	// fType := "application/vnd.google-apps.folder"
	// fType := "inode/directory"
	// MimeType text
	// tType := "application/vnd.google-apps.file"
	tType := "text/plain"

	lf, err := getLocalFiles()
	if err != nil {
		return
	}

	df, _ := drive.GDriveFiles()

	// Walks local file array
	for _, f := range lf {

		e := false

		// Verify if file exist
		for _, d := range df {
			if f == d {
				e = true
			}
		}

		// Case not file exist
		if e == false {

			//File Concat
			fCon := fmt.Sprintf("./backups/%s/backup-dbchat-%s", f, f)

			//Get the content file
			cf, err := ioutil.ReadFile(fCon)
			if err != nil {
				fmt.Println(err)
			}

			rcf, err := io.Reader(cf)
			if err != nil {
				fmt.Println(err)
			}

			_, err = drive.CreateFile(rcf, f, tType, fId)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}
