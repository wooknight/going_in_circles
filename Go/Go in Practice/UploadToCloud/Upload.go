package main

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"

	"cloud.google.com/go/storage"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

var BACKUP_BUCKET = "mts-ramesh-sandbox-1"

func main() {
	//create our backup bucket . this is where we will save all our files , inbound and outbound

	// credentials in env comes in escaped, need to unescape
	creds, err := ioutil.ReadFile("ramesh-sandbox-755a791f557a.json")
	if err != nil {
		log.Printf("Failed to ioen cert %v", err)
		return
	}
	// fmt.Print(string(creds))

	ctx := context.Background()

	mtsBackup, err := storage.NewClient(ctx, option.WithCredentialsJSON(creds))
	if err != nil {
		log.Printf("GCP UploadLocalToRemote failed . client not created\t%v", err)
		return
	}

	if err := mtsBackup.Bucket(BACKUP_BUCKET).Create(ctx, "ramesh-sandbox", nil); err != nil {
		log.Printf("could not create bucket %v", err)
		return

		// TODO: handle error.
	}
	mtsBackupAllFiles := mtsBackup.Bucket(BACKUP_BUCKET) //bucket handle
	l := "randi"
	f, err := os.Open(l)
	if err != nil {
		log.Printf("failed to open file %s %v", l, err)
		return
	}
	defer f.Close()
	r := "inbound"
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := mtsBackupAllFiles.Object(r + path.Base(l)).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		log.Printf("failed to copy file %s %v", l, err)
		return
	}
	if err := wc.Close(); err != nil {
		log.Printf("failed to close remote file %s %v", l, err)
		return
	}
	return
}
