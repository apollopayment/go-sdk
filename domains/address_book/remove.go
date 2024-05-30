package address_book

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) Remove(ctx context.Context, addressId string) responses.BaseResponse {
	var res string

	_res := d.requester.Request(ctx, "address-book/remove", map[string]string{"addressId": addressId}, &res)

	return _res
}
