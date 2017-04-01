require './app'
require 'resque/server'


Resque::Server.use Rack::Auth::Basic do |username, password|
  username == 'thrust' && password == 'rules'
end

run Rack::URLMap.new \
  "/"       => Sinatra::Application,
  "/resque" => Resque::Server.new
