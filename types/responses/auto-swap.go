package responses

import "time"

const (
	AutoSwapStatusProcessed   string = "PROCESSED"
	AutoSwapStatusError       string = "ERROR"
	AutoSwapStatusRejected    string = "REJECTED"
	AutoSwapStatusPending     string = "PENDING"
	AutoSwapStatusWithdrawing string = "WITHDRAWING"
)

const (
	AutoSwapRiskLow    string = "Low"
	AutoSwapRiskMedium string = "Medium"
	AutoSwapRiskHigh   string = "High"
	AutoSwapRiskSevere string = "Severe"
)

type AutoSwap struct {
	Id                   string    `json:"id"`
	OrganizationId       string    `json:"organizationId"`
	Status               string    `json:"status"`
	Message              string    `json:"message"`
	AddressRiskLevel     string    `json:"addressRiskLevel"`
	CurrencyFrom         string    `json:"currencyFrom"`
	NetworkFrom          string    `json:"networkFrom"`
	CurrencyTo           string    `json:"currencyTo"`
	NetworkTo            string    `json:"networkTo"`
	AmountFrom           string    `json:"amountFrom"`
	AmountFromUSD        string    `json:"amountFromUSD"`
	AmountTo             string    `json:"amountTo"`
	AmountToUSD          string    `json:"amountToUSD"`
	AmountToReceive      string    `json:"amountToReceive"`
	Rate                 string    `json:"rate"`
	BlockchainFeeFrom    string    `json:"blockchainFeeFrom"`
	BlockchainFeeFromUSD string    `json:"blockchainFeeFromUSD"`
	BlockchainFeeTo      string    `json:"blockchainFeeTo"`
	BlockchainFeeToUSD   string    `json:"blockchainFeeToUSD"`
	ServiceFee           string    `json:"serviceFee"`
	WebhookUrl           string    `json:"webhookUrl"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
