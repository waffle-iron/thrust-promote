require './lib/pelvis/jobs/check_soundcloud'
require './lib/pelvis/db/redb'


describe CheckSoundcloudJob do
	describe '#check_adrian_soundcloud' do
		before do
			ReDB.empty_favorites('adrian')
			@sc_job = CheckSoundcloudJob.new
			@sc_job.check_adrian_soundcloud
		end	

		it "should retrieve and store favorites" do
			conn = ReDB.get_conn
			stored_favorites = ReDB.get_table('users')
				.filter({'name' => 'adrian'})
				.pluck('favorites')
				.run(conn).to_a
			expect(stored_favorites.length).to be > 0
		end
	end
end