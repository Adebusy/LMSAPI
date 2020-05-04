package model

import uuid "github.com/satori/go.uuid"

// Account example
type Account struct {
	ID   int       `json:"id" example:"1" format:"int64"`
	Name string    `json:"name" example:"account name"`
	UUID uuid.UUID `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

//AccountOne for account
func AccountOne(id int) (Account, error) {
	for _, v := range accounts {
		if id == v.ID {
			return v, nil
		}
	}
	return Account{}, nil
}

var accountMaxID = 3
var accounts = []Account{
	{ID: 1, Name: "account_1"},
	{ID: 2, Name: "account_2"},
	{ID: 3, Name: "account_3"},
}
