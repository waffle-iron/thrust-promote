require 'resque'
require './lib/pelvis/downloads/audiodl'
require './lib/pelvis/jobs/render_video'

class PrepareAudio
	# Download Audio
	@queue = :audio

	def self.perform(metadata)
		audio_uri = metadata["audio_uri"]
		filename = AudioDL.download_from_soundcloud(audio_uri)
		metadata["audio_filepath"] = filename
		Resque.enqueue(RenderVideo, metadata)
	end
end