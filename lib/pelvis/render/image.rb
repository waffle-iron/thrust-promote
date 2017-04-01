require 'mini_magick'

class ImageRender
	attr_reader :image

	def initialize(image_file)
		@image_file = image_file
		@image = MiniMagick::Image.new(image_file)
	end

	def resize_to_full_screen
		@image.combine_options do |b|
			b.resize "1280x720!"
		end
		self.save
	end

	def save
		image_data = @image.to_blob
		File.delete(@image_file)
		new_file = File.open(@image_file, 'wb')
		new_file.write(image_data)
		new_file.close
	end

	def watermark(watermark_file)
		@image.combine_options do |c|
		  c.gravity 'SouthWest'
		  c.draw 'image Over 0,0 0,0 "#{watermark_file}"'
		end
	end
end