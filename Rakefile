# Resque tasks
require 'yaml'
require 'resque/tasks'

namespace :resque do

  task :setup do
    require 'resque'
    require './lib/pelvis/jobs/check_soundcloud'
    require './lib/pelvis/jobs/prepare_image'
    require './lib/pelvis/jobs/prepare_audio'
    require './lib/pelvis/jobs/render_video'
    require './lib/pelvis/jobs/upload_to_youtube'
    REDIS_HOST_NAME = ENV['REDIS_PORT_6379_TCP_ADDR'] || '127.0.0.1'
    Resque.redis = "#{REDIS_HOST_NAME}:6379"
  end
end

namespace :rethinkdb do
  task :setup_db do
    require './setup_db.rb'
    setup_db
  end
end
