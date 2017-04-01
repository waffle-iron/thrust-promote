require 'resque'
require './lib/pelvis/uploads/soundcloud'
require './lib/pelvis/db/redb'


class UploadToSoundcloud
    @queue = :soundcloud

    def self.perform(metadata)
        video_filepath = metadata["video_filepath"]
        title = metadata["title"]
        video = SoundcloudUpload.upload_to_channel(video_filepath, title)
        metadata["yt_link"] = "https://www.youtube.com/watch?v=#{video.id}"
        metadata["status"] = "UPLOADED"

        # store in db
        ReDB.store_video(metadata)
    end
end