package main

import (
    "log"
    "io/ioutil"
    "cloud.google.com/go/storage"
    "golang.org/x/net/context"
    "google.golang.org/api/option"
)


func DownloadFromGCS(urlPath string, filename string) string {
    ctx := context.Background()

    // set a project ID
    // projectID := "thrust"

    client, err := storage.NewClient(ctx, option.WithServiceAccountFile("thrust-5f3eaea7e015.json"))
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    // Set bucket name
    bucketName := "thrust-media"

    // Create bucket instance
    bucket := client.Bucket(bucketName)
    object := bucket.Object(urlPath)

    // create reader
    rc, err := object.NewReader(ctx)
    if err != nil {
        // reader
        panic(err)
    }
    data, err := ioutil.ReadAll(rc)
    rc.Close()
    if err != nil {
        // Handle error
        panic(err)
    }

    if err := ioutil.WriteFile(filename, data, 644); err != nil {
        panic(err)
    }

    return filename
}