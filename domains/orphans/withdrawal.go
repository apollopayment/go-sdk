package orphans

import (
	"context"
	"apollopayment_sdk/types/requests"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) Withdrawal(ctx context.Context, payload requests.OrphansWithdrawal) responses.BaseResponseGeneric[*responses.OrphanTransaction] {
	var res responses.OrphanTransaction

	_res := d.requester.Request(ctx, "orphan-deposits/withdrawal", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.OrphanTransaction](_res)
}
