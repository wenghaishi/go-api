package tools

import "time"

type mockDb struct{}

var mockLoginDetails = map[string]LoginDetail{
	"alex": {
		AuthToken: "12345abc",
		Username:  "alex",
	},
	"john": {
		AuthToken: "12345def",
		Username:  "john",
	},
	"marie": {
		AuthToken: "12345ghi",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"john": {
		Coins:    200,
		Username: "john",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

func (d *mockDb) GetUserLoginDetails(username string) *LoginDetail {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetail{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}
