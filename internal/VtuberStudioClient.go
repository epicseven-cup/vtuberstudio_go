package internal

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
)

type VtuberStudioClient struct {
	url      *url.URL
	status   chan bool
	receiver chan []byte
	outgoing chan interface{}
	handler  map[string]Handler
}

func NewVtuberStudioClient(url url.URL) *VtuberStudioClient {

	log.Printf("Connecting to %v", url)
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Fatalf("dial error, exiting: %v", err)
	}

	status := make(chan bool, 1)
	receiver := make(chan []byte, 1)
	outgoing := make(chan interface{}, 1)

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

	}()

	go func() {
		for {
			select {
			case <-status:
				os.Exit(0)
			case message := <-outgoing:
				log.Printf("Sending message: %s", message)
				err := conn.WriteJSON(message)
				if err != nil {
					return
				}
			}
		}
	}()

	return &VtuberStudioClient{
		url:      &url,
		status:   status,
		receiver: receiver,
		outgoing: outgoing,
	}
}

func (v *VtuberStudioClient) Close() {
	// Bad close, I need to sleep
	v.status <- true
}

func (v *VtuberStudioClient) AddHandler(messageType string, handler Handler) error {
	_, ok := v.handler[messageType]
	if ok {
		return fmt.Errorf("handler for message type %s already exists", messageType)
	}
	v.handler[messageType] = handler
	return nil
}

func (v *VtuberStudioClient) SendMessage(m interface{}) {
	v.outgoing <- m
}

func (v *VtuberStudioClient) AuthenticatePlugin(name string, developerName string, iconBase64 string) {

	m := AuthenticationRequest{
		RequestMessageBase: RequestMessageBase{
			ApiName:     API_NAME,
			ApiVersion:  API_VERSION,
			RequestID:   uuid.New().String(),
			MessageType: "AuthenticationTokenRequest",
		},
		AuthenticationRequestData: AuthenticationRequestData{
			PluginName:      name,
			PluginDeveloper: developerName,
			PluginIcon:      iconBase64,
		},
	}
	v.SendMessage(m)
}
