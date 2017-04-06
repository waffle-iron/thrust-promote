
class Downloader
    def download_audio( url)
        filename = `youtube-dl --extract-audio #{url} --get-filename -o "/tmp/%(title)s.%(ext)s"`
        cmd = `youtube-dl -o "/tmp/%(title)s.%(ext)s" --extract-audio #{url}`
        filename.strip
    end

    def download_from_gcs( url )
      storage =  Google::Cloud::Storage.new(
        project: "thrust",
        keyfile: "../thrust-5f3eaea7e015.json"
      )
      bucket = storage.bucket "thrust-media"
      bucket.file url
    end

    def self.download_from_soundcloud( soundcloud_url )
        self.new.download_audio(soundcloud_url)
    end

    def self.download_from_youtube( youtube_url )
        self.new.download_audio(youtube_url)
    end

    def self.download_from_gcs( gcs_file_path ) 
        self.new.download_from_gcs(gcs_file_path)
    end
end
