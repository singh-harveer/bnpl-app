package handlers_test

import (
	"bnpl/service"
	"encoding/json"
	"testing"
)

func TestAddGetMerchant(t *testing.T) {

}

func prepareMerchantBody(t *testing.T, name, email string, discount float64) []byte {
	var m = service.Merchant{
		Name:     name,
		Email:    email,
		Discount: discount,
	}

	var body, err = json.Marshal(m)
	if err != nil {
		t.Fatalf("failed to prepare body:%v", err)
	}

	return body
}
