require 'soundcloud'
require './lib/workers/configure'


adrian_credentials = CFG['soundcloud']['adrian']    

client_hash = {
    :client_id => CFG['soundcloud']['client_id'],
    :client_secret => CFG['soundcloud']['client_secret'],
    :username => adrian_credentials['username'],
    :password => adrian_credentials['password']
}

class SoundcloudUpload
    attr_reader :account

    def initialize
        @account = Soundcloud.new(client_hash)
    end

    def upload_to_channel(audio_uri, title, tags: ["music"], privacy_status: "private")
        @account.post('/tracks', :track => {
          :title => title,
          :asset_data => File.new(audio_uri, 'rb')
        })
    end

    def self.upload_to_channel(audio_uri, title, tags: ["music"], privacy_status: "private")
        SoundcloudUpload.new.upload_to_channel(audio_uri, title, 
            tags: tags, privacy_status: privacy_status)
    end
end