package tools

import (
	"duback/drive"
	"fmt"
	"os"
	"strconv"
	"time"
)

func BackupRotine() {
	//Define ticker
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		bChat()
		bEco()
		bPay()
	}
}

func bChat() {

	p := os.Getenv("CHAT_PORT")

	fmt.Println("C P", p)

	pi, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Erro de conversão 'chat'", err)
	}

	var dbchat drive.DBStruct
	dbchat.Host = os.Getenv("CHAT_IP")
	dbchat.Port = pi
	dbchat.User = os.Getenv("CHAT_USER")
	dbchat.Password = os.Getenv("CHAT_PASSWORD")
	dbchat.Name = os.Getenv("CHAT_DBNAME")

	err = drive.BackupGen(dbchat)
	if err != nil {
		fmt.Println("Erro ao gerar backup", err)
	}
}

func bEco() {

	p := os.Getenv("ECO_PORT")

	fmt.Println("E P", p)

	pi, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Erro de conversão 'dbeco'", err)
	}

	var dbeco drive.DBStruct
	dbeco.Host = os.Getenv("ECO_IP")
	dbeco.Port = pi
	dbeco.User = os.Getenv("ECO_USER")
	dbeco.Password = os.Getenv("ECO_PASSWORD")
	dbeco.Name = os.Getenv("ECO_DBNAME")

	err = drive.BackupGen(dbeco)
	if err != nil {
		fmt.Println("Erro ao gerar backup", err)
	}
}

func bPay() {

	p := os.Getenv("PAY_PORT")

	fmt.Println("P P", p)

	pi, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Erro de conversão 'dbpay'", err)
	}

	var dbpay drive.DBStruct
	dbpay.Host = os.Getenv("PAY_IP")
	dbpay.Port = pi
	dbpay.User = os.Getenv("PAY_USER")
	dbpay.Password = os.Getenv("PAY_PASSWORD")
	dbpay.Name = os.Getenv("PAY_DBNAME")

	err = drive.BackupGen(dbpay)
	if err != nil {
		fmt.Println("Erro ao gerar backup", err)
	}
}
