package invoices

import (
	"context"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) GetInvoice(ctx context.Context, invoiceId string) responses.BaseResponseGeneric[*responses.Invoice] {
	var res responses.Invoice

	_res := d.requester.Request(ctx, "get-invoice", map[string]string{"invoiceId": invoiceId}, &res)

	return responses.ConvertBase[*responses.Invoice](_res)
}
