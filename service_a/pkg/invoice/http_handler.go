// api/handler/handler.go
package invoice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type InvoiceHandler struct {
	srv *InvoiceService
}

func NewInvoiceHandler(srv *InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{srv: srv}
}

// GetInvoices handles the GET /invoices/{id} endpoint
func (ih *InvoiceHandler) GetInvoiceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	invoice, err := ih.srv.GetInvoiceById(id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}

	if invoice != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(invoice)
		return
	}

	http.Error(w, fmt.Sprintf("No invoice with id: %s found.", id), http.StatusNotFound)
}

// CreateInvoice handles the POST /invoices endpoint
func (ih *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := ih.srv.CreateInvoice(invoice)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}
