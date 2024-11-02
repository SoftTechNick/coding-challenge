// repository/invoice.go
package invoice

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite Treiber für SQL-Paket
)

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(dbPath string) (*InvoiceRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	repo := &InvoiceRepository{db: db}
	return repo, nil
}

func (repo *InvoiceRepository) InitDB() error {
	query := `
    CREATE TABLE IF NOT EXISTS invoices (
        id TEXT PRIMARY KEY,
        customer_name TEXT,
        amount REAL,
        due_date TEXT
    );`
	_, err := repo.db.Exec(query)
	return err
}

func (repo *InvoiceRepository) InsertInvoice(invoice Invoice) error {
	query := `INSERT INTO invoices (id, customer_name, amount, due_date) VALUES (?, ?, ?, ?)`
	_, err := repo.db.Exec(query, invoice.Id, invoice.CustomerName, invoice.Amount, invoice.DueDate)
	return err
}

func (repo *InvoiceRepository) GetInvoiceById(id string) (*Invoice, error) {
	query := `SELECT id, customer_name, amount, due_date FROM invoices WHERE id = ?`
	row := repo.db.QueryRow(query, id)

	var invoice Invoice
	err := row.Scan(&invoice.Id, &invoice.CustomerName, &invoice.Amount, &invoice.DueDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &invoice, nil
}

// closes the database connection
func (repo *InvoiceRepository) Close() error {
	return repo.db.Close()
}
