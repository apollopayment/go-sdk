package requests

type ValidateAddress struct {
	Address string `json:"address"`
	Network string `json:"network"`
}

func (v ValidateAddress) ToMap() map[string]any {
	payload := map[string]any{}

	if v.Address != "" {
		payload["address"] = v.Address
	}
	if v.Network != "" {
		payload["network"] = v.Network
	}

	return payload
}
