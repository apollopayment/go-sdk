package auto_swaps

import (
	"context"
	"apollopayment_sdk/types/requests"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) Create(ctx context.Context, payload requests.CreateAutoSwap) responses.BaseResponseGeneric[*responses.AutoSwap] {
	var res responses.AutoSwap

	_res := d.requester.Request(ctx, "auto_swaps/create", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.AutoSwap](_res)
}
