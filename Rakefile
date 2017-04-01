# Resque tasks
require 'yaml'
require 'resque/tasks'
require 'sinatra/activerecord'
require 'sinatra/activerecord/rake'
require './app'

namespace :resque do
  task :setup do
    require 'resque'
    require './lib/workers/jobs/prepare_image'
    require './lib/workers/jobs/prepare_audio'
    require './lib/workers/jobs/render_video'
    require './lib/workers/jobs/upload_to_youtube'
    REDIS_HOST_NAME = ENV['REDIS_PORT_6379_TCP_ADDR'] || '127.0.0.1'
    Resque.redis = "#{REDIS_HOST_NAME}:6379"
  end
end

