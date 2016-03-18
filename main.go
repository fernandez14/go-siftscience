package gosift

import (
	"gopkg.in/jmcvetta/napping.v3"

	"fmt"
)

const (
	siftEndpoint = "https://api.siftscience.com/v203/events"
)

var ApiKey string
var UserId string
var SessionId string

func Track(event string, data map[string]interface{}) (error) {

	var siftErr SiftError{}

	payload := map[string]interface{}{}
	result := map[string]interface{}{}

	if ApiKey != "" {
		payload["$api_key"] = ApiKey
	}

	if UserId != "" {
		payload["$user_id"] = UserId
	}

	if SessionId != "" {
		payload["$session_id"] = SessionId
	}

	for key, value := range data {
		payload[key] = value
	}

	resp, err := napping.Post(siftEndpoint, &payload, &result, &siftErr)

	if resp.Status() == 200 && err == nil {
		return nil
	}

	return siftErr
}