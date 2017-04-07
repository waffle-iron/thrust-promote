require './lib/workers/jobs/transcode_audio'

describe TranscodeAudio do
    describe '#run' do
        before(:all) do
            ENV["APP_ENV"] ||= "test"
            Resque.inline = true
        end

        it "should transcode audio end to end" do
            payload = {
                :source_path => 'test/unstaged/audio/test.flac',
                :target_path => 'test/staged/audio/test.mp3',
                :transcode_type => 'mp3',
                :track_id => 1
            }
            TranscodeAudio.run(payload)
        end
    end
end