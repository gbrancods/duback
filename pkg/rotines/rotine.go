package rotines

import (
	"log"
	"os/exec"
	"time"
)

func BackupRotine() {
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
