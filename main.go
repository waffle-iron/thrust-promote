package main

import (
	"github.com/go-martini/martini"
	"log"
)

const WORKER_COUNT = 5;

func main() {
	taskMap := map[string]interface{}{
		"transcode_audio": TranscodeAudio,
		"transcode_video": TranscodeVideo,
		"social_send":     SocialSend,
		"event_send":      EventSend,
		"release_send":    ReleaseSend,
	}
	machine := &Machine{}
	db := &Postges{}
	log.Info("Registering Tasks...")
	machine.RegisterTasks(taskMap)
	log.Info("Launching Workers...")
	if err := machine.LaunchWorkers(WORKER_COUNT); err != nil {
		log.Fatalf("Failed to launch workers: %v", err)
		panic(err)
	}

	m := martini.Classic()
	m.Use(martini.Static("public"))

	m.Map(machine)
	m.Map(db)

	m.Group("/api", func(r martini.Router) {
		r.Post("/transcode/audio", CreateTranscodeAudioTask)
		r.Post("/transcode/video", CreateTranscodeVideoTask)
		r.Post("/social/send", CreateSocialSendTask)
		r.Post("/event/send", CreaetEventSendTask)
		r.Post("/release/send", CreateReleaseSendTask)
	})


	m.Run()
}
