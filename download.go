import (
    "log"
    "cloud.google.com/go/storage"
    "golang.org/x/net/context"
)


func DownloadFromGcs(string urlPath, string filename) string {
    ctx := context.Background()

    // set a project ID
    projectID := "thrust"

    client, err := storage.NewClient(ctx)
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
    }
    data, err := ioutil.ReadAll(rc)
    rc.Close()
    if err != nil {
        // Handle error
    }

    if err := ioutil.WriteFile(filename, data); err != nil {
        panic(err)
    }

    return filename

}