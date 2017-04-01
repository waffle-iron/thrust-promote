require "bundler/setup"
Bundler.require(:default)

require './lib/workers/jobs/transcode_audio'
require './lib/workers/jobs/transcode_video'

require 'json'

set :bind, '0.0.0.0'
set :port, 5000

get '/' do 
  File.read("public/index.html")
end

post '/transcode/audio' do
    """
        payload:
            - source_url
            - target_url
            - transcode_type
            - user_id
    """
    request_body = request.body.read
    payload = JSON.parse(request_body)
    TranscodeAudioJob.run(payload)
end

post '/transcode/video' do
    """
        payload:
            - source_url
            - target_url
            - image_url
            - user_id
    """
    request_body = request.body.read
    payload = JSON.parse(request_body)
end

post '/social/send' do
    """
        payload:
            - access_token
            - message
            - other_params
            - user_id
    """
    request_body = request.body.read
    payload = JSON.parse(request_body)
end

post '/event/send' do
    """
        payload:
            - access_token
            - message
            - other_params
            - user_id
    """
    request_body = request.body.read
    payload = JSON.parse(request_body)
end

post '/release/send' do
    """
        payload:
            - access_token
            - url
            - other_params
            - user_id
    """
    request_body = request.body.read
    payload = JSON.parse(request_body)
end