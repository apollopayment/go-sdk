package auto_swaps

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) Get(ctx context.Context, autoSwapId string) responses.BaseResponseGeneric[*responses.AutoSwap] {
	var res responses.AutoSwap

	_res := d.requester.Request(ctx, "auto-swaps/get", map[string]string{"id": autoSwapId}, &res)

	return responses.ConvertBase[*responses.AutoSwap](_res)
}
