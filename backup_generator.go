package main

type db struct {
	date string
	port string
	dbName string
	fileName string
	ip string
	password string
	login string
}

func backupGen(db db) {
	
	pg_dump --dbname=postgresql://postgres:000DHPdhpDHP000@192.168.55.250:5431/postgres > backup-dbchat-16-12-2022

	command := Sprintf("pg_dump --dbname=postgresql://%s:%s@%s:%d/%s > backup-%s-%s",db.login, db.password, db.ip, db.port, )

	cmd := exec.Command("pg_dump", "--dbname=postgresql://%s:%s@%s:%d/%s", ">", "backup-dbchat-16-12-2022")

	returnVar := -1
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	compareString := string("192.168.55.250:" + port + " - accepting connections")
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