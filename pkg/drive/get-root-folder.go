package drive

import (
	"fmt"
)

// Get root folder id
// Case folder not found, create
func GetRootFolder() (rfid, dfid, sfid string) {

	fn, fi := GDriveFiles()
	n := "bee-back"

	for i, f := range fn {

		if f == n {
			rfid = fi[i]
			return
		}
	}

	rf, err := CreateFile(nil, n, FType, "")
	if err != nil {
		fmt.Println("Error on create root folder", err)
	}

	df, err := CreateFile(nil, "dev-env", FType, rf.Id)
	if err != nil {
		fmt.Println("Error on create dev-env folder", err)
	}

	sf, err := CreateFile(nil, "SIC", FType, rf.Id)
	if err != nil {
		fmt.Println("Error on create SIC folder", err)
	}

	rfid = rf.Id
	dfid = df.Id
	sfid = sf.Id

	return
}
