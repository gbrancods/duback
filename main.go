package main

import (
	"duback/drive"
	"duback/tools"
)

func main() {

	// Gera as credentials
	drive.DriveCredenGen()

	// Gera o token
	drive.DriveTokenGen()

	// Lista os arquivos
	tools.BackupRotine()
}
