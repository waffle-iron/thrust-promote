require "bundler/setup"
Bundler.require(:default)

require './lib/pelvis/jobs/check_soundcloud'
require 'json'

set :bind, '0.0.0.0'
set :port, 5000

get '/' do 
  File.read("public/index.html")
end

post '/add_from_soundcloud' do
	request_body = request.body.read
	payload = JSON.parse(request_body)
	CheckSoundcloudJob.run(payload['title'], payload['track_url'])
end

get '/connect_soundcloud' do
    # redirect to soundcloud
end

get '/soundcloud_likes' do
end

