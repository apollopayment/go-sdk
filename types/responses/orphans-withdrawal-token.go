package responses

import "time"

type OrphansWithdrawalToken struct {
	Currency         string    `json:"currency"`
	Network          string    `json:"network"`
	FeeSource        string    `json:"feeSource"`
	BlockchainFee    string    `json:"blockchainFee"`
	BlockchainFeeUSD string    `json:"blockchainFeeUSD"`
	ServiceFee       string    `json:"serviceFee"`
	ServiceFeeUSD    string    `json:"serviceFeeUSD"`
	Amount           string    `json:"amount"`
	AmountTo         string    `json:"amountTo"`
	Price            string    `json:"price"`
	Token            string    `json:"token"`
	Expires          time.Time `json:"expires"`
}
