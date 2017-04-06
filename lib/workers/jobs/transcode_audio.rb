require 'resque'
require 'securerandom'
require './lib/workers/tools/uploader'
require './lib/workers/tools/downloader'

class TranscodeAudio
    def self.run(payload)
        payload['action'] = 'transcode_audio'
        payload[:id] = SecureRandom.uuid
        # go ahead and implement the full process
        # we'll break it down later
        audio_uri = metadata["source_path"]
        # download file
        filename = AudioDL.download_from_gcs(audio_uri)
        # transcode
        target_file = AudioRender.run(filename, metadata["transcode_type"])
        # upload file
        Uploader.upload_to_gcs(target_file, metadata["target_path"])
        # update db
        track = Track.find(metadata["track_id"])
        track.files.create(
          :staged => true,
          :track_id => track.id,
          :url_path => metadata[:target_path],
          :file_type => "audio_#{metadata['transcode_type']}"
        )
    end
end