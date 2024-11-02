package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SoftTechNick/coding-challenge/service_a/pkg/invoice"
	_ "github.com/mattn/go-sqlite3"
)

// Test the interaction between handler, service, repo and bus. Including the correct storage of the invoice in the in-memory DB.
func TestPostInvoice(t *testing.T) {
	// ARRANGE //
	invoice_repo, err := invoice.NewInvoiceRepository(":memory:")
	if err != nil {
		t.Fatalf("Failed to create in memory db: %v", err)
	}
	err = invoice_repo.InitDB()
	if err != nil {
		t.Fatalf("Failed to init memory db: %v", err)
	}
	defer invoice_repo.Close()

	messaging_service := NewMockMessagingService()
	invoice_service := invoice.NewInvoiceService(invoice_repo, messaging_service)
	handler := invoice.NewInvoiceHandler(invoice_service)

	invoice := invoice.Invoice{
		Id:           "1",
		CustomerName: "John Doe",
		Amount:       150.00,
		DueDate:      "01.02.2031",
	}
	body, _ := json.Marshal(invoice)

	// create a post req
	req, err := http.NewRequest("POST", "/invoices", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// ACT //
	handler.CreateInvoice(rr, req)
	id := strings.TrimSpace(rr.Body.String())
	id = strings.Trim(id, "\"")

	// ASSERT //
	// validate status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %v, got %v", http.StatusCreated, status)
	}

	// validate the stored invoice obj
	invoice_from_db, err := invoice_repo.GetInvoiceById(id)
	if err != nil {
		t.Errorf("Failed to retrieve stored invoice from database: %v", err)
	}
	if invoice_from_db == nil {
		t.Errorf("Missing invoice with id: %s", id)
	}

	if invoice_from_db.Id != id {
		t.Errorf("Failed to store invoice in database: %v", err)
	}

	// vaidate publishing the invoice in the messaging service
	if len(messaging_service.PublishedMessages) != 1 {
		t.Errorf("Expected 1 message to be published, got %v", len(messaging_service.PublishedMessages))
	}
}

type MockMessagingService struct {
	PublishedMessages map[string][]interface{}
}

func NewMockMessagingService() *MockMessagingService {
	return &MockMessagingService{PublishedMessages: make(map[string][]interface{})}
}

func (m *MockMessagingService) Publish(subject string, message interface{}) error {
	m.PublishedMessages[subject] = append(m.PublishedMessages[subject], message)
	return nil
}

func (m *MockMessagingService) Subscribe(subject string, handler func(message string)) error {
	// Mock Subscribe-Funktionalität
	return nil
}

func (m *MockMessagingService) Close() {
	// Mock Close-Funktionalität
}
