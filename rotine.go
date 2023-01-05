package main

import (
	"log"
	"os/exec"
	"time"
)

func backupRotine() {
	//Define ticker
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		callB()
	}
}

func callB() {
	cmd := exec.Command("./backup.sh")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
