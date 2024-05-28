package responses

import "time"

type AdvancedBalanceDepositAddress struct {
	Currency string    `json:"currency"`
	Network  string    `json:"network"`
	Address  string    `json:"address"`
	Tag      *string   `json:"tag"`
	Until    time.Time `json:"until"`
}
