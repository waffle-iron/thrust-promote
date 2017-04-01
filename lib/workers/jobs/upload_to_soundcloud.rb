require 'resque'
require './lib/workers/uploads/soundcloud'


class UploadToSoundcloud
    @queue = :soundcloud

    def self.perform(metadata)
        video_filepath = metadata["video_filepath"]
        title = metadata["title"]
        video = SoundcloudUpload.upload_to_channel(video_filepath, title)
        metadata["yt_link"] = "https://www.youtube.com/watch?v=#{video.id}"
        metadata["status"] = "UPLOADED"

        # TODO store in release->track db
    end
end