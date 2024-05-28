package responses

import "time"

const (
	AddressTransactionStatusProcessed string = "processed"
	AddressTransactionStatusError     string = "error"
	AddressTransactionStatusRejected  string = "rejected"
	AddressTransactionStatusPending   string = "pending"
)
const (
	AddressTransactionTypeWithdrawal string = "withdrawal"
	AddressTransactionTypeDeposit    string = "deposit"
)

type AddressTransaction struct {
	Id              string    `json:"id"`
	Status          string    `json:"status"`
	Type            string    `json:"type"`
	Currency        string    `json:"currency"`
	Network         string    `json:"network"`
	AddressFrom     string    `json:"addressFrom"`
	AddressTo       string    `json:"addressTo"`
	Amount          string    `json:"amount"`
	Tx              string    `json:"tx"`
	FeeCurrency     *string   `json:"feeCurrency"`
	FeeAmount       *string   `json:"feeAmount"`
	FeeAmountUSD    *string   `json:"feeAmountUSD"`
	WithdrawalId    *string   `json:"withdrawalId"`
	OrphanDepositId *string   `json:"orphanDepositId"`
	CreatedAt       time.Time `json:"createdAt"`
}
