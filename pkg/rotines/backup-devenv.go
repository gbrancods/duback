package rotines

import (
	"duback/pkg/app"
	"duback/pkg/drive"
	"log"
	"os/exec"
	"time"
)

func DevEnvRotine() {

	rf := drive.GetRootFolder()

	ticker := time.NewTicker(30 * time.Minute)

	for range ticker.C {
		call(rf)
	}
}

func call(rf string) {
	cmd := exec.Command("./scripts/backup-dev-env.sh")

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	app.DevSync(rf)
}
