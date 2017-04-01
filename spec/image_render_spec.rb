require './lib/pelvis/render/image'


describe ImageRender do
	describe '#new' do
		before do
			@image_path = 'data/example.jpg'
		end

		it "should take 1 parameter" do
			@image_renderer = ImageRender.new(@image_path)
			expect(@image_renderer).to be_an_instance_of(ImageRender)
		end
	end

	describe '#resize_to_full_screen' do
		before do
			@image_path = 'data/dropbox-example.jpg'
			@image_renderer = ImageRender.new(@image_path)
		end

		it "should equal full screen height and width" do
			@image_renderer.resize_to_full_screen
			expect(@image_renderer.image.height).to eq(720)
			expect(@image_renderer.image.width).to eq(1280)
		end
	end
end