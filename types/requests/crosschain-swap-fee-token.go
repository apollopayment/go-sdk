package requests

type CrosschainSwapFeeToken struct {
	AdvancedBalanceId string `json:"advancedBalanceId"`
	CurrencyFrom      string `json:"currencyFrom"`
	CurrencyTo        string `json:"currencyTo"`
	NetworkFrom       string `json:"networkFrom"`
	NetworkTo         string `json:"networkTo"`
	Amount            string `json:"amount"`
}

func (c CrosschainSwapFeeToken) ToMap() map[string]any {
	payload := map[string]any{}

	if c.AdvancedBalanceId != "" {
		payload["advancedBalanceId"] = c.AdvancedBalanceId
	}
	if c.CurrencyFrom != "" {
		payload["currencyFrom"] = c.CurrencyFrom
	}
	if c.CurrencyTo != "" {
		payload["currencyTo"] = c.CurrencyTo
	}
	if c.NetworkFrom != "" {
		payload["networkFrom"] = c.NetworkFrom
	}
	if c.NetworkTo != "" {
		payload["networkTo"] = c.NetworkTo
	}
	if c.Amount != "" {
		payload["amount"] = c.Amount
	}

	return payload
}
