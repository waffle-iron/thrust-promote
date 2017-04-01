require_relative '../tools/helpers'


class VideoRender
	attr_reader :video_file

	def initialize(image_file, audio_file, video_file: "output.mp4")
		@image_file = image_file
		@audio_file = audio_file
		@video_file = video_file
	end

	def run
		audio_length = get_audio_length(@audio_file)

		puts "Running command: ffmpeg -loop 1 -f image2 -i #{@image_file} \
		 -i #{@audio_file} -c:v libx264 -c:a aac -strict experimental \
		 -b:a 192k -t #{audio_length} #{@video_file}"

		cmd = `ffmpeg -y -loop 1 -f image2 -i #{@image_file} \
		 -i "#{@audio_file}" -c:v libx264 -c:a aac -strict experimental \
		 -b:a 192k -t #{audio_length} #{@video_file}`
	end

	def self.run(image_file, audio_file, video_file)
		VideoRender.new(image_file, audio_file, video_file: video_file).run
	end
end



