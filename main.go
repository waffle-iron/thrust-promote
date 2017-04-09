package main

import (
	"github.com/go-martini/martini"
	"log"
	config "github.com/ammoses89/thrust-workers/config"
	db_ "github.com/ammoses89/thrust-workers/db"
)

const WORKER_COUNT = 5;


func main() {
	cfg := config.LoadConfig()
	taskMap := map[string]interface{}{
		"transcode_audio": TranscodeAudio,
		"transcode_video": TranscodeVideo,
		"social_send":     SocialSend,
		"event_send":      EventSend,
		"release_send":    ReleaseSend,
	}
	machine := NewMachine(cfg.Redis.Development)
	db := db_.NewPostgres(&cfg.Db.Development)
	log.Println("Registering Tasks...")
	machine.RegisterTasks(taskMap)
	log.Println("Launching Workers...")
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
		r.Post("/event/send", CreateEventSendTask)
		r.Post("/release/send", CreateReleaseSendTask)
	})


	m.Run()
}
