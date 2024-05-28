package requests

type CreateBusinessAddress struct {
	AdvancedBalanceId string `json:"advancedBalanceId"`
	Currency          string `json:"currency"`
	Network           string `json:"network"`
	Alias             string `json:"alias"`
	Comment           string `json:"comment"`
}

func (c CreateBusinessAddress) ToMap() map[string]any {
	payload := map[string]any{}

	if c.AdvancedBalanceId != "" {
		payload["advancedBalanceId"] = c.AdvancedBalanceId
	}
	if c.Currency != "" {
		payload["currency"] = c.Currency
	}
	if c.Network != "" {
		payload["network"] = c.Network
	}
	if c.Alias != "" {
		payload["alias"] = c.Alias
	}
	if c.Comment != "" {
		payload["comment"] = c.Comment
	}

	return payload
}
