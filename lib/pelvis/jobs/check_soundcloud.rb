require './lib/pelvis/db/redb'
require './lib/pelvis/configure'
require './lib/pelvis/jobs/prepare_image'
require 'soundcloud'
require 'instagram'
require 'securerandom'
require 'resque'
require 'json'

class CheckSoundcloudJob

	def check_adrian_soundcloud
		adrian_credentials = CFG['soundcloud']['adrian']	

		client_hash = {
			:client_id => CFG['soundcloud']['client_id'],
			:client_secret => CFG['soundcloud']['client_secret'],
			:username => adrian_credentials['username'],
			:password => adrian_credentials['password']
		}

		client = Soundcloud.new(client_hash)
		adrian_favorites = []
		all_favorites = client.get('/me/favorites', :limit => 500)
		for favorite in all_favorites
			if favorite.kind == 'track'
				favorite_hash = {
					:track_id => favorite.id,
					:created_at => favorite.created_at,
					:title => favorite.title,
					:url => favorite.permalink_url,
					:user => favorite.user.username
				}
				adrian_favorites.push favorite_hash
			end
		end

		conn = ReDB.get_conn
		# first time just store
		ReDB.get_table('users')
			.filter{|user| user["name"].eq("adrian")}
			.update({'favorites' => adrian_favorites})
			.run(conn)
	end

	def choose_photo
		conn = ReDB.get_conn
		image = ReDB.get_table('images').sample(1).run(conn)[0]
		image["uri"]
	end

	def self.run(title, track_url)
		cs = CheckSoundcloudJob.new
		# image_uri = cs.choose_photo
		metadata = {
			:id => SecureRandom.uuid,
			:user => '',
			:stage => '',
			:status => 'QUEUED',
			:audio_uri => track_url,
			:title => title,
			:uploaded_by => 'thrust',
			:yt_link => ''
		}
		Resque.enqueue(PrepareImage, metadata)
	end
end