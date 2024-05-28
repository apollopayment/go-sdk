package requests

type GetUserAddress struct {
	Id           string `json:"id"`
	Currency     string `json:"currency"`
	Network      string `json:"network"`
	RenewAddress bool   `json:"renewAddress"`
}

func (g GetUserAddress) ToMap() map[string]any {
	payload := map[string]any{}

	if g.Id != "" {
		payload["id"] = g.Id
	}
	if g.Currency != "" {
		payload["currency"] = g.Currency
	}
	if g.Network != "" {
		payload["network"] = g.Network
	}

	payload["renewAddress"] = g.RenewAddress

	return payload
}
