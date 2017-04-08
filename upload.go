import (
    "log"
    "cloud.google.com/go/storage"
    "golang.org/x/net/context"
)


func UploadFromGcs(string urlPath, string filename) int {
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
    wc := object.NewWriter(ctx)
    data, err := ioutil.ReadFile(filename)
    if _, err := wc.Write(data); err != nil {
        log.Fatalf("Failed to write file: %v", err) 
    }

    if err := wc.Close(); err != nil {
        log.Fatalf("Failed to save file: %v", err) 
    }

    return 1
}