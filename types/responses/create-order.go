package responses

import "time"

type CreateOrder struct {
	Status            string    `json:"status"`
	Link              string    `json:"link"`
	Amount            string    `json:"amount"`
	AdvancedBalanceId string    `json:"advancedBalanceId"`
	Currency          string    `json:"currency"`
	Network           string    `json:"network"`
	Address           string    `json:"address"`
	Tag               *string   `json:"tag"`
	AddressId         string    `json:"addressId"`
	OrderId           string    `json:"orderId"`
	ClientOrderId     string    `json:"clientOrderId"`
	Description       string    `json:"description"`
	SuccessWebhook    *string   `json:"successWebhook"`
	ErrorWebhook      *string   `json:"errorWebhook"`
	ReturnUrl         *string   `json:"returnUrl"`
	ExpiresAt         time.Time `json:"expiresAt"`
	CreatedAt         time.Time `json:"createdAt"`
}
