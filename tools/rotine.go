package tools

import (
	"duback/drive"
	"os"

	"github.com/robfig/cron/v3"
)

func BackupRotine() {

	c := cron.New()
	c.AddFunc("@every 1m", func() {
		bChat()
		bEco()
		bPay()
	})
}

func backups() {
	bChat()
	bEco()
	bPay()
}

func bChat() {

	var dbchat drive.DBStruct
	dbchat.IP = os.Getenv("CHAT_IP")
	dbchat.Port = os.Getenv("CHAT_PORT")
	dbchat.User = os.Getenv("CHAT_USER")
	dbchat.Password = os.Getenv("CHAT_PASSWORD")

	drive.BackupGen(dbchat)
}

func bEco() {

	var dbeco drive.DBStruct
	dbeco.IP = os.Getenv("ECO_IP")
	dbeco.Port = os.Getenv("ECO_PORT")
	dbeco.User = os.Getenv("ECO_USER")
	dbeco.Password = os.Getenv("ECO_PASSWORD")

	drive.BackupGen(dbeco)
}

func bPay() {

	var dbpay drive.DBStruct
	dbpay.IP = os.Getenv("PAY_IP")
	dbpay.Port = os.Getenv("PAY_PORT")
	dbpay.User = os.Getenv("PAY_USER")
	dbpay.Password = os.Getenv("PAY_PASSWORD")

	drive.BackupGen(dbpay)
}
