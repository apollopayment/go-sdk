package blockchain_addresses

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) BusinessAddresses(ctx context.Context, advancedBalanceId string) responses.BaseResponseGeneric[*[]responses.BusinessAddress] {
	res := []responses.BusinessAddress{}

	_res := d.requester.Request(ctx, "account-addresses", map[string]string{"advancedBalanceId": advancedBalanceId}, &res)

	return responses.ConvertBase[*[]responses.BusinessAddress](_res)
}
