package requests

type AdvancedBalanceDepositAddress struct {
	AdvancedBalanceId string `json:"advancedBalanceId"`
	Currency          string `json:"currency"`
	Network           string `json:"network"`
}

func (a AdvancedBalanceDepositAddress) ToMap() map[string]any {
	payload := map[string]any{}

	if a.AdvancedBalanceId != "" {
		payload["advancedBalanceId"] = a.AdvancedBalanceId
	}
	if a.Currency != "" {
		payload["currency"] = a.Currency
	}
	if a.Network != "" {
		payload["network"] = a.Network
	}

	return payload
}
