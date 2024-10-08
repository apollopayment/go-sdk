package personal_addresses

import (
	"context"
	"github.com/apollopayment/go-sdk/types/requests"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) CreateUser(ctx context.Context, payload requests.CreatePersonalUser) responses.BaseResponseGeneric[*responses.PersonalUserWithAddresses] {
	var res responses.PersonalUserWithAddresses

	_res := d.requester.Request(ctx, "personal-addresses/create-user", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.PersonalUserWithAddresses](_res)
}
