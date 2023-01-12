package drive

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func CreateFile(content io.Reader, name, mimeType, parentId string) (*drive.File, error) {

	ctx := context.Background()

	client := GetClient()

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  []string{parentId},
	}

	var file *drive.File

	if mimeType == "text/plain" {
		file, err = srv.Files.Create(f).Media(content).Do()
	} else if mimeType == "application/vnd.google-apps.folder" {
		file, err = srv.Files.Create(f).Do()
	} else {
		fmt.Println("Inválid mime type")
	}

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	return file, nil
}
