require './lib/pelvis/render/image'
require './lib/pelvis/render/video'

describe VideoRender do
	describe '#new' do
		before do
			@audio_path = 'data/test.mp3'
			@image_path = 'data/dropbox-example.jpg'
		end

		it "should take 2 parameters" do
			@video_renderer = VideoRender.new(@image_path, @audio_path)
			expect(@video_renderer).to be_an_instance_of(VideoRender)
		end

		it "should have default video_file output.mp4"  do
			@video_renderer = VideoRender.new(@image_path, @audio_path)
			expect(@video_renderer.video_file).to eq("output.mp4")
		end
	end

	describe '#run' do
		before do
			@audio_path = 'data/test.mp3'
			@image_path = 'data/dropbox-example.jpg'
			ImageRender.new(@image_path).resize_to_full_screen
			@video_renderer = VideoRender.new(@image_path, @audio_path)
		end

		xit "should render video to output.mp4" do
			@video_renderer.run
			expect(File.exists?('output.mp4')).to eq(true)
		end
	end
end
