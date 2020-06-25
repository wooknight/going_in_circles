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
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var backupBucket = "mts-ramesh-sandbox"
var mtsBackupAllFiles *storage.BucketHandle

func main() {
	//create our backup bucket . this is where we will save all our files , inbound and outbound

	// credentials in env comes in escaped, need to unescape
	creds, err := ioutil.ReadFile("ramesh-sandbox-81dfd482f50f.json")
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
	projectID := "ramesh-sandbox"

	// if err := mtsBackup.Bucket(BACKUP_BUCKET).Create(ctx, "ramesh-sandbox", nil); err != nil {
	// 	log.Printf("could not create bucket %v", err)
	// 	return

	// 	// TODO: handle error.
	// }
	it := mtsBackup.Buckets(ctx, projectID)
	log.Debug().Msgf("Buckets %v", it)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			//we did not find the bucket so create it
			if err := mtsBackup.Bucket(backupBucket).Create(ctx, projectID, nil); err != nil {
				log.Error().Err(err).Msgf("could not create bucket %s for project %s", backupBucket, projectID)
				return
			}
			mtsBackupAllFiles = mtsBackup.Bucket(backupBucket) //bucket handle
			break
		}
		if err != nil {
			log.Error().Err(err).Msgf("Error while looping for buckets in project %s", projectID)
			return
		}
		if battrs.Name == backupBucket {
			mtsBackupAllFiles = mtsBackup.Bucket(backupBucket) //bucket handle
			log.Debug().Msgf("bucket %s found for project %s", backupBucket, projectID)
			break
		}
	}
	mtsBackupAllFiles := mtsBackup.Bucket(backupBucket) //bucket handle
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
