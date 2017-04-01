require 'soundcloud'


client_hash = {
    :client_id => CFG['soundcloud']['client_id'],
    :client_secret => CFG['soundcloud']['client_secret'],
    :redirect_uri => 'http://localhost:5000/callback'
}

client = Soundcloud.new(client_hash)