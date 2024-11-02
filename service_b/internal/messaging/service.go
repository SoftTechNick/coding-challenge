// messaging_service.go
package messaging

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type MessagingService struct {
	conn *nats.Conn
}

func NewMessagingService(url string) (*MessagingService, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &MessagingService{conn: conn}, nil
}

// publishes a message under the specified subject
func (ms *MessagingService) Publish(subject string, message interface{}) error {
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = ms.conn.Publish(subject, msgBytes)
	if err != nil {
		return err
	}

	return ms.conn.Flush()
}

// subscribes a message under the specified subject
func (ms *MessagingService) Subscribe(subject string, handler func(message string)) error {
	_, err := ms.conn.Subscribe(subject, func(msg *nats.Msg) {
		handler(string(msg.Data))
	})
	return err
}

// closes the messaging service connection
func (ms *MessagingService) Close() {
	ms.conn.Close()
}
