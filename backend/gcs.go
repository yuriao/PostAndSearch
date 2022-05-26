package main

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

const (
	BUCKET_NAME = "go_project_051922"
)

func saveToGCS(r io.Reader, objectName string) (string, error) {
	ctx := context.Background()
	fmt.Println("chk 00")
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}

	fmt.Println("chk 01")
	object := client.Bucket(BUCKET_NAME).Object(objectName)
	wc := object.NewWriter(ctx)
	if _, err := io.Copy(wc, r); err != nil {
		return "", err
	}
	fmt.Println("chk 02")
	if err := wc.Close(); err != nil {
		return "", err
	}
	fmt.Println("chk 03")
	if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", err
	}
	fmt.Println("chk 04")
	attrs, err := object.Attrs(ctx)
	if err != nil {
		return "", err
	}

	fmt.Printf("Image is saved to GCS: %s\n", attrs.MediaLink)
	return attrs.MediaLink, nil
}
