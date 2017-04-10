package main

import (
    "fmt"
    // "os/exec"
)

func CreateTranscodeVideoTask() string {
    // TODO add task to worker
    fmt.Println("Save Task")
    return "{\"status\": 200}"
}

func TranscodeVideo(task *Task) {
    // cmd := `ffmpeg -y -loop 1 -f image2 -i #{@image_file} \
    //      -i "#{@audio_file}" -c:v libx264 -c:a aac -strict experimental \
    //      -b:a 192k -t #{audio_length} #{@video_file}`
}