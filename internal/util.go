package internal

import (
	"encoding/json"
	"log"
)

//https://stackoverflow.com/questions/47911187/golang-elegantly-json-decode-different-structures

/*
Two different idea here, that I can have a metadata wrapper structure and have json raw value to parse it correctlu
Or
I can just have a bunch of if statements and parse it
*/

func ParseByMessageType(message []byte) ([]interface{}, error) {
	response := RespondMessageBase{}
	err := json.Unmarshal(message, &response)

	if err != nil {
		log.Fatalf("json unmarshal error: %v", err)
	}

	switch response.MessageType {
		case:

		case API_VERSION:

	}


	return nil, nil
}
