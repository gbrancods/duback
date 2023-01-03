package main

import "duback/drive"

func main() {

	// Gera as credentials
	drive.DriveCredenGen()

	// Gera o token
	drive.DriveTokenGen()

	// Lista os arquivos
	drive.ListFiles()
}
