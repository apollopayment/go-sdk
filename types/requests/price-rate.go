package requests

type PriceRate struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (p PriceRate) ToMap() map[string]any {
	payload := map[string]any{}

	if p.From != "" {
		payload["from"] = p.From
	}
	if p.To != "" {
		payload["to"] = p.To
	}

	return payload
}
