package responses

import "time"

const (
	TransactionStatusProcessed string = "processed"
	TransactionStatusError     string = "error"
	TransactionStatusRejected  string = "rejected"
	TransactionStatusPending   string = "pending"
)

const (
	OrderStatusInit      string = "init"
	OrderStatusError     string = "error"
	OrderStatusProcessed string = "processed"
	OrderStatusPending   string = "pending"
	OrderStatusExpired   string = "expired"
	OrderStatusPartial   string = "partial"
	OrderStatusOverpaid  string = "overpaid"
)

type OrderTransaction struct {
	Id        string `json:"id"`
	Status    string `json:"status"`
	Currency  string `json:"currency"`
	Network   string `json:"network"`
	Amount    string `json:"amount"`
	Tx        string `json:"tx"`
	Sender    string `json:"sender"`
	PriceUSD  string `json:"priceUSD"`
	AmountUSD string `json:"amountUSD"`
}

type Order struct {
	Id                string             `json:"id"`
	AdvancedBalanceId string             `json:"advancedBalanceId"`
	Currency          string             `json:"currency"`
	Network           string             `json:"network"`
	Status            string             `json:"status"`
	OrderId           string             `json:"orderId"`
	Description       *string            `json:"description"`
	Address           string             `json:"address"`
	Tag               *string            `json:"tag"`
	AddressId         string             `json:"addressId"`
	Amount            string             `json:"amount"`
	Received          string             `json:"received"`
	SuccessWebhook    *string            `json:"successWebhook"`
	ErrorWebhook      *string            `json:"errorWebhook"`
	ReturnUrl         *string            `json:"returnUrl"`
	Link              string             `json:"link"`
	ExpiresAt         time.Time          `json:"expiresAt"`
	CreatedAt         time.Time          `json:"createdAt"`
	UpdatedAt         time.Time          `json:"updatedAt"`
	Transactions      []OrderTransaction `json:"transactions"`
	OrphanDeposits    []struct {
	} `json:"orphanDeposits"`
}
