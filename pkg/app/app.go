package app

import (
	"bytes"
	"duback/pkg/drive"
	"fmt"
	"os"
)

func App() (err error) {

	// TODO: Obter id da pasta automaticamente

	// Backup folder id
	fId := "1BuuDMi4Z9YrelWffjAqvAPOTSZisCsA3"
	// MimeType folder
	//fType := "inode/directory"
	// MimeType text
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
		if !e {

			//File Concat
			fCon := fmt.Sprintf("./backups/%s/backup-dbchat-%s", f, f)

			//Get the content file
			cf, err := os.ReadFile(fCon)
			if err != nil {
				fmt.Println(err)
			}

			// Convert byte slice to io.Reader
			cr := bytes.NewReader(cf)

			_, err = drive.CreateFile(cr, f, tType, fId)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}
