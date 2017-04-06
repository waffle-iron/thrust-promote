
class Uploader

    def upload_to_gcs( file_path, url )
      storage =  Google::Cloud::Storage.new(
        project: "thrust",
        keyfile: "../thrust-5f3eaea7e015.json"
      )
      bucket = storage.bucket "thrust-media"
      bucket.create_file file_path, url
    end

    def self.upload_to_gcs(file_path, gcs_file_path ) 
        self.new.upload_to_gcs(file_path, gcs_file_path)
    end
end

