require './lib/pelvis/uploads/youtube'
require 'yt'

Yt.configure do |config|
  config.log_level = :debug
end

describe YoutubeUpload  do
	before(:all){ @yt_upload = YoutubeUpload.new}

	describe '#new' do

		it "should take 0 paramenters and be an instance of YoutubeUpload" do
			expect(@yt_upload).to be_an_instance_of(YoutubeUpload)
		end

		it "should create a youtube api account" do 
			expect(@yt_upload.account).to be_an_instance_of(Yt::Account)
		end
	end	

end