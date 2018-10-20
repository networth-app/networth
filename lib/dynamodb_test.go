package nwlib

import (
	_ "github.com/networth-app/networth/dotenv"

	"testing"

	"github.com/plaid/plaid-go/plaid"
	"github.com/stretchr/testify/assert"
)

var (
	db                   = NewDynamoDBClient()
	institutionID        = "ins_1"
	testUsername         = "b6989907-cba1-4ffb-b0f3-1258cb689ba0"
	testUsernameNotExist = "test_not_exist_username@networth.app"
	accountID            = "Vxk3QMnVmNhaJKmlrXg5tj7q5keD3bfW4BnnE"
)

func TestSetTransaction(t *testing.T) {
	username := "test_set_transaction@networth.app"
	trans := plaid.Transaction{
		ID:        "1",
		AccountID: "1",
		Amount:    1,
	}

	err := db.SetTransaction(username, trans)
	assert.Equal(t, err, nil)
}

func TestSetAccount(t *testing.T) {
	username := "test_set_account@networth.app"
	itemID := "1"
	account := &plaid.Account{
		AccountID: "1",
		Name:      "test",
	}

	// set for specific ins
	err := db.SetAccount(username, institutionID, account)
	assert.Equal(t, err, nil)

	// set for all
	err = db.SetAccount(username, itemID, account)
	assert.Equal(t, err, nil)
}

func TestSetToken(t *testing.T) {
	db := NewDynamoDBClient()
	username := "test_set_token@networth.app"
	token := &Token{
		ItemID:        "1",
		AccessToken:   "1",
		InstitutionID: institutionID,
	}

	err := db.SetToken(username, token)
	assert.Equal(t, nil, err)
}

// func TestGetToken(t *testing.T) {
// 	db := NewDynamoDBClient()
// 	username := "test_set_token@networth.app"
// 	invalidInstitutionID := "ins_1_invalid"

// 	// get without ins_id
// 	tokens := db.GetToken(username, "")
// 	assert.Equal(t, tokens.Tokens[len(tokens.Tokens)-1].InstitutionID, institutionID)

// 	// get using ins_id
// 	tokens = db.GetToken(username, institutionID)
// 	assert.Equal(t, tokens.Tokens[len(tokens.Tokens)-1].InstitutionID, institutionID)

// 	// // get using invalid ins_id
// 	tokens = db.GetToken(username, invalidInstitutionID)
// 	assert.Equal(t, len(tokens.Tokens) == 0, true)
// }

func TestGetAccounts(t *testing.T) {
	db := NewDynamoDBClient()

	// get using existing username
	accounts, err := db.GetAccounts(testUsername)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(accounts) > 0, true)

	// get using non-exist username
	accounts, err = db.GetAccounts(testUsernameNotExist)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(accounts), 0)
}

func TestGetNetworth(t *testing.T) {
	db := NewDynamoDBClient()

	// get using existing username
	networth, err := db.GetNetworth(testUsername)
	assert.Equal(t, err, nil)
	assert.Equal(t, networth.Networth > 0, true)
	assert.Equal(t, networth.Assets > 0, true)
	assert.Equal(t, networth.Liabilities > 0, true)

	// get using non-exist username
	networth, err = db.GetNetworth(testUsernameNotExist)
	assert.Equal(t, err, nil)
	assert.Equal(t, networth.Networth, 0.0)
	assert.Equal(t, networth.Assets, 0.0)
	assert.Equal(t, networth.Liabilities, 0.0)
}

func TestGetNetworthByDateRange(t *testing.T) {
	db := NewDynamoDBClient()
	startDate := "2018-10-06"
	endDate := "2018-10-13"
	badStartDate := "2050-10-06"
	badEndDate := "2050-10-13"

	// get using existing range
	networth, err := db.GetNetworthByDateRange(testUsername, startDate, endDate)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(networth) > 0, true)
	assert.Equal(t, networth[0].Networth > 0, true)

	// get using invalid range
	networth, err = db.GetNetworthByDateRange(testUsername, badStartDate, badEndDate)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(networth), 0)
}

// func TestGetTransaction(t *testing.T) {
// 	// get using existing username
// 	transactions, err := db.GetTransactions(testUsername, accountID)
// 	assert.Equal(t, err, nil)
// 	assert.Equal(t, len(transactions) > 0, true)

// 	// get using non-exist username
// 	// accounts, err = db.GetAccounts(testUsernameNotExist)
// 	// assert.Equal(t, err, nil)
// 	// assert.Equal(t, len(accounts), 0)
// }

func TestGetAllUsers(t *testing.T) {
	_, err := db.GetAllUsers()
	assert.Equal(t, err, nil)
}
