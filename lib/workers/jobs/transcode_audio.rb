require 'resque'
require 'securerandom'
require './lib/workers/tools/uploader'
require './lib/workers/tools/downloader'
require './lib/workers/render/audio'
require './lib/workers/models/track'

class TranscodeAudio
    def self.run(payload)
        payload['action'] = 'transcode_audio'
        payload[:id] = SecureRandom.uuid
        # go ahead and implement the full process
        # we'll break it down later
        audio_uri = payload[:source_path]
        # download file
        filename = Downloader.download_from_gcs(audio_uri)
        # transcode
        target_file = AudioRender.run(filename, payload[:transcode_type])
        # upload file
        Uploader.upload_to_gcs(target_file, payload[:target_path])
        # update db
        unless ENV["APP_ENV"] == "test"
            track = Track.find(payload[:track_id])
            if track
                track.files.create(
                  :staged => true,
                  :track_id => track.id,
                  :url_path => payload[:target_path],
                  :file_type => "audio_#{payload[:transcode_type]}"
                )
            end
        end
    end
end