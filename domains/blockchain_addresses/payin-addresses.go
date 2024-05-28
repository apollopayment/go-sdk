package blockchain_addresses

import (
	"context"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) PayInAddresses(ctx context.Context, advancedBalanceId string) responses.BaseResponseGeneric[*[]responses.PayInAddress] {
	res := []responses.PayInAddress{}

	_res := d.requester.Request(ctx, "account-addresses", map[string]string{"advancedBalanceId": advancedBalanceId}, &res)

	return responses.ConvertBase[*[]responses.PayInAddress](_res)
}
