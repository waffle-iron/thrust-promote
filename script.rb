require './lib/pelvis/configure'
require './lib/pelvis/db/redb'
require './lib/pelvis/tools/dropbox'
require './lib/pelvis/tools/helpers'
require './lib/pelvis/jobs/check_soundcloud'
require './lib/pelvis/render/image'
# bulk insert data
require './lib/pelvis/render/video'



def empty_favorites
	ReDB.empty_favorites('adrian')
end


def get_random_image_uri
  conn = ReDB.get_conn
  cursor = ReDB.get_table('images').sample(1).run(conn)
  puts cursor[0]["uri"]
end


def test_dropbox
  puts Dropbox.new(CFG['dropbox']['access_token'])
  	.download_file("/ThrustImages/blur-blurred-city-2055.jpg", 'output.jpg')
end

def test_soundcloud
	CheckSoundcloudJob.new.check_adrian_soundcloud
end


def test_resize
	@image_path = 'data/dropbox-example.jpg'
	@image_renderer = ImageRender.new(@image_path)
	@image_renderer.resize_to_full_screen
end

def test_mp3_info
	@audio_path = 'data/test.mp3'
	@image_path = 'data/dropbox-example.jpg'
	@video_renderer = VideoRender.new(@image_path, @audio_path)
	@video_renderer.run
end


empty_favorites

