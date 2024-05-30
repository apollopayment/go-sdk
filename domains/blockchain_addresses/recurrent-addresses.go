package blockchain_addresses

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) RecurrentAddresses(ctx context.Context, advancedBalanceId string) responses.BaseResponseGeneric[*[]responses.RecurrentAddress] {
	res := []responses.RecurrentAddress{}

	_res := d.requester.Request(ctx, "recurrent-addresses", map[string]string{"advancedBalanceId": advancedBalanceId}, &res)

	return responses.ConvertBase[*[]responses.RecurrentAddress](_res)
}
