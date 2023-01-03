package drive

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

type DBStruct struct {
	Port     string
	Name     string
	IP       string
	Password string
	User     string
}

func BackupGen(db DBStruct) error {

	d := time.Now().Day()
	m := time.Now().Month()
	y := time.Now().Year()

	date := fmt.Sprintf("%d-%d-%d", d, m, y)

	arg1 := fmt.Sprintf("--dbname=postgresql://%s:%s@%s:%s/postgres", db.Name, db.Password, db.IP, db.Port)

	arg2 := fmt.Sprintf("backup-%s-%s", db.Name, date)

	cmd := exec.Command("pg_dump", arg1, ">", arg2)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		// TODO - send email
	}

	compareString := string("192.168.55.250:" + db.Port + " - accepting connections")
	compareStringOut := string(out.String())

	// Compara se as strings
	if compareString == compareStringOut {
		return nil
	} else {
		return errors.New("erro ao conectar com o banco")
	}
}
