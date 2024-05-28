package responses

import "time"

const (
	InvoiceStatusCreated   string = "CREATED"
	InvoiceStatusInit      string = "INIT"
	InvoiceStatusPending   string = "PENDING"
	InvoiceStatusProcessed string = "PROCESSED"
	InvoiceStatusPartial   string = "PARTIAL"
	InvoiceStatusRejected  string = "REJECTED"
	InvoiceStatusError     string = "ERROR"
	InvoiceStatusExpired   string = "EXPIRED"
	InvoiceStatusOverpaid  string = "OVERPAID"
)

type Invoice struct {
	Id                              string    `json:"id"`
	AdvancedBalanceId               string    `json:"advancedBalanceId"`
	ExternalId                      string    `json:"externalId"`
	OrderId                         *string   `json:"orderId"`
	OrderLink                       *string   `json:"orderLink"`
	InvoiceLink                     string    `json:"invoiceLink"`
	Status                          string    `json:"status"`
	Order                           *string   `json:"order"`
	Description                     *string   `json:"description"`
	Currency                        string    `json:"currency"`
	Amount                          string    `json:"amount"`
	ReceivedNetwork                 *string   `json:"receivedNetwork"`
	ReceivedCurrency                *string   `json:"receivedCurrency"`
	ReceivedAmount                  string    `json:"receivedAmount"`
	ReceivedAmountInInvoiceCurrency string    `json:"receivedAmountInInvoiceCurrency"`
	Rate                            string    `json:"rate"`
	IncludeFee                      bool      `json:"includeFee"`
	AdditionalFees                  []string  `json:"additionalFees"`
	InsurancePercent                string    `json:"insurancePercent"`
	SlippagePercent                 string    `json:"slippagePercent"`
	WebhookURL                      string    `json:"webhookURL"`
	ReturnUrl                       string    `json:"returnUrl"`
	ExpiresAt                       time.Time `json:"expiresAt"`
	CreatedAt                       time.Time `json:"createdAt"`
	Currencies                      []struct {
		Currency string `json:"currency"`
		Networks []struct {
			Name   string `json:"name"`
			Amount string `json:"amount"`
		} `json:"networks"`
	} `json:"currencies"`
}
