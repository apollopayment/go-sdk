package blockchain_addresses

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) PayInAddresses(ctx context.Context) responses.BaseResponseGeneric[*[]responses.PayInAddress] {
	res := []responses.PayInAddress{}

	_res := d.requester.Request(ctx, "account-addresses", map[string]string{}, &res)

	return responses.ConvertBase[*[]responses.PayInAddress](_res)
}
