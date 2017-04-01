require 'typhoeus'
require 'json'

BASE_URL = "https://api.dropbox.com/2/"
CONTENT_URL = "https://api-content.dropbox.com/2/"

class Dropbox
	def initialize(access_token)
		@access_token = access_token
		@base_url = BASE_URL
	end

	def build_request(method, path, data, additional_headers:{})
		headers = {
			"Authorization" => "Bearer #{@access_token}",
		}

		if additional_headers
			headers.merge!(additional_headers)
		end

		if path == 'files/upload'
			data = File.open(data, 'r').read
		end

		request = Typhoeus::Request.new(
			@base_url + path,
			method: method,
			body: data,
			headers: headers
		)

		request
	end

	def download_file(filename, output_file)
		@base_url = CONTENT_URL
		path = 'files/download'
		additional_header = {
			"Dropbox-API-Arg" => {
				"path" => filename
			}.to_json,
			"Content-Type" => ""
		}
		data = nil
		request = build_request(:post, path, data, 
			additional_headers:additional_header)

		downloaded_file = File.open "#{output_file}", 'wb'
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
		response = request.run
	end

	def upload_file(local_file_path, dest_filename)
		@base_url = CONTENT_URL
		path = 'files/upload'
		additional_header = {
			"Dropbox-API-Arg" => {
				"path" => dest_filename,
				"mode" => "overwrite"
			}.to_json,
			"Content-Type" => "application/octet-stream"
		}
		data = "#{local_file_path}"
		request = build_request(:post, path, data, 
			additional_headers:additional_header)
		response = request.run
		response.body
	end

	def list_folder(folder)
		path = 'files/list_folder'

		additional_header {
		}

		data = {
			"path" => folder
		}.to_json

		request = build_request(:post, path, data, 
			additional_headers:additional_headers)
		request.run
	end

	def get_metadata(filename)
		path = 'files/get_metadata'

		additional_header {
		}

		data = {
			"path" => filename
		}.to_json
		
		response = build_request(:post, path, data, 
			additional_headers:additional_headers)
	end

	def get_current_account()
		path = 'users/get_current_account'

		additional_header {
		}

		data = "null"
		response = build_request(:post, path, data, 
			additional_headers:additional_header)
		response
	end
end