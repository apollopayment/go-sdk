package responses

type AvailableCurrency struct {
	Currency        string `json:"currency"`
	Alias           string `json:"alias"`
	AllowDeposit    bool   `json:"allowDeposit"`
	AllowWithdrawal bool   `json:"allowWithdrawal"`
	PriceUSD        string `json:"priceUSD"`
	Networks        []struct {
		Name                    string `json:"name"`
		Alias                   string `json:"alias"`
		Contract                string `json:"contract"`
		AddressRegex            string `json:"addressRegex"`
		TagRegex                string `json:"tagRegex"`
		AllowDeposit            bool   `json:"allowDeposit"`
		AllowWithdrawal         bool   `json:"allowWithdrawal"`
		AllowCrosschainBridge   bool   `json:"allowCrosschainBridge"`
		AllowCrosschainSwapFrom bool   `json:"allowCrosschainSwapFrom"`
		AllowCrosschainSwapTo   bool   `json:"allowCrosschainSwapTo"`
		AllowAutoSwapFrom       bool   `json:"allowAutoSwapFrom"`
		AllowAutoSwapTo         bool   `json:"allowAutoSwapTo"`
		WithdrawalFee           string `json:"withdrawalFee"`
		//WithdrawalMin           string `json:"withdrawalMin"`
		Confirmations    int  `json:"confirmations"`
		UnderMaintenance bool `json:"underMaintenance"`
	} `json:"networks"`
}
