require 'db/redis'
=begin

Stages:
	- Download Image
	- Image Render (Resize and Watermark)
	- Download Audio
	- Video Render 
	- Upload to s3
	- Upload to YT


metadata = {
	:id => '',
	:audio_uri => '',
	:image_uri => '',
	:artist_name => '',
	:title => ''
	:video_file => '',
	:yt_link => ''
}

=end


