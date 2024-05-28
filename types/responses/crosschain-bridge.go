package responses

import "time"

const (
	CrosschainBridgeStatusCreated   string = "CREATED"
	CrosschainBridgeStatusPending   string = "PENDING"
	CrosschainBridgeStatusError     string = "ERROR"
	CrosschainBridgeStatusRejected  string = "REJECTED"
	CrosschainBridgeStatusProcessed string = "PROCESSED"
)

type CrosschainBridge struct {
	Id                string    `json:"id"`
	ClientId          string    `json:"clientId"`
	AdvancedBalanceId string    `json:"advancedBalanceId"`
	Currency          string    `json:"currency"`
	NetworkFrom       string    `json:"networkFrom"`
	NetworkTo         string    `json:"networkTo"`
	Status            string    `json:"status"`
	RejectMessage     *string   `json:"rejectMessage"`
	Amount            string    `json:"amount"`
	AmountUSD         string    `json:"amountUSD"`
	BlockchainFee     string    `json:"blockchainFee"`
	BlockchainFeeUSD  string    `json:"blockchainFeeUSD"`
	ServiceFeeUSD     string    `json:"serviceFeeUSD"`
	WebhookUrl        string    `json:"webhookUrl"`
	CreatedAt         time.Time `json:"createdAt"`
}
