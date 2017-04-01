require 'aws-sdk'
require 'yaml'

conf = YAML::load(File.open('config.yml'))

class S3Upload
	def initialize
		@s3 = AWS::S3.new(access_key_id: conf['access_key'], 
			secret_access_key: conf['secret_key'], 
			region: 'us-standard')
	end

	def get_bucket(bucket_name)
		@s3.buckets[bucket_name]
	end

end
