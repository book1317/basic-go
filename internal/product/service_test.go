package product

import (
	"academy/internal/app"
	"academy/internal/model"
	"fmt"
	"reflect"
	"testing"
)

var Mock_Merchant = []model.Merchant{
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

// Product in Mongo
/* 1 */
// {
//     "_id" : ObjectId("5fc08fcac1bd846579fe6ed3"),
//     "merchant_id" : ObjectId("5fc07c40ff6e6d7a37f761b8"),
//     "name" : "Coke",
//     "amount" : 100.01,
//     "stocks" : 5
// }

// /* 2 */
// {
//     "_id" : ObjectId("5fc090aac1bd846579fe6f9b"),
//     "merchant_id" : ObjectId("5fc07c40ff6e6d7a37f761b8"),
//     "name" : "Coke",
//     "amount" : 250.35,
//     "stocks" : 1
// }

// /* 3 */
// {
//     "_id" : ObjectId("5fc093e947998d3c09a788cc"),
//     "merchant_id" : ObjectId("5fc07c40ff6e6d7a37f761b8"),
//     "name" : "Spize",
//     "amount" : 129.0,
//     "stocks" : 2
// }

// /* 4 */
// {
//     "_id" : ObjectId("5fc094e459b6ecc3d87ea0e1"),
//     "merchant_id" : ObjectId("5fc07c40ff6e6d7a37f761b8"),
//     "name" : "Lactasoy",
//     "amount" : 5.5,
//     "stocks" : 2
// }

// /* 5 */
// {
//     "_id" : ObjectId("5fc0952bf1866762da955488"),
//     "merchant_id" : ObjectId("5fc07c40ff6e6d7a37f761b8"),
//     "name" : "Vitamilke",
//     "amount" : 10.0,
//     "stocks" : 2
// }

func TestService_GetProductsByMerchantId(t *testing.T) {
	tsc := []struct {
		merchantID   string
		expectedData []model.Product
	}{
		{
			merchantID: "5fc07c40ff6e6d7a37f761b9",
			expectedData: []model.Product{
				{
					Name:   "KK",
					Amount: 10.0,
					Stocks: 2,
				},
			},
		},
	}

	state := "dev"
	c := app.NewConfig(state)
	if err := c.Init_Test("../../configs"); err != nil {
		fmt.Println("config init error", err)
		return
	}

	productRepo := NewRepo(c)
	srv := NewService(c, productRepo)
	for _, tc := range tsc {
		data, err := srv.GetListAllProductByMerchantId(nil, tc.merchantID)
		if err != nil {
			t.Error("found error ==>", err)
			continue
		}

		if !reflect.DeepEqual(data, tc.expectedData) {
			t.Errorf("expected data %+v, but got %+v", tc.expectedData, data)
		}
	}
}
