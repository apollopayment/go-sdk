package blockchain_addresses

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) FindByAddress(ctx context.Context, address string) responses.BaseResponseGeneric[*responses.ExtendedAddress] {
	var res responses.ExtendedAddress

	_res := d.requester.Request(ctx, "addresses/find-by-address", map[string]string{"address": address}, &res)

	return responses.ConvertBase[*responses.ExtendedAddress](_res)
}
