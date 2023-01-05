package drive

import (
	"fmt"

	pg "github.com/habx/pg-commands"
)

type DBStruct struct {
	Port     int
	Name     string
	Host     string
	Password string
	User     string
}

func BackupGen(db DBStruct) error {

	dump, _ := pg.NewDump(&pg.Postgres{
		Host:     db.Host,
		Port:     db.Port,
		DB:       db.Name,
		Username: db.User,
		Password: db.Password,
	})

	dumpExec := dump.Exec(pg.ExecOptions{StreamPrint: false})
	if dumpExec.Error != nil {
		fmt.Println(dumpExec.Error.Err)
		fmt.Println(dumpExec.Output)
	} else {
		fmt.Println("Dump success")
		fmt.Println(dumpExec.Output)
	}

	return nil
}
