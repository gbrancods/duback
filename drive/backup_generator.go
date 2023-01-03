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
	DBName   string
	IP       string
	Password string
	Login    string
}

func BackupGen(db DBStruct) error {

	d := time.Now().Day()
	m := time.Now().Month()
	y := time.Now().Year()

	date := fmt.Sprintf("%d-%d-%d", d, m, y)

	arg1 := fmt.Sprintf("--dbname=postgresql://%s:%s@%s:%s/postgres", DBStruct.Login, DBStruct.Password, DBStruct.IP, DBStruct.Port)
	arg2 := fmt.Sprintf("backup-%s-%s", DBStruct.DBName, date)

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

	compareString := string("192.168.55.250:" + db.port + " - accepting connections")
	compareStringOut := string(out.String())

	// Compara se as strings
	if compareString == compareStringOut {
		return nil
	} else {
		return errors.New("erro ao conectar com o banco")
	}
}
