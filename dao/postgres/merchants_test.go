package postgres_test

import (
	"bnpl/dao/postgres"
	"bnpl/service"
	"context"
	"errors"
	"testing"
	"time"
)

const (
	timeout = 120 * time.Second
)

func TestAddGetAndDeleteMerchants(t *testing.T) {
	t.Parallel()
	var ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var merchantMngr, err = postgres.NewLocalDBMerchantManager()
	if err != nil {
		t.Fatalf("failed to create merchant manager object :%v", err)
	}

	var testcases = []struct {
		name string
		in   *service.Merchant
		err  error
	}{
		{
			name: "Ok",
			in: &service.Merchant{
				Name:     "merchant2",
				Email:    "m2@email.com",
				Discount: 10,
			},
		},
		{
			name: "duplicateEmail",
			in: &service.Merchant{
				Name:     "merchant1",
				Email:    "m3@email.com",
				Discount: 10,
			},
			err: errors.New("unique constraint error"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

			var err = merchantMngr.AddMerchant(ctx, tc.in)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got %v want %v", err, tc.err)
			}

			var got service.Merchant
			got, err = merchantMngr.GetMerchantByID(ctx, service.ID(tc.in.ID))
			if err != nil {
				t.Fatalf("failed to retrieve merchant: %v", err)
			}
			if !comPareTwoMerchants(t, &got, tc.in) {
				t.Fatalf("got %v want %v", got, tc.in)
			}

			err = merchantMngr.DeleteMerchant(ctx, service.ID(tc.in.ID))
			if err != nil {
				t.Fatalf("failed to delete merchant: %v", err)
			}
		})
	}
}

func comPareTwoMerchants(t *testing.T, m1, m2 *service.Merchant) bool {
	t.Helper()

	if m1.ID != m2.ID || m1.Name != m2.Name ||
		m1.Email != m2.Email || m1.Discount != m2.Discount {
		return false
	}

	return true
}
