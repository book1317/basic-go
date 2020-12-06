package merchant

import (
	"academy/internal/app"
	"academy/internal/model"
	"fmt"
	"testing"
)

var Mock_Data = []model.Merchant{
	{
		// ID:          ObjectId("5fc07c40ff6e6d7a37f761b8"),
		Name:        "Bookie",
		BankAccount: "45678",
		Username:    "1zBimVd1yQ",
		Password:    "DPuMm94oAGzajcF2YGWT",
	},
	{
		// ID:          ObjectId("5fc09f80051cad00e9b0de2c"),
		Name:        "jame",
		BankAccount: "12345678dd",
		Username:    "BRqJqQKsL0",
		Password:    "hhyokaAt0D0CKtZBl8AQ",
	},
	{
		// ID:          ObjectId("5fc09fbb89c66ad962c5f19d"),
		Name:        "john",
		BankAccount: "5436666",
		Username:    "oG9Q3f9yyu",
		Password:    "c5h6C1OYKh9rSqTtOJh7",
	},
	{
		// ID:          ObjectId("5fc09fc389c66ad962c5f19e"),
		Name:        "michale",
		BankAccount: "8888888",
		Username:    "nVpXb6IHc9",
		Password:    "LDDpzqHFULRXsun3Fh9z",
	},
}

func TestService_GetMerchantById(t *testing.T) {
	tsc := []struct {
		merchantID   string
		expectedData model.Merchant
	}{
		{
			merchantID: "5fc07c40ff6e6d7a37f761b8",
			expectedData: model.Merchant{
				Name:        "Bookie",
				BankAccount: "45678",
				Username:    "1zBimVd1yQ",
				Password:    "DPuMm94oAGzajcF2YGWT",
			},
		},
		{
			merchantID: "5fc09f80051cad00e9b0de2c",
			expectedData: model.Merchant{
				Name:        "jame",
				BankAccount: "12345678dd",
				Username:    "BRqJqQKsL0",
				Password:    "hhyokaAt0D0CKtZBl8AQ",
			},
		},
	}

	state := "dev"
	c := app.NewConfig(state)
	if err := c.Init_Test("../../configs"); err != nil {
		fmt.Println("config init error", err)
		return
	}

	merchantRepo := NewRepo(c)
	srv := NewService(c, merchantRepo)
	for _, tc := range tsc {
		data, err := srv.GetMerchantById(nil, tc.merchantID)
		if err != nil {
			t.Error("found error ==>", err)
			continue
		}

		if data != tc.expectedData {
			t.Errorf("expected data %+v, but got %+v", tc.expectedData, data)
		}
	}
}
