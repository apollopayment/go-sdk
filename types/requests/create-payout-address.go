package requests

type CreatePayOutAddress struct {
	Currency string `json:"currency"`
	Network  string `json:"network"`
}

func (c CreatePayOutAddress) ToMap() map[string]any {
	payload := map[string]any{}

	if c.Currency != "" {
		payload["currency"] = c.Currency
	}
	if c.Network != "" {
		payload["network"] = c.Network
	}

	return payload
}
