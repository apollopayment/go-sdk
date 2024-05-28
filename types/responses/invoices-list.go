package responses

type InvoicesList struct {
	Invoices []Invoice `json:"invoices"`
	Total    int       `json:"total"`
}
