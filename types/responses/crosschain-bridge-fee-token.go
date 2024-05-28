package responses

import "time"

type CrosschainBridgeFeeToken struct {
	AdvancedBalanceId string    `json:"advancedBalanceId"`
	BlockchainFee     string    `json:"blockchainFee"`
	BlockchainFeeUSD  string    `json:"blockchainFeeUSD"`
	ServiceFeeUSD     string    `json:"serviceFeeUSD"`
	Amount            string    `json:"amount"`
	AmountUSD         string    `json:"amountUSD"`
	Token             string    `json:"token"`
	ExpiresAt         time.Time `json:"expiresAt"`
}
