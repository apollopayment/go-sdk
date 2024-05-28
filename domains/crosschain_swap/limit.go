package crosschain_swap

import (
	"context"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) Limit(ctx context.Context) responses.BaseResponseGeneric[*responses.CrosschainLimit] {
	var res responses.CrosschainLimit

	_res := d.requester.Request(ctx, "swaps/limit", map[string]string{}, &res)

	return responses.ConvertBase[*responses.CrosschainLimit](_res)
}
