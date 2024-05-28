package responses

import "time"

const (
	CrosschainSwapStatusCreated   string = "CREATED"
	CrosschainSwapStatusPending   string = "PENDING"
	CrosschainSwapStatusError     string = "ERROR"
	CrosschainSwapStatusRejected  string = "REJECTED"
	CrosschainSwapStatusProcessed string = "PROCESSED"
)

const (
	CrosschainSwapBlockchainFeeSourceAddress string = "ADDRESS"
	CrosschainSwapBlockchainFeeSourceAdvance string = "ADVANCE"
)

type CrosschainSwap struct {
	Id                          string    `json:"id"`
	ClientId                    string    `json:"clientId"`
	AdvancedBalanceId           string    `json:"advancedBalanceId"`
	AddressFromId               string    `json:"addressFromId"`
	AddressToId                 string    `json:"addressToId"`
	CurrencyFrom                string    `json:"currencyFrom"`
	CurrencyTo                  string    `json:"currencyTo"`
	NetworkFrom                 string    `json:"networkFrom"`
	NetworkTo                   string    `json:"networkTo"`
	Status                      string    `json:"status"`
	RejectMessage               *string   `json:"rejectMessage"`
	AmountFrom                  string    `json:"amountFrom"`
	AmountTo                    string    `json:"amountTo"`
	Price                       string    `json:"price"`
	ServiceBlockchainFeeSource  string    `json:"serviceBlockchainFeeSource"`
	ServiceBlockchainFee        string    `json:"serviceBlockchainFee"`
	ServiceBlockchainFeeUSD     string    `json:"serviceBlockchainFeeUSD"`
	ProviderBlockchainFeeSource string    `json:"providerBlockchainFeeSource"`
	ProviderBlockchainFee       string    `json:"providerBlockchainFee"`
	ProviderBlockchainFeeUSD    string    `json:"providerBlockchainFeeUSD"`
	ServiceFeeSource            string    `json:"serviceFeeSource"`
	ServiceFee                  string    `json:"serviceFee"`
	ServiceFeeUSD               string    `json:"serviceFeeUSD"`
	WebhookUrl                  string    `json:"webhookUrl"`
	CreatedAt                   time.Time `json:"createdAt"`
}
