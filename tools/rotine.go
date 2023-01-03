package tools

import (
	"duback/drive"
	"os"

	"github.com/robfig/cron"
)

func backupRotine() {

	var dbchat drive.DBStruct
	dbchat.port = os.GetEnv("CHAT_PORT")

	var dbecommerce DBStruct
	var dbpay DBStruct

	c := cron.New()
	c.AddFunc("@every 1m", drive.CreateFile())

}
