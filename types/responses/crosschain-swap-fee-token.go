package responses

import "time"

type CrosschainSwapFeeToken struct {
	AdvancedBalanceId           string    `json:"advancedBalanceId"`
	AmountFrom                  string    `json:"amountFrom"`
	AmountTo                    string    `json:"amountTo"`
	ServiceBlockchainFeeSource  string    `json:"serviceBlockchainFeeSource"`
	ServiceBlockchainFee        string    `json:"serviceBlockchainFee"`
	ServiceBlockchainFeeUSD     string    `json:"serviceBlockchainFeeUSD"`
	ProviderBlockchainFeeSource string    `json:"providerBlockchainFeeSource"`
	ProviderBlockchainFee       string    `json:"providerBlockchainFee"`
	ProviderBlockchainFeeUSD    string    `json:"providerBlockchainFeeUSD"`
	ServiceFeeSource            string    `json:"serviceFeeSource"`
	ServiceFee                  string    `json:"serviceFee"`
	ServiceFeeUSD               string    `json:"serviceFeeUSD"`
	Price                       string    `json:"price"`
	Token                       string    `json:"token"`
	ExpiresAt                   time.Time `json:"expiresAt"`
}
