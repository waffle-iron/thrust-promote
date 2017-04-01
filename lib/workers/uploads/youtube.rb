require 'yt'
require_relative '../configure'

Yt.configure do |config|
  config.client_id = CFG['youtube']['client_id']
  config.client_secret = CFG['youtube']['client_secret']
end

class YoutubeUpload
	attr_reader :account

	def initialize
		@account = Yt::Account.new refresh_token: CFG['youtube']['refresh_token']
	end

	def upload_to_channel(video_uri, title, tags: ["music"], privacy_status: "private")
		@account.upload_video(video_uri, title: title, tags: tags, 
			privacy_status:privacy_status)
	end

	def self.upload_to_channel(video_uri, title, tags: ["music"], privacy_status: "private")
		YoutubeUpload.new.upload_to_channel(video_uri, title, tags: tags, privacy_status: privacy_status)
	end
end