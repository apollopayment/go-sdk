package requests

type CreateWithdrawal struct {
	//AdvancedBalanceId string  `json:"advancedBalanceId"`
	AddressId string  `json:"addressId"`
	Address   string  `json:"address"`
	Tag       *string `json:"tag"`
	Amount    string  `json:"amount"`
	FeeToken  string  `json:"feeToken"`
}

type CreateAsyncWithdrawal struct {
	//AdvancedBalanceId string  `json:"advancedBalanceId"`
	AddressId  string  `json:"addressId"`
	Address    string  `json:"address"`
	Tag        *string `json:"tag"`
	Amount     string  `json:"amount"`
	FeeToken   string  `json:"feeToken"`
	WebhookUrl string  `json:"webhookUrl"`
}

func (c CreateWithdrawal) ToMap() map[string]any {
	payload := map[string]any{}

	//if c.AdvancedBalanceId != "" {
	//	payload["advancedBalanceId"] = c.AdvancedBalanceId
	//}
	if c.AddressId != "" {
		payload["addressId"] = c.AddressId
	}
	if c.Address != "" {
		payload["address"] = c.Address
	}
	if c.Tag != nil && *c.Tag != "" {
		payload["tag"] = c.Tag
	}
	if c.Amount != "" {
		payload["amount"] = c.Amount
	}
	if c.FeeToken != "" {
		payload["feeToken"] = c.FeeToken
	}

	return payload
}
func (c CreateAsyncWithdrawal) ToMap() map[string]any {
	payload := map[string]any{}

	//if c.AdvancedBalanceId != "" {
	//	payload["advancedBalanceId"] = c.AdvancedBalanceId
	//}
	if c.AddressId != "" {
		payload["addressId"] = c.AddressId
	}
	if c.Address != "" {
		payload["address"] = c.Address
	}
	if c.Tag != nil && *c.Tag != "" {
		payload["tag"] = c.Tag
	}
	if c.Amount != "" {
		payload["amount"] = c.Amount
	}
	if c.FeeToken != "" {
		payload["feeToken"] = c.FeeToken
	}
	if c.WebhookUrl != "" {
		payload["webhookUrl"] = c.WebhookUrl
	}

	return payload
}
