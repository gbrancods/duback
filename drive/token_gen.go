package drive

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type driveToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expiry       string `json:"expiry"`
}

func DriveTokenGen() {

	credentialJsonStringB64 := os.Getenv("DRIVE_TOKEN")

	credentialJsonString, err := base64.StdEncoding.DecodeString(credentialJsonStringB64)
	if err != nil {
		panic(err)
	}

	var credentialJsonModel driveToken

	json.Unmarshal([]byte(credentialJsonString), &credentialJsonModel)

	content, err := json.Marshal(credentialJsonModel)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("token.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
