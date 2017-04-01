

class YoutubeDL
	def download_audio( url)
		filename = `youtube-dl --extract-audio #{url} --get-filename -o "/tmp/%(title)s.%(ext)s"`
		cmd = `youtube-dl -o "/tmp/%(title)s.%(ext)s" --extract-audio #{url}`
		filename.strip
	end

	def self.download_from_soundcloud( soundcloud_url )
		self.new.download_audio(soundcloud_url)
	end

	def self.download_from_youtube( youtube_url )
		self.new.download_audio(youtube_url)
	end
end


class AudioDL < YoutubeDL
end
