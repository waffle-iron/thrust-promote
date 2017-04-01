require 'resque'
require './lib/pelvis/db/redb'
require './lib/pelvis/jobs/check_soundcloud'
require './lib/pelvis/jobs/prepare_image'
require './lib/pelvis/jobs/prepare_audio'
require './lib/pelvis/jobs/render_video'
require './lib/pelvis/jobs/upload_to_youtube'



describe CheckSoundcloudJob do
	describe '#run' do
		before(:all) do
			Resque.inline = true
		end

		xit "should run full pipeline" do
			CheckSoundcloudJob.run('kevin abstract - not on doasm 03', 'https://soundcloud.com/kevinabstract/not-on-doasm-03')
		end
	end
end