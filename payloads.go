package main

import (
    "time"
)

type TestPayload struct {
    id string `json:id`
    status string `json:status`
    Message string `json:message`
}

type AudioTranscodePayload struct {
    id string `json:id`
    status string `json:status`
    SourceUrl string `json:source_url`
    TargetUrl string `json:target_url`
    TranscodeType string `json:transcode_type`
    TrackID int `json:track_id`
}

type VideoTranscodePayload struct {
    id string `json:id`
    status string `json:status`
    SourceUrl string `json:source_url`
    TargetUrl string `json:target_url`
    ImageUrl string `json:image_url`
    TranscodeType string `json:transcode_type`
    TrackID int `json:track_id`
}

type SocialSendPayload struct {
    id string `json:id`
    status string `json:status`
    Service string `json:service`
    Message string `json:message`
    PublishAt time.Time `json:publish_at`
    SocialID int `json:social_id`
}