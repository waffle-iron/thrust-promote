require 'resque'
require './lib/workers/render/video'
require './lib/workers/jobs/upload_to_youtube'

class RenderVideo
	# Render Video
	@queue = :video

	def self.perform(metadata)
		image_filepath = metadata["image_filepath"]
		audio_filepath = metadata["audio_filepath"]

		id = metadata["id"]
		video_filepath = "/tmp/#{id}_video.mp4"

		# download image to target file
		VideoRender.run(image_filepath, audio_filepath, 
			video_filepath)

		metadata["video_filepath"] = video_filepath

		# Enqueue for youtube upload
		Resque.enqueue(UploadToYoutube, metadata)
	end
end