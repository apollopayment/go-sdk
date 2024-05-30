package blockchain_addresses

import (
	"context"
	"github.com/apollopayment/go-sdk/types/requests"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) CreatePayoutAddress(ctx context.Context, payload requests.CreatePayOutAddress) responses.BaseResponseGeneric[*responses.PayOutAddress] {
	var res responses.PayOutAddress

	_res := d.requester.Request(ctx, "payout-address", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.PayOutAddress](_res)
}
