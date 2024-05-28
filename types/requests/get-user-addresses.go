package requests

type GetUserAddressBalance struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type GetUserAddresses struct {
	Id       string                 `json:"id"`
	IsActive bool                   `json:"isActive"`
	Currency []string               `json:"currency"`
	Network  []string               `json:"network"`
	Balance  *GetUserAddressBalance `json:"balance"`
	Offset   int                    `json:"offset"`
	Limit    int                    `json:"limit"`
}

func (g GetUserAddresses) ToMap() map[string]any {
	payload := map[string]any{}

	if g.Id != "" {
		payload["id"] = g.Id
	}

	if g.Currency != nil && len(g.Currency) != 0 {
		payload["currency"] = g.Currency
	}
	if g.Network != nil && len(g.Network) != 0 {
		payload["network"] = g.Network
	}

	if g.Balance != nil {
		balancePayload := map[string]any{}

		if g.Balance.From != "" {
			balancePayload["from"] = g.Balance.From
		}
		if g.Balance.To != "" {
			balancePayload["to"] = g.Balance.To
		}

		payload["balance"] = balancePayload
	}

	if g.Limit > 0 {
		payload["limit"] = g.Limit
	} else {
		payload["limit"] = 100
	}

	if g.Offset >= 0 {
		payload["offset"] = g.Offset
	}

	payload["isActive"] = g.IsActive

	return payload
}
