package blockchain_addresses

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) FindByID(ctx context.Context, addressId string) responses.BaseResponseGeneric[*responses.ExtendedAddress] {
	var res responses.ExtendedAddress

	_res := d.requester.Request(ctx, "addresses/find-by-id", map[string]string{"id": addressId}, &res)

	return responses.ConvertBase[*responses.ExtendedAddress](_res)
}
