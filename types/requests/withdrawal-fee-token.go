package requests

type WithdrawalFeeToken struct {
	AdvancedBalanceId string `json:"advancedBalanceId"`
	AddressId         string `json:"addressId"`
	Amount            string `json:"amount"`
	Native            bool   `json:"native"`
}

func (w WithdrawalFeeToken) ToMap() map[string]any {
	payload := map[string]any{}

	if w.AdvancedBalanceId != "" {
		payload["advancedBalanceId"] = w.AdvancedBalanceId
	}
	if w.AddressId != "" {
		payload["addressId"] = w.AddressId
	}
	if w.Amount != "" {
		payload["amount"] = w.Amount
	}

	payload["native"] = w.Native

	return payload
}
