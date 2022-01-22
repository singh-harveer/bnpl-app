package mocks_test

import (
	"bnpl/mocks"
	"bnpl/service"
	"context"
	"errors"
	"testing"
)

func TestAddAndGetUsers(t *testing.T) {
	t.Parallel()
	var ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var merchantMngr, err = mocks.NewLocalDBUserManager()
	if err != nil {
		t.Fatalf("failed to create users manager object :%v", err)
	}

	var testcases = []struct {
		name string
		in   *service.User
		err  error
	}{
		{
			name: "Ok",
			in: &service.User{
				Name:        "user1",
				Email:       "u1@email.com",
				CreditLimit: 1000,
			},
		},
		{
			name: "duplicateEmail",
			in: &service.User{
				Name:        "user1",
				Email:       "u1@email.com",
				CreditLimit: 1000,
			},
			err: errors.New("unique constraint error"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var err = merchantMngr.AddUser(ctx, tc.in)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got %v want %v", err, tc.err)
			}

			if tc.err != nil {
				// If error is expected skip further validtions
				return
			}
			var got service.User
			got, err = merchantMngr.GetUserByName(ctx, tc.in.Name)
			if err != nil {
				t.Fatalf("failed to retrieve user: %v", err)
			}
			if !comPareTwoUsers(t, &got, tc.in) {
				t.Fatalf("got %v want %v", got, tc.in)
			}
		})
	}
}

func TestAddAndDeleteUsers(t *testing.T) {
	t.Parallel()
	var ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var merchantMngr, err = mocks.NewLocalDBUserManager()
	if err != nil {
		t.Fatalf("failed to create user manager object :%v", err)
	}

	var testcases = []struct {
		name string
		in   *service.User
		err  error
	}{
		{
			name: "Ok",
			in: &service.User{
				Name:        "user10",
				Email:       "u10@email.com",
				CreditLimit: 1000,
			},
			err: errors.New("not found"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var err = merchantMngr.AddUser(ctx, tc.in)
			if err != nil {
				t.Fatalf("failed to add user: %v", err)
			}

			err = merchantMngr.DeleteUserByName(ctx, tc.in.Name)
			if err != nil {
				t.Fatalf("failed to delete user: %v", err)
			}

			var got service.User
			got, err = merchantMngr.GetUserByName(ctx, tc.in.Name)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got %v want %v", err, tc.err)
			}

			if tc.err != nil {
				// If error is expected skip further validtions
				return
			}
			if !comPareTwoUsers(t, &got, tc.in) {
				t.Fatalf("got %v want %v", got, tc.in)
			}
		})
	}
}

func TestAddAndGetAllUsers(t *testing.T) {
	t.Parallel()
	var ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var merchantMngr, err = mocks.NewLocalDBUserManager()
	if err != nil {
		t.Fatalf("failed to create user manager object :%v", err)
	}

	var users = []*service.User{
		{
			Name:        "u5",
			Email:       "u5@email.com",
			CreditLimit: 1000,
		},
		{
			Name:        "u6",
			Email:       "u6@email.com",
			CreditLimit: 1000,
		},
		{
			Name:        "u7",
			Email:       "u7@email.com",
			CreditLimit: 1000,
		},
	}

	for _, v := range users {
		err = merchantMngr.AddUser(ctx, v)
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
			var got, err = merchantMngr.GetAllUsers(ctx)
			if err != nil {
				t.Fatalf("failed to retrieve users: %v", err)
			}

			if len(got) != 3 {
				t.Fatalf("failed to retrives all users")
			}
		})
	}
}

func comPareTwoUsers(t *testing.T, u1, u2 *service.User) bool {
	t.Helper()

	if u1.ID != u2.ID || u1.Name != u2.Name ||
		u1.Email != u2.Email || u1.CreditLimit != u2.CreditLimit {
		return false
	}

	return true
}
