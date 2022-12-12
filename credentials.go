package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type driveCreden struct {
	Installed struct {
		ClientID                string   `json:"client_id"`
		ProjectID               string   `json:"project_id"`
		AuthURI                 string   `json:"auth_uri"`
		TokenURI                string   `json:"token_uri"`
		AuthProviderX509CertURL string   `json:"auth_provider_x509_cert_url"`
		ClientSecret            string   `json:"client_secret"`
		RedirectUris            []string `json:"redirect_uris"`
	} `json:"installed"`
}

func driveCredenGen() {

	credentialJsonStringB64 := os.Getenv("DRIVE_CREDENTIALS")

	credentialJsonString, err := base64.StdEncoding.DecodeString(credentialJsonStringB64)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("credentials.json", credentialJsonString, 0600)
	if err != nil {
		log.Fatal(err)
	}
}

type driveToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expiry       string `json:"expiry"`
}

func driveTokenGen() {

	credentialJsonStringB64 := os.Getenv("DRIVE_TOKEN")

	credentialJsonString, err := base64.StdEncoding.DecodeString(credentialJsonStringB64)
	if err != nil {
		panic(err)
	}

	var credentialJsonModel driveToken

	json.Unmarshal([]byte(credentialJsonString), &credentialJsonModel)

	//--------------
	content, err := json.Marshal(credentialJsonModel)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("token.json", content, 0600)
	if err != nil {
		log.Fatal(err)
	}
}
