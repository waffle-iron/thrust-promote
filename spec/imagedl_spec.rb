require './lib/pelvis/downloads/imagedl'

describe ImageDL do
	describe "#new" do
		before do
			@image_downloader = ImageDL.new
		end

	    it "takes zero parameters and returns an ImageDL object" do
	        expect(@image_downloader).to be_an_instance_of(ImageDL)
	    end
	end

	describe "#download_from_url" do
		before do
			@image_downloader = ImageDL.new	
		end

		it "saves file to given path" do
			@image_path = "data/example.jpg" 
			@image_downloader.download_from_url "https://placekitten.com/g/200/300", @image_path
			expect(File.exists?(@image_path)).to eq(true)
		end
	end

	describe "#download_from_dropbox" do
		before do
			@image_downloader = ImageDL.new	
		end

		it "saves dropbox file to given path" do
			@image_path = "data/dropbox-example.jpg" 
			@image_downloader.download_from_dropbox "/ThrustImages/blur-blurred-city-2055.jpg", @image_path
			expect(File.exists?(@image_path)).to eq(true)
		end
	end
end