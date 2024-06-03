package blockchain_addresses

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) RecurrentAddresses(ctx context.Context) responses.BaseResponseGeneric[*[]responses.RecurrentAddress] {
	res := []responses.RecurrentAddress{}

	_res := d.requester.Request(ctx, "recurrent-addresses", map[string]string{}, &res)

	return responses.ConvertBase[*[]responses.RecurrentAddress](_res)
}
