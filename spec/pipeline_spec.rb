require 'resque'
require './lib/workers/jobs/transcode_audio'
require './lib/workers/jobs/prepare_image'
require './lib/workers/jobs/prepare_audio'
require './lib/workers/jobs/render_video'
require './lib/workers/jobs/upload_to_youtube'



describe CheckSoundcloudJob do
	describe '#run' do
		before(:all) do
			Resque.inline = true
		end

		it "should run full pipeline" do
            payload = {
                :unstaged_audio_uri => 'test/unstaged/audio/test.flac'
                :transcodes => ['.mp3', '.wav'],
                :user_id => 1,
                :track_id => 1
            }
			TranscodeAudio.run(payload)
		end
	end
end