require 'resque'
require './lib/workers/downloads/audiodl'
require './lib/workers/jobs/prepare_image'

class PrepareAudio
	# Download Audio
	@queue = :audio

	def self.perform(metadata)
		audio_uri = metadata["audio_uri"]
		filename = AudioDL.download_from_gcs(audio_uri)
		metadata["audio_filepath"] = filename
        if metadata['video_render']
    		Resque.enqueue(PrepareImage, metadata)
        end
	end
end