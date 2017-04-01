require 'resque'
require './lib/pelvis/downloads/imagedl'
require './lib/pelvis/render/image'
require './lib/pelvis/jobs/prepare_audio'

class PrepareImage
	# Download Image
	@queue = :image

	def self.choose_photo
		conn = ReDB.get_conn
		image = ReDB.get_table('images').sample(1).run(conn)[0]
		image["uri"]
		# TODO marke image retrieved as in use
	end

	def self.perform(metadata)
		image_uri = self.choose_photo
		id = metadata["id"]
		filename = "/tmp/#{id}_image.jpg"

		# download image to target file
		ImageDL.download_from_url(image_uri, filename)
		metadata["image_uri"] = image_uri
		metadata["image_filepath"] = filename

	    # resize image for Youtube	
		ImageRender.new(filename).resize_to_full_screen

		# Enqueue for prepare Audio
		Resque.enqueue(PrepareAudio, metadata)
	end
end