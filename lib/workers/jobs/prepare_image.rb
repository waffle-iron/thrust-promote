require 'resque'
require './lib/workers/downloads/imagedl'
require './lib/workers/jobs/render_video'

class PrepareImage
	# Download Image
	@queue = :image

	def self.perform(metadata)
		id = metadata["id"]
		filename = "/tmp/#{id}_image.jpg"

		# download image to target file
		ImageDL.download_from_url(metadata["image_uri"], filename)
		metadata["image_filepath"] = filename

	    # resize image for Youtube	
		ImageRender.new(filename).resize_to_full_screen

		# Enqueue for prepare Audio
		Resque.enqueue(RenderVideo, metadata)
	end
end