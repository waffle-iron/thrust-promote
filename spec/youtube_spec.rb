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

	# describe '#upload_to_channel' do
	# 	before do 
	# 		@title = 'Kevin Abstract - not on doasm 03'
	# 		@yt_upload.upload_to_channel('data/output.mp4', @title)
	# 		@video = @yt_upload.account.videos.where(title: @title).first
	# 	end


	# 	it "should upload to test channel" do
	# 		expect(@yt_upload.account.channel.title).to eq("Thrust Music")
	# 	end	

	# 	it "should upload with requested title" do
	# 		expect(@video).not_to eq(nil)
	# 		expect(@video.title).to eq(@title)
	# 	end

	# 	it "should upload video as private to channel" do
	# 		expect(@video.privacy_status).to eq('private')
	# 	end

	# 	after {@video.delete}
	# end
	
end