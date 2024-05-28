package orphans

import (
	"context"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) GetDeposit(ctx context.Context, orphanId string) responses.BaseResponseGeneric[*responses.OrphanTransaction] {
	var res responses.OrphanTransaction

	_res := d.requester.Request(ctx, "orphan-deposits/get-deposit", map[string]string{"id": orphanId}, &res)

	return responses.ConvertBase[*responses.OrphanTransaction](_res)
}
