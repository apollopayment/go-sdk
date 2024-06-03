package apollopayment_sdk

import (
	"context"
	"github.com/apollopayment/go-sdk/domains/address_book"
	"github.com/apollopayment/go-sdk/domains/advanced_account"
	"github.com/apollopayment/go-sdk/domains/auto_swaps"
	"github.com/apollopayment/go-sdk/domains/base"
	"github.com/apollopayment/go-sdk/domains/blockchain_addresses"
	"github.com/apollopayment/go-sdk/domains/crosschain_bridge"
	"github.com/apollopayment/go-sdk/domains/crosschain_swap"
	"github.com/apollopayment/go-sdk/domains/invoices"
	"github.com/apollopayment/go-sdk/domains/orders"
	"github.com/apollopayment/go-sdk/domains/orphans"
	"github.com/apollopayment/go-sdk/domains/personal_addresses"
	"github.com/apollopayment/go-sdk/domains/webhooks"
	"github.com/apollopayment/go-sdk/domains/withdrawals"
	"github.com/apollopayment/go-sdk/noncer"
	"github.com/apollopayment/go-sdk/requester"
)

type Client struct {
	noncer *noncer.Noncer

	AddressBook         *address_book.Domain
	AdvancedBalance     *advanced_account.Domain
	AutoSwaps           *auto_swaps.Domain
	Base                *base.Domain
	BlockchainAddresses *blockchain_addresses.Domain
	Bridge              *crosschain_bridge.Domain
	Swaps               *crosschain_swap.Domain
	Invoices            *invoices.Domain
	Orders              *orders.Domain
	Orphans             *orphans.Domain
	PersonalAddresses   *personal_addresses.Domain
	Webhooks            *webhooks.Domain
	Withdrawals         *withdrawals.Domain
}

func New(public, private string) *Client {
	_requester := requester.New(public, private)

	client := &Client{
		noncer: noncer.New(),

		AddressBook:         address_book.New(_requester),
		AdvancedBalance:     advanced_account.New(_requester),
		AutoSwaps:           auto_swaps.New(_requester),
		Base:                base.New(_requester),
		BlockchainAddresses: blockchain_addresses.New(_requester),
		Bridge:              crosschain_bridge.New(_requester),
		Swaps:               crosschain_swap.New(_requester),
		Invoices:            invoices.New(_requester),
		Orders:              orders.New(_requester),
		Orphans:             orphans.New(_requester),
		PersonalAddresses:   personal_addresses.New(_requester),
		Webhooks:            webhooks.New(_requester),
		Withdrawals:         withdrawals.New(_requester),
	}

	advancedBalances := client.AdvancedBalance.AdvancedBalances(context.Background())

	if advancedBalances.Success {
		_advancedBalanced := *advancedBalances.Response

		for _, balance := range _advancedBalanced {
			_requester.SetAdvancedBalance(balance.AdvancedBalanceId)
		}
	}

	return client
}
