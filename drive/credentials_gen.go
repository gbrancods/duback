package drive

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

func DriveCredenGen() {

	credentialJsonStringB64 := os.Getenv("DRIVE_CREDENTIALS")

	credentialJsonString, err := base64.StdEncoding.DecodeString(credentialJsonStringB64)
	if err != nil {
		panic(err)
	}

	var credentialJsonModel driveCreden

	json.Unmarshal([]byte(credentialJsonString), &credentialJsonModel)

	content, err := json.Marshal(credentialJsonModel)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(content))

	err = os.WriteFile("credentials.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
