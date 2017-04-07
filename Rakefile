# Resque tasks
require 'yaml'
require 'resque/tasks'
require 'sinatra/activerecord'
require 'sinatra/activerecord/rake'
require 'active_record'
require './app'

include ActiveRecord::Tasks

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

root = File.expand_path '..', __FILE__
DatabaseTasks.env = ENV['ENV'] || 'development'
conf = File.join root, 'db/database.yml'
DatabaseTasks.database_configuration = YAML.load(File.read(conf))
DatabaseTasks.db_dir = File.join root, 'db'
DatabaseTasks.fixtures_path = File.join root, 'test/fixtures'
DatabaseTasks.migrations_paths = [File.join(root, 'db/migrate')]
DatabaseTasks.root = root

task :environment do
  ActiveRecord::Base.configurations = DatabaseTasks.database_configuration
  ActiveRecord::Base.establish_connection DatabaseTasks.env.to_sym
end

load 'active_record/railties/databases.rake'