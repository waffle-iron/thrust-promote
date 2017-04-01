require 'resque'
require './lib/workers/uploads/youtube'
require './lib/workers/models/track'


class UploadToYoutube
	@queue = :youtube

	def self.perform(metadata)
		video_filepath = metadata["video_filepath"]
		title = metadata["title"]
		video = YoutubeUpload.upload_to_channel(video_filepath, title)
		metadata["yt_link"] = "https://www.youtube.com/watch?v=#{video.id}"
		metadata["status"] = "UPLOADED"

		# update model
	end
end