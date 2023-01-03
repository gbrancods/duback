package main

func main() {

	// Gera as credentials
	drive.driveCredenGen()

	// Gera o token
	drive.driveTokenGen()

	drive.listFiles()
}
