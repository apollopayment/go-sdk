package responses

import "time"

const (
	OrphanStatusPending   string = "PENDING"
	OrphanStatusProcessed string = "PROCESSED"
	OrphanStatusError     string = "ERROR"
	OrphanStatusRejected  string = "REJECTED"
)
const (
	OrphanStageDeposit    string = "DEPOSIT"
	OrphanStageWithdrawal string = "WITHDRAWAL"
)

type OrphanTransaction struct {
	Id             string  `json:"id"`
	OrganizationId string  `json:"organizationId"`
	OrderId        *string `json:"orderId"`
	Stage          string  `json:"stage"`
	Status         string  `json:"status"`
	Currency       string  `json:"currency"`
	Network        string  `json:"network"`
	Amount         string  `json:"amount"`
	CanWithdrawal  bool    `json:"canWithdrawal"`
	InTransaction  struct {
		AddressType string `json:"addressType"`
		AddressId   string `json:"addressId"`
		Address     string `json:"address"`
		TxId        string `json:"txId"`
		Amount      string `json:"amount"`
		Status      string `json:"status"`
	} `json:"inTransaction"`
	OutTransactions *struct {
		Address      string    `json:"address"`
		TxId         string    `json:"txId"`
		Amount       string    `json:"amount"`
		Status       string    `json:"status"`
		FeeAmount    string    `json:"feeAmount"`
		FeeAmountUSD string    `json:"feeAmountUSD"`
		WithdrawalId string    `json:"withdrawalId"`
		CreatedAt    time.Time `json:"createdAt"`
	} `json:"outTransactions"`
	CreatedAt time.Time `json:"createdAt"`
}
