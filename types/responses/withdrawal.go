package responses

import "time"

const (
	WithdrawalStatusInit      string = "init"
	WithdrawalStatusError     string = "error"
	WithdrawalStatusPending   string = "pending"
	WithdrawalStatusProcessed string = "processed"
	WithdrawalStatusRejected  string = "rejected"
)

type Withdrawal struct {
	Id                string    `json:"id"`
	AdvancedBalanceId string    `json:"advancedBalanceId"`
	AddressId         string    `json:"addressId"`
	Currency          string    `json:"currency"`
	Network           string    `json:"network"`
	Tx                string    `json:"tx"`
	Status            string    `json:"status"`
	Address           string    `json:"address"`
	Tag               *string   `json:"tag"`
	Amount            string    `json:"amount"`
	FeeCurrency       string    `json:"feeCurrency"`
	FeeSource         string    `json:"feeSource"`
	FeeAmount         string    `json:"feeAmount"`
	CreatedAt         time.Time `json:"createdAt"`
}

type AsyncWithdrawal struct {
	Id                string    `json:"id"`
	AdvancedBalanceId string    `json:"advancedBalanceId"`
	AddressId         string    `json:"addressId"`
	Currency          string    `json:"currency"`
	Network           string    `json:"network"`
	Tx                *string   `json:"tx"`
	Status            string    `json:"status"`
	Address           string    `json:"address"`
	Tag               *string   `json:"tag"`
	Amount            string    `json:"amount"`
	FeeCurrency       string    `json:"feeCurrency"`
	FeeSource         string    `json:"feeSource"`
	FeeAmount         string    `json:"feeAmount"`
	WebhookUrl        string    `json:"webhookUrl"`
	CreatedAt         time.Time `json:"createdAt"`
}
