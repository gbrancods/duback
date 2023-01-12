package app

import (
	"duback/pkg/drive"
	"fmt"
)

func App() (err error) {

	// TODO: Obter id da pasta automaticamente

	// Backup folder id
	fId := "1BuuDMi4Z9YrelWffjAqvAPOTSZisCsA3"

	// MimeType folder
	fType := "application/vnd.google-apps.folder"

	// MimeType text
	//tType := "text/plain"

	lf, err := getLocalFiles()
	if err != nil {
		return
	}

	df, _ := drive.GDriveFiles()

	// Walks local folders array
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

			//Create folder
			ff, err = drive.CreateFile(nil, f, fType, fId)
			if err != nil {
				fmt.Println(err)
			}

			//File path concat
			fcc := fmt.Sprintf("./backups/%s/backup-dbchat-%s", ff.Name, f)
			fce := fmt.Sprintf("./backups/%s/backup-dbeco-%s", ff.Name, f)
			fcp := fmt.Sprintf("./backups/%s/backup-dbpay-%s", ff.Name, f)

			fc = string[
				fcc,
				fce,
				fcp]

			//Get the content file
			cf, err := os.ReadFile(fCon)
			if err != nil {
				fmt.Println(err)
			}
			
			//Create file
			ff, err = drive.CreateFile(nil, f, fType, ff.Id)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}
