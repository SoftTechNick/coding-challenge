// messaging_service.go
package messaging

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type MessagingService interface {
	Publish(subject string, message interface{}) error
	Subscribe(subject string, handler func(message string)) error
	Close()
}

type MessagingServiceImpl struct {
	conn *nats.Conn
}

func NewMessagingService(url string) (*MessagingServiceImpl, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &MessagingServiceImpl{conn: conn}, nil
}

// publishes a message under the specified subject
func (ms *MessagingServiceImpl) Publish(subject string, message interface{}) error {
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
func (ms *MessagingServiceImpl) Subscribe(subject string, handler func(message string)) error {
	_, err := ms.conn.Subscribe(subject, func(msg *nats.Msg) {
		handler(string(msg.Data))
	})
	return err
}

// closes the messaging service connection
func (ms *MessagingServiceImpl) Close() {
	ms.conn.Close()
}
