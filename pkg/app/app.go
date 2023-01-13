package app

import (
	"duback/pkg/drive"
	"fmt"
	"os"
)

func App() (err error) {

	// Backup folder id
	fid := drive.GetBackupFolder()

	// Get SIC files
	slf, err := getLocalFiles(SicF)
	if err != nil {
		return
	}

	// Get dev-env folders
	// dlf, err := getLocalFiles(DevF)
	// if err != nil {
	// 	return
	// }

	df, _ := drive.GDriveFiles()

	// Walks local folders array
	// dev_env
	for _, f := range slf {

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
			ff, err := drive.CreateFile(nil, f, drive.FType, fid)
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
