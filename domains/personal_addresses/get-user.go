package personal_addresses

import (
	"context"
	"apollopayment_sdk/types/requests"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) GetUser(ctx context.Context, payload requests.CreatePersonalUser) responses.BaseResponseGeneric[*responses.PersonalUserWithAddresses] {
	var res responses.PersonalUserWithAddresses

	_res := d.requester.Request(ctx, "personal_addresses/create-user", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.PersonalUserWithAddresses](_res)
}
