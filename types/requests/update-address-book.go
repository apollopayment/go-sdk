package requests

type UpdateAddressBook struct {
	AddressId string `json:"addressId"`
	Alias     string `json:"alias"`
	Comment   string `json:"comment"`
}

func (u UpdateAddressBook) ToMap() map[string]any {
	payload := map[string]any{}

	if u.AddressId != "" {
		payload["addressId"] = u.AddressId
	}
	if u.Alias != "" {
		payload["alias"] = u.Alias
	}
	if u.Comment != "" {
		payload["comment"] = u.Comment
	}

	return payload
}
