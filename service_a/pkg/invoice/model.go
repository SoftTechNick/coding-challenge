package invoice

type Invoice struct {
	Id           string  `json:"id"`
	CustomerName string  `json:"customer_name"`
	Amount       float64 `json:"amount"`
	DueDate      string  `json:"due_date"`
}
