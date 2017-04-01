require 'clockwork'
require './lib/pelvis/jobs/check_soundcloud'
require './lib/pelvis/jobs/prepare_image'
require './lib/pelvis/jobs/prepare_audio'
require './lib/pelvis/jobs/render_video'
require './lib/pelvis/jobs/upload_to_youtube'

module Clockwork
	every(5.minutes, 'check_soundcloud.run'){ CheckSoundcloudJob.run }
end