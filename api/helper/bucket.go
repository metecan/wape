package helper

import (
	"io"
	"os"

	storage_go "github.com/supabase-community/storage-go"
)

type Bucket struct {
	Name  string
	Raw   string
	Token string
}

func InitBucket() *Bucket {
	return &Bucket{
		Name:  os.Getenv("BUCKET_NAME"),
		Raw:   os.Getenv("BUCKET_RAW"),
		Token: os.Getenv("BUCKET_TOKEN"),
	}
}

func (b Bucket) Upload(fileName string, fileContent io.Reader) string {
	// Creating new client and access to Bucket
	client := storage_go.NewClient(b.Raw, b.Token, nil)

	// Uploading file to bucket
	client.UploadFile(b.Name, fileName, fileContent)

	// Getting the public url of file from from bucket
	publicAddr := client.GetPublicUrl(b.Name, fileName)

	// Returning the public url
	return publicAddr.SignedURL
}
