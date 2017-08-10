[![Stories in Ready](https://badge.waffle.io/ammoses89/thrust-promote.png?label=ready&title=Ready)](https://waffle.io/ammoses89/thrust-promote?utm_source=badge)
# Thrust Promote
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



### Workers



#### Concept

1. The workers are run by a `Machine` 
2. `Machine` runs like a server and spins the workers
3. Each `Worker` is pulls from the redis server `Broker`  struct 
4. The `Broker` struct converts the redis data JSON to a golang map that contains the task needed to be run
5. The `Worker` calls that task and the task runs to completion



#### Machine

#### Worker

#### Broker

#### Task

A task is a struct representation of:

* a `func` to be called to perform the task, and
* arguments to be passed to the corresponding `func`





