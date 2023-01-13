package drive

import (
	"fmt"
)

// Get root folder id
// Case folder not found, create
func GetBackupFolder() (fid string) {

	fn, fi := GDriveFiles()
	n := "bee-back"

	for i, f := range fn {
		if f == n {
			fid = fi[i]
			return
		}
	}

	ff, err := CreateFile(nil, n, FType, "")
	if err != nil {
		fmt.Println(err)
	}

	fid = ff.Id

	return
}
