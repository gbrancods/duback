package drive

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type db struct {
	date     string
	port     string
	dbName   string
	fileName string
	ip       string
	password string
	login    string
}

func backupGen(db db) int {

	//pg_dump --dbname=postgresql://postgres:000DHPdhpDHP000@192.168.55.250:5431/postgres > backup-dbchat-16-12-2022

	d := time.Now().Day()
	m := time.Now().Month()
	y := time.Now().Year()

	date := fmt.Sprintf("%d-%d-%d", d, m, y)

	arg1 := fmt.Sprintf("--dbname=postgresql://%s:%s@%s:%d/postgres", db.login, db.password, db.ip, db.port)
	arg2 := fmt.Sprintf("backup-%s-%s", db.dbName, date)

	cmd := exec.Command("pg_dump", arg1, ">", arg2)

	returnVar := -1
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	compareString := string("192.168.55.250:" + db.port + " - accepting connections")
	compareStringOut := string(out.String())
	resultCompare := strings.Compare(compareString+"\n", compareStringOut)

	if resultCompare == 0 {
		returnVar = 0
	} else {
		returnVar = 2
	}

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		returnVar = 1
	}

	return returnVar
}
