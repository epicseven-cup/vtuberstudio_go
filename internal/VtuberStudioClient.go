package internal

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

type VtuberStudioClient struct {
	url      *url.URL
	status   chan bool
	receiver chan []byte
	outgoing chan []byte
}

func NewVtuberStudioClient(url url.URL) *VtuberStudioClient {

	log.Printf("Connecting to %v", url)
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Fatalf("dial error, exiting: %v", err)
	}

	status := make(chan bool, 1)
	receiver := make(chan []byte, 1)
	outgoing := make(chan []byte, 1)

	go func() {
		for {
			amount, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatalf("read error, exiting: %v", err)
			}
			log.Printf("Received message: %s", message)
			log.Printf("Message size: %d", amount)

			receiver <- message
		}

	}

	return &VtuberStudioClient{
		url:      &url,
		status:   status,
		receiver: receiver,
		outgoing: outgoing,
	}
}
