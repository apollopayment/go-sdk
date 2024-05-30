package blockchain_addresses

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) PayOutAddresses(ctx context.Context, advancedBalanceId string) responses.BaseResponseGeneric[*[]responses.PayOutAddress] {
	res := []responses.PayOutAddress{}

	_res := d.requester.Request(ctx, "payout-addresses", map[string]string{"advancedBalanceId": advancedBalanceId}, &res)

	return responses.ConvertBase[*[]responses.PayOutAddress](_res)
}
