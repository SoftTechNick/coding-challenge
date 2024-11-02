// repository/invoice.go
package invoice

import (
	"log"

	"github.com/SoftTechNick/coding-challenge/service_a/internal/messaging"
	"github.com/google/uuid"
)

// InvoiceRepository verwaltet die Datenbankoperationen für die Invoice-Entität
type InvoiceService struct {
	repo              *InvoiceRepository
	messaging_service messaging.MessagingService
}

func NewInvoiceService(repo *InvoiceRepository, messagingService messaging.MessagingService) *InvoiceService {
	return &InvoiceService{repo: repo, messaging_service: messagingService}
}

func (srv *InvoiceService) GetInvoiceById(id string) (*Invoice, error) {
	invoice, err := srv.repo.GetInvoiceById(id)
	return invoice, err
}

// Creates a new ID and stores the invoice, returning the created ID if the operation was successful.
func (srv *InvoiceService) CreateInvoice(invoice Invoice) (*string, error) {
	id := uuid.New()
	invoice.Id = id.String()
	err := srv.repo.InsertInvoice(invoice)

	if err != nil {
		return nil, err
	}

	// at this point, the error should only be logged, as there is currently no rollback for the db and it can therefore not be handled correctly, except for logging
	if err := srv.messaging_service.Publish("invoice", invoice); err != nil {
		log.Printf("An unexpected error occured while sending invoice message: {%s}", err)
	}

	return &invoice.Id, nil
}
