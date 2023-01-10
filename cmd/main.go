package main

import (
	"duback/pkg/drive"
	"duback/pkg/rotines"
)

func main() {

	// Create credentials
	drive.CredenGen()

	// Create token
	drive.TokenGen()

	// Start backup rotine
	rotines.BackupRotine()
}
