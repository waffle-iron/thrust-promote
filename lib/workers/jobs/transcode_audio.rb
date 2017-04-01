require 'resque'
require 'securerandom'
require './lib/workers/jobs/prepare_audio'

class TranscodeAudio
    def self.run(payload)
        payload['action'] = 'transcode_audio'
        payload[:id] = SecureRandom.uuid
        payload[:status] = 'QUEUED'
        Resque.enqueue(PrepareAudio, payload)
    end
end