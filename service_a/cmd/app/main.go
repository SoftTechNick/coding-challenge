package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/SoftTechNick/coding-challenge/service_a/internal/messaging"
	"github.com/SoftTechNick/coding-challenge/service_a/pkg/invoice"
)

func main() {
	// getting env vars
	db_path := os.Getenv("SQLITE_DB_PATH")
	if db_path == "" {
		log.Fatalf("Missing env var 'SQLITE_DB_PATH'")
	}
	bus_url := os.Getenv("NATS_URL")
	if bus_url == "" {
		log.Fatalf("Missing env var 'NATS_URL'")
	}
	app_port := os.Getenv("APP_PORT")
	if app_port == "" {
		log.Fatalf("Missing env var 'APP_PORT'")
	}

	// setting up repo, services and handler
	invoice_repository, err := invoice.NewInvoiceRepository(db_path)
	if err != nil {
		log.Fatalf("Error while creating invoice repository: %s", err.Error())
	}
	invoice_repository.InitDB()
	defer invoice_repository.Close()
	messaging_service, err := messaging.NewMessagingService(bus_url)
	if err != nil {
		log.Fatalf("Error while creating messaging service: %s", err.Error())
	}
	defer messaging_service.Close()
	invoice_service := invoice.NewInvoiceService(invoice_repository, messaging_service)
	invoice_handler := invoice.NewInvoiceHandler(invoice_service)

	// set up the routes and assign handlers
	r := mux.NewRouter()
	r.HandleFunc("/invoices", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			invoice_handler.CreateInvoice(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	r.HandleFunc("/invoices/{id}", invoice_handler.GetInvoiceById).Methods("GET")

	// format port and start the server
	portf := fmt.Sprintf(":%s", app_port)
	log.Printf("Server is listening on port %s...", app_port)
	log.Fatal(http.ListenAndServe(portf, r))
}
