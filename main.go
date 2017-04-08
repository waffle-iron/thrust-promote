package main

import (
  "github.com/go-martini/martini"
)

func main() {
  m := martini.Classic()
  m.Use(martini.Static("public"))

  m.Group("/api", func(r martini.Router) {
    r.Post("/transcode/audio", CreateTranscodeAudioTask)
    r.Post("/transcode/video", CreateTranscodeVideoTask)
    r.Post("/social/send", CreateSocialSendTask)
    r.Post("/event/send", CreaetEventSendTask)
    r.Post("/release/send", CreateReleaseSendTask)
  })

  m.Run()
}
