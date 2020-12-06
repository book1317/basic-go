package model

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Merchant struct {
	Name        string `json:"name" bson:"name"`
	BankAccount string `json:"bank_account" bson:"bank_account"`
	Username    string `json:"username" bson:"username"`
	Password    string `json:"password" bson:"password"`
}

type Product struct {
	Name   string  `json:"name" bson:"name"`
	Amount float64 `json:"amount" bson:"amount"`
	Stocks int     `json:"stocks" bson:"stocks"`
}
