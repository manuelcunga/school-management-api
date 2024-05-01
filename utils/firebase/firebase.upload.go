// pkg/util/upload.go
package firebase

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

var (
	storageClient *storage.Client
	wg            sync.WaitGroup
)

func InitFirebase() error {

	opt := option.WithCredentialsFile("config/firebase/Credentials/school-management-b6bd7-firebase-adminsdk-v1z7j-d7301ecd03.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return fmt.Errorf("error initializing Firebase App: %v", err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		return fmt.Errorf("error initializing Firebase Storage client: %v", err)
	}
	storageClient = client

	return nil
}

func UploadExcel(uploadControl <-chan struct{}, errorFileUpload chan<- string, filename string) error {
	defer wg.Done()

	<-uploadControl
	fmt.Printf("Uploading file %s\n", filename)

	ctx := context.Background()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		errorFileUpload <- filename
		return err
	}
	defer file.Close()

	bucket, err := storageClient.DefaultBucket()
	if err != nil {
		fmt.Printf("Error getting default bucket: %v\n", err)
		errorFileUpload <- filename
		return err
	}

	obj := bucket.Object(filename)
	wc := obj.NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		fmt.Printf("Error uploading file to Firebase Storage: %v\n", err)
		errorFileUpload <- filename
		return err
	}
	if err := wc.Close(); err != nil {
		fmt.Printf("Error closing writer: %v\n", err)
		errorFileUpload <- filename
		return err
	}

	fmt.Printf("File %s uploaded successfully\n", filename)
	return nil
}

func GetStorageClient() *storage.Client {
	return storageClient
}
