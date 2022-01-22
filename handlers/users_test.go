package handlers_test

import (
	"bnpl/service"
	"encoding/json"
	"testing"
)

func TestAddGetUser(t *testing.T) {
	// TODO-add test cases.

}

func prepareUserBody(t *testing.T, name, email string, creditLimt float64) []byte {
	var m = service.User{
		Name:        name,
		Email:       email,
		CreditLimit: creditLimt,
	}

	var body, err = json.Marshal(m)
	if err != nil {
		t.Fatalf("failed to prepare body:%v", err)
	}

	return body
}
