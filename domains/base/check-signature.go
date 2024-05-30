package base

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) CheckSignature(ctx context.Context) responses.BaseResponseGeneric[*responses.CheckSignResponse] {
	var res responses.CheckSignResponse

	_res := d.requester.Request(ctx, "test-signature", map[string]string{}, &res)

	return responses.ConvertBase[*responses.CheckSignResponse](_res)
}
