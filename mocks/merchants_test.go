package mocks_test

import (
	"bnpl/mocks"
	"bnpl/service"
	"context"
	"errors"
	"testing"
	"time"
)

const (
	timeout = 120 * time.Second
)

func TestAddAndGetMerchants(t *testing.T) {
	t.Parallel()
	var ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var merchantMngr, err = mocks.NewLocalDBMerchantManager()
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
				Name:     "merchant1",
				Email:    "m1@email.com",
				Discount: 10,
			},
		},
		{
			name: "duplicateEmail",
			in: &service.Merchant{
				Name:     "merchant1",
				Email:    "m1@email.com",
				Discount: 10,
			},
			err: errors.New("unique constraint error"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var err = merchantMngr.AddMerchant(ctx, tc.in)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got %v want %v", err, tc.err)
			}

			if tc.err != nil {
				// If error is expected skip further validtions
				return
			}
			var got service.Merchant
			got, err = merchantMngr.GetMerchantByName(ctx, tc.in.Name)
			if err != nil {
				t.Fatalf("failed to retrieve merchant: %v", err)
			}
			if !comPareTwoMerchants(t, &got, tc.in) {
				t.Fatalf("got %v want %v", got, tc.in)
			}
		})
	}
}

func TestAddAndDeleteMerchants(t *testing.T) {
	t.Parallel()
	var ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var merchantMngr, err = mocks.NewLocalDBMerchantManager()
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
				Name:     "merchant10",
				Email:    "m10@email.com",
				Discount: 10,
			},
			err: errors.New("not found"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var err = merchantMngr.AddMerchant(ctx, tc.in)
			if err != nil {
				t.Fatalf("failed to add merchant: %v", err)
			}

			err = merchantMngr.DeleteMerchantByName(ctx, tc.in.Name)
			if err != nil {
				t.Fatalf("failed to delete merchant: %v", err)
			}

			var got service.Merchant
			got, err = merchantMngr.GetMerchantByName(ctx, tc.in.Name)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got %v want %v", err, tc.err)
			}

			if tc.err != nil {
				// If error is expected skip further validtions
				return
			}
			if !comPareTwoMerchants(t, &got, tc.in) {
				t.Fatalf("got %v want %v", got, tc.in)
			}
		})
	}
}

func TestAddAndGetAllMerchants(t *testing.T) {
	t.Parallel()
	var ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var merchantMngr, err = mocks.NewLocalDBMerchantManager()
	if err != nil {
		t.Fatalf("failed to create merchant manager object :%v", err)
	}

	var merchants = []*service.Merchant{
		{
			Name:     "m5",
			Email:    "m5@email.com",
			Discount: 10,
		},
		{
			Name:     "m6",
			Email:    "m6@email.com",
			Discount: 15,
		},
		{
			Name:     "m7",
			Email:    "m7@email.com",
			Discount: 5,
		},
	}

	for _, v := range merchants {
		err = merchantMngr.AddMerchant(ctx, v)
		if err != nil {
			t.Fatalf("failed to add merchant: %v", err)
		}
	}

	var testcases = []struct {
		name string
		in   []*service.Merchant
		err  error
	}{
		{
			name: "Ok",
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			var got, err = merchantMngr.GetAllMerchants(ctx)
			if err != nil {
				t.Fatalf("failed to retrieve merchants: %v", err)
			}

			if len(got) != 3 {
				t.Fatalf("failed to retrives all merchants")
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
