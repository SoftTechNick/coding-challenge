package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/SoftTechNick/coding-challenge/service_b/internal/messaging"
	"github.com/SoftTechNick/coding-challenge/service_b/pkg/invoice"
)

func main() {
	bus_url := os.Getenv("NATS_URL")
	messaging_service, err := messaging.NewMessagingService(bus_url)
	if err != nil {
		log.Printf("Error while creating messaging service: %s", err.Error())
	}
	defer messaging_service.Close()
	if err := messaging_service.Subscribe("invoice", func(message string) {
		var invoice invoice.Invoice
		if err := json.Unmarshal([]byte(message), &invoice); err != nil {
			log.Printf("Error while deserialization the invoice: %s", err.Error())
		}
		log.Printf("Received following invoice: %+v", invoice)

	}); err != nil {
		log.Printf("Error while subscribing invoice subject: %s", err.Error())
	}
	// wait to receive messages
	select {}
}
