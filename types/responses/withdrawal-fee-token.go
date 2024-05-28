package responses

import "time"

const (
	WithdrawalTokenBlockchainFeeSourceAdvancedBalance string = "advanced_balance"
	WithdrawalTokenBlockchainFeeSourceWallet          string = "wallet"
	WithdrawalTokenBlockchainFeeSourceNative          string = "native"
)

type WithdrawalFeeToken struct {
	BlockchainFeeCurrencyUSDRate string    `json:"blockchainFeeCurrencyUSDRate"`
	BlockchainFeeCurrency        string    `json:"blockchainFeeCurrency"`
	WithdrawalCurrencyUSDRate    string    `json:"withdrawalCurrencyUSDRate"`
	BlockchainFeeSource          string    `json:"blockchainFeeSource"`
	BlockchainFee                string    `json:"blockchainFee"`
	BlockchainFeeUSD             string    `json:"blockchainFeeUSD"`
	ServiceFee                   string    `json:"serviceFee"`
	ServiceFeeUSD                string    `json:"serviceFeeUSD"`
	WithdrawalMin                string    `json:"withdrawalMin"`
	Token                        string    `json:"token"`
	Until                        time.Time `json:"until"`
}
