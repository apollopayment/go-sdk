package crosschain_swap

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) Get(ctx context.Context, crosschainSwapId string) responses.BaseResponseGeneric[*responses.CrosschainSwap] {
	var res responses.CrosschainSwap

	_res := d.requester.Request(ctx, "swaps/get", map[string]string{"id": crosschainSwapId}, &res)

	return responses.ConvertBase[*responses.CrosschainSwap](_res)
}
