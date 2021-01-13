package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"backend/args"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
)

/*
GetSignedUrl(filename string)
*/
func GetSignedURL(filename string) string {

	saKey, err := ioutil.ReadFile(args.Args.PrivateKey)
	if err != nil {
		log.Println(err)
	}

	cfg, err := google.JWTConfigFromJSON(saKey)
	if err != nil {
		log.Println(err)
	}

	bucket := args.Args.S3Bucket
	method := "GET"
	expires := time.Now().Add(time.Minute * 1)
	url, err := storage.SignedURL(bucket, filename, &storage.SignedURLOptions{
		GoogleAccessID: cfg.Email,
		PrivateKey:     cfg.PrivateKey,
		Method:         method,
		Expires:        expires,
		Scheme:         storage.SigningSchemeV4,
	})
	if err != nil {
		fmt.Println("\n \n GCS Error " + err.Error())
	}
	fmt.Println(url)
	return url
}
