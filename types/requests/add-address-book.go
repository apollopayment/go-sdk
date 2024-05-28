package requests

type AddAddressBook struct {
	Address  string   `json:"address"`
	Networks []string `json:"networks"`
	Alias    string   `json:"alias"`
	Comment  string   `json:"comment"`
}

func (a AddAddressBook) ToMap() map[string]any {
	payload := map[string]any{}

	if a.Address != "" {
		payload["address"] = a.Address
	}
	if a.Alias != "" {
		payload["alias"] = a.Alias
	}
	if a.Comment != "" {
		payload["comment"] = a.Comment
	}

	if a.Networks != nil && len(a.Networks) > 0 {
		payload["networks"] = a.Networks
	}

	return payload
}
