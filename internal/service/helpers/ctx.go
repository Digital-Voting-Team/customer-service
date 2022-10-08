package helpers

import (
	"context"
	"customer-service/internal/data"

	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	addressesQCtxKey
	personsQCtxKey
	customersQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxAddressesQ(entry data.AddressesQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, addressesQCtxKey, entry)
	}
}

func AddressesQ(r *http.Request) data.AddressesQ {
	return r.Context().Value(addressesQCtxKey).(data.AddressesQ).New()
}

func CtxPersonsQ(entry data.PersonsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, personsQCtxKey, entry)
	}
}

func PersonsQ(r *http.Request) data.PersonsQ {
	return r.Context().Value(personsQCtxKey).(data.PersonsQ).New()
}

func CtxCustomersQ(entry data.CustomersQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, customersQCtxKey, entry)
	}
}

func CustomersQ(r *http.Request) data.CustomersQ {
	return r.Context().Value(customersQCtxKey).(data.CustomersQ).New()
}
