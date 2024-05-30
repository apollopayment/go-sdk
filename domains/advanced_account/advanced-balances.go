package advanced_account

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) AdvancedBalances(ctx context.Context) responses.BaseResponseGeneric[*[]responses.AdvancedBalance] {
	res := []responses.AdvancedBalance{}

	_res := d.requester.Request(ctx, "advanced-balances", map[string]string{}, &res)

	return responses.ConvertBase[*[]responses.AdvancedBalance](_res)
}
