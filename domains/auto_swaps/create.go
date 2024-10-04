package auto_swaps

import (
	"context"
	"github.com/apollopayment/go-sdk/types/requests"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) Create(ctx context.Context, payload requests.CreateAutoSwap) responses.BaseResponseGeneric[*responses.AutoSwap] {
	var res responses.AutoSwap

	_res := d.requester.Request(ctx, "auto-swaps/create", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.AutoSwap](_res)
}
