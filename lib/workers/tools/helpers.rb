require 'mp3info'
require 'fastimage'


def get_audio_length(audio_file)
	Mp3Info.open(audio_file) do |mp3info|
	  seconds = mp3info.length
	  audio_length = Time.at(seconds).utc.strftime("%H:%M:%S.%L")
	end
end

def get_image_size(image_uri)
	size = FastImage.size(image_uri)
	size_hash = {
		:w => size[0],
		:h => size[1]
	}
	size_hash
end

def assert_image_type(image_uri)
	type = FastImage.type(image_uri)
	raise "Unsupported image type" unless type == :jpg
end