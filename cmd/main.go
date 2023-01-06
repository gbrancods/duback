package main

import (
	"duback/pkg/drive"
	"duback/pkg/rotines"
)

func main() {

	// Gera as credentials
	drive.CredenGen()

	// Gera o token
	drive.TokenGen()

	// Inicia rotina de backup
	rotines.BackupRotine()
}
