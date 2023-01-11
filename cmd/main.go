package main

import (
	"duback/pkg/drive"
	"duback/pkg/rotines"
)

func main() {

	// Create credentials
	drive.CredenGen()

	// Create token
	drive.GetClient()

	// Start backup rotine
	rotines.BackupRotine()
}
