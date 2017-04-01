require 'typhoeus'
require './lib/pelvis/tools/dropbox'
require './lib/pelvis/configure'


class ImageDL
	def initialize
	end

	def download_from_url(url, output_file)
		downloaded_file = File.open "#{output_file}", 'wb'
		request = Typhoeus::Request.new(url)
		request.on_headers do |response|
		  if response.code != 200
		    raise "Request failed"
		  end
		end
		request.on_body do |chunk|
		  downloaded_file.write(chunk)
		end
		request.on_complete do |response|
		  downloaded_file.close
		end
		request.run
	end

	def download_from_dropbox(filepath, output_file)
		Dropbox.new(CFG['dropbox']['access_token']).download_file(filepath, output_file)
	end

	def self.download_from_url(url, output_filename)
		ImageDL.new.download_from_url(url, output_filename)
	end
end