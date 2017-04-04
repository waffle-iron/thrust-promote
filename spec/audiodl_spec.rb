require './lib/workers/downloads/audiodl'

describe AudioDL do

	describe "#new" do
		before do
			@audio_downloader = AudioDL.new	
		end

	    it "takes zero parameters and returns an AudioDL object" do
	        expect(@audio_downloader).to be_an_instance_of(AudioDL)
	    end
	end

	describe "#download_audio" do
		before do
			url = "https://soundcloud.com/kevinabstract/not-on-doasm-03"
			@audio_filename = AudioDL.new.download_audio(url)
		end

		it "takes one parameter and returns a string" do
			expect(@audio_filename).to be_an_instance_of(String)
		end

		it "takes one parameter and returns a filename with song title" do
			expect(@audio_filename).to eq('/tmp/not on doasm 03 (prod. albert gordon).mp3')
		end
	end

end