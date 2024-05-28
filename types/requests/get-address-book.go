package requests

type GetAddressBook struct {
	Page     int      `json:"page"`     // Page number (starting from 0)
	Limit    int      `json:"limit"`    // Items length per page
	Networks []string `json:"networks"` // Filter by networks
}

func (g GetAddressBook) ToMap() map[string]any {
	payload := map[string]any{}

	if g.Page > 0 {
		payload["page"] = g.Page
	} else {
		payload["page"] = 0
	}

	if g.Limit > 0 {
		payload["limit"] = g.Limit
	}

	if g.Networks != nil && len(g.Networks) != 0 {
		payload["networks"] = g.Networks
	}

	return payload
}
