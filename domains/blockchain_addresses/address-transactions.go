package blockchain_addresses

import (
	"context"
	"github.com/apollopayment/go-sdk/types/requests"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) AddressTransactions(ctx context.Context, payload *requests.AddressTransactions) responses.BaseResponseGeneric[*[]responses.AddressTransaction] {
	res := []responses.AddressTransaction{}

	_res := d.requester.Request(ctx, "addresses/address-transactions", payload.ToMap(), &res)

	return responses.ConvertBase[*[]responses.AddressTransaction](_res)
}
