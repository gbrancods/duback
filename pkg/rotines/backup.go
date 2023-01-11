package rotines

import (
	"duback/pkg/app"
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
	cmd := exec.Command("./scripts/backup.sh")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	app.App()
}
