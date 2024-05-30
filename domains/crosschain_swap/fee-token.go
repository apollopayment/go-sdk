package crosschain_swap

import (
	"context"
	"github.com/apollopayment/go-sdk/types/requests"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) FeeToken(ctx context.Context, payload requests.CrosschainSwapFeeToken) responses.BaseResponseGeneric[*responses.CrosschainSwapFeeToken] {
	var res responses.CrosschainSwapFeeToken

	_res := d.requester.Request(ctx, "swaps/fee-token", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.CrosschainSwapFeeToken](_res)
}
