package blockchain_addresses

import (
	"context"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) SetMeta(ctx context.Context, addressId string, meta interface{}) responses.BaseResponse {
	var res responses.ExtendedAddress

	_res := d.requester.Request(ctx, "addresses/set-meta", map[string]any{"id": addressId, "meta": meta}, &res)

	return _res
}
