package responses

const (
	TariffTypePercent string = "PERCENT"
	TariffTypeFixed   string = "FIXED"
)

const (
	TariffActionTransactionInit         string = "TRANSACTION_INIT"
	TariffActionTransactionProcessed    string = "TRANSACTION_PROCESSED"
	TariffActionTransfer                string = "TRANSFER"
	TariffActionOrderInit               string = "ORDER_INIT"
	TariffActionOrderDeposit            string = "ORDER_DEPOSIT"
	TariffActionInternalTransfer        string = "INTERNAL_TRANSFER"
	TariffActionWalletWithdrawal        string = "WALLET_WITHDRAWAL"
	TariffActionRecurrentWithdrawal     string = "RECURRENT_WITHDRAWAL"
	TariffActionPersonalWithdrawal      string = "PERSONAL_WITHDRAWAL"
	TariffActionPayoutWithdrawal        string = "PAYOUT_WITHDRAWAL"
	TariffActionRecurrentDeposit        string = "RECURRENT_DEPOSIT"
	TariffActionPersonalDeposit         string = "PERSONAL_DEPOSIT"
	TariffActionPayoutDeposit           string = "PAYOUT_DEPOSIT"
	TariffActionExchangeInternal        string = "EXCHANGE_INTERNAL"
	TariffActionExchangeAuto            string = "EXCHANGE_AUTO"
	TariffActionBridgeInternal          string = "BRIDGE_INTERNAL"
	TariffActionBridgeExternal          string = "BRIDGE_EXTERNAL"
	TariffActionKytTransaction          string = "KYT_TRANSACTION"
	TariffActionKytWithdrawalAddress    string = "KYT_WITHDRAWAL_ADDRESS"
	TariffActionKytAddress              string = "KYT_ADDRESS"
	TariffActionOrphanDepositWithdrawal string = "ORPHAN_DEPOSIT_WITHDRAWAL"
	TariffActionSepaWithdrawal          string = "SEPA_WITHDRAWAL"
	TariffActionSepaDeposit             string = "SEPA_DEPOSIT"
	TariffActionCryptoExchange          string = "FIAT_CRYPTO_EXCHANGE"
	TariffActionPayoutAutoSwap          string = "PAYOUT_AUTO_SWAP"
)

type AdvancedBalance struct {
	AdvancedBalanceId             string   `json:"advancedBalanceId"`
	Currency                      string   `json:"currency"`
	Blocked                       bool     `json:"blocked"`
	Balance                       string   `json:"balance"`
	AvailableCurrenciesForDeposit []string `json:"availableCurrenciesForDeposit"`
	Tariffs                       []struct {
		Id                   string `json:"id"`
		Action               string `json:"action"`
		Amount               string `json:"amount"`
		Type                 string `json:"type"`
		MinAmount            string `json:"minAmount"`
		MaxAmount            string `json:"maxAmount"`
		InvoiceAdditionalFee bool   `json:"invoiceAdditionalFee"`
	} `json:"tariffs"`
}
