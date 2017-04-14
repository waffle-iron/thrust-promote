package main
/*
    We want to be able to find an access token in the db
    but should we need to refresh? - Buffer delays the send
    and notifies user that a refresh is needed
    but if we have a refresh token we should be able to 
    make that request
*/

import (
    "testing"
    "encoding/json"
    "github.com/stretchr/testify/assert"
)

func TestSocialSend(t *testing.T) {
    /* 
    find the access token from the db
    and send the message
    - if the publishedAt date is in the future
    schedule the post
    */

    // AccessToken := "836439866992779265-D0t5sPjtuNgh0TdCebK1T37ofBX6rDG"
    // insert into DB


    payload := SocialSendPayload{
        Service: "twitter",
        Message: "hello twitter! #myfirsttweet",
        SocialID: 1,
    }

    metadata, err := json.Marshal(payload)
    task := NewTask("social_send", string(metadata))
    status, err := SocialSend(task)
    if assert.NoError(t, err) {
        assert.Equal(t, status, true, "Successful send")
    }

}