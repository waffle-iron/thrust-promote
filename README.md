# Thrust Workers
This will be HTTP triggered worker


### REST requests
* `/transcode/audio` will transcode an audio file and upload to GCS
* `/transcode/video` will render a video with and image and upload file to GCS
* `/social/send` will send a message to a social account
* `/events/send`will send an event to Songkick, BandsInTown
* `/release/send` will send to YT, Soundcloud, other places




### Audio Transcoder

Prerequisites:

* ffmpeg

* Google Cloud Platform downloader/uploader

  Once finished will upload to both GCS

### Video Renderer

Prerequisites:
* ffmpeg
* rvideo

Once finished will upload to both GCS and YT



### Send Requests

Prerequisites

* User tokens
* data



Postgres will be used to query and update user data