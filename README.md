# ApolloPayment SDK

ApolloPayment SDK is a comprehensive JavaScript library designed to streamline the integration of blockchain-based payment solutions into various applications. This SDK provides developers with the tools necessary to facilitate secure, transparent, and efficient transactions on supported blockchain networks.

Key Features:

- **Ease of Integration**: Simplifies the process of adding blockchain payment capabilities to your applications.
- **Security**: Ensures high-level security for all transactions using blockchain technology.
- **Transparency**: Leverages the transparency of blockchain networks to provide clear and verifiable transaction records.
- **Multi-Network Support**: Supports multiple blockchain networks, providing flexibility and scalability.

For detailed documentation, installation guides, and API references, please visit ApolloPayment Documentation.

This package makes it easy [ApolloPayment Api](https://docs.apollopayment.io/).

## Installation

`go get github.com/apollopayment/go-sdk`

## Use

Go to your personal account
[https://app.apollopayment.io/api-keys](https://app.apollopayment.io/api-keys)
and get api-keys.

*Substitute keys in class call:*

```go
package main

import (
	"context"
	apollopayment_sdk "github.com/apollopayment/go-sdk"
)

func main() {
	client := apollopayment_sdk.New("__PUBLIC_KEY__", "__PRIVATE_KEY__")
}
```

### Check signature

You can test your signature within this method.

```go
checkSignature := client.Base.CheckSignature(context.Background())

if !checkSignature.Success {
	panic(checkSignature.Error)
}

if checkSignature.Response.CheckSignatureResult {
    fmt.Println("Signature correct")
} else {
    fmt.Println("Signature incorrect", checkSignature.Response.Errors)
}
```

### Fetch available currencies

Get list of available currencies for depositing/withdrawing

```go
availableCurrencies := client.Base.AvailableCurrencies(context.Background())

if !availableCurrencies.Success {
    panic(availableCurrencies.Error)
}

list := *availableCurrencies.Response

for _, currency := range list {
	fmt.Println("%s (%s) = %s", currency.Currency, currency.Alias, currency.PriceUSD)

    if len(currency.Networks) > 0 {
        fmt.Println("\tnetworks:");

		for _, network := range currency.Networks {
			fmt.Println("\t\t%s (%s)", network.Name, network.Alias)
        }
    }
}
```

### Get currencies price-rate

Get price rate from one currency to another


```go
price := client.Base.PriceRate(context.Background(), "ETH", "USDT")

if !price.Success {
    panic(price.Error)
}

fmt.Println("Price: %s", price.Response)
```

### Get advanced balances info

Get info about advanced balance by its id

```go
balance := client.AdvancedBalance.AdvancedBalance(context.Background(), balanceId)

if !balance.Success {
    panic(balance.Error)
}

fmt.Println("[%s] (%s)\n\tAvalable for deposit: %s", balance.Response.AdvancedBalanceId, balance.Response.Currency, string.Join(balance.Response.AvailableCurrenciesForDeposit, ", "))
```

Or get list of advanced balances of user

```go
balances := client.AdvancedBalance.AdvancedBalance(context.Background(), balanceId)

if !balances.Success {
    panic(balances.Error)
}

list := *balances.Response

for _, balance := range list {
    fmt.Println("[%s] (%s)\n\tAvalable for deposit: %s", balance.AdvancedBalanceId, balance.Currency, string.Join(balance.AvailableCurrenciesForDeposit, ", "))
}
```

### Create order

```go
package main

import (
	"context"
	"errors"
	"github.com/apollopayment/go-sdk"
	"github.com/apollopayment/go-sdk/types/requests"
)

func main() {
	orderLink, err := createOrder("USDT", "tron", "1000")
	
	if err != nil {
		panic(err)
	}
	
	//
}

func createOrder(currency, network, amount string) (string, error) {
	ctx := context.Background()

	client := apollopayment_sdk.New("__PUBLIC_KEY__", "__PRIVATE_KEY__")

	advancedBalances := client.AdvancedBalance.AdvancedBalances(ctx)

	if !advancedBalances.Success {
		return "", errors.New("Error on get advanced balances list: " + advancedBalances.Error.Message)
	}

	_advancedBalances := *advancedBalances.Response
	advancedBalance := _advancedBalances[0]

	order := client.Orders.MakeOrder(ctx, requests.CreateOrder{
		AdvancedBalanceId: advancedBalance.AdvancedBalanceId,
		Currency:          currency,
		Network:           network,
		Amount:            amount,
		ErrorWebhook:      "https://merchant.domain/webhook-url",
		SuccessWebhook:    "https://merchant.domain/webhook-url",
		ReturnUrl:         "https://merchant.domain",
		Order:             "Order #1234567",
		Description:       "Buy some item",
	})

	if !order.Success {
		return "", errors.New("Error on create order: " + order.Error.Message)
	}

	return order.Response.Link, nil
}
```

### Auto-swap to external address

```go
package main

import (
	"context"
	"errors"
	"github.com/apollopayment/go-sdk"
	"github.com/apollopayment/go-sdk/types/requests"
)

func main() {
	autoSwapId, err := makeWithdrawal("USDT", "tron", "TUfVHn...DDC", "100")
	
	if err != nil {
		panic(err)
	}
	
	// 
}

func makeWithdrawal(currency, network, address, amount string) (string, error) {
	ctx := context.Background()

	client := apollopayment_sdk.New("__PUBLIC_KEY__", "__PRIVATE_KEY__")

	swap := client.AutoSwaps.Create(ctx, requests.CreateAutoSwap{
		Address:    address,
		Currency:   currency,
		Network:    network,
		AmountTo:   amount,
		WebhookUrl: "https://merchant.domain/webhook-url",
	})

	if !swap.Success {
		return "", errors.New("Error on make withdrawal: " + swap.Error.Message)
	}

	return swap.Response.Id, nil
}
```
