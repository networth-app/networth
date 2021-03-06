package nwlib

import (
	"log"
	"strings"
)

// Networth struct
type Networth struct {
	Networth    float64 `json:"networth"`
	Assets      float64 `json:"assets"`
	Liabilities float64 `json:"liabilities"`
	DateTime    string  `json:"sort"`
	UpdatedAt   string  `json:"updated_at"`
}

// SyncNetworth save net worth to db for current datetime
func SyncNetworth(db *DynamoDBClient, username string) error {
	assets := 0.0
	liabilities := 0.0
	accountCache := make(map[string]bool)
	accounts, err := db.GetAccounts(username)

	if err != nil {
		log.Printf("Problem getting accounts: %+v\n", err)
		return err
	}

	for _, account := range accounts {
		if _, ok := accountCache[account.AccountID]; !ok {
			if isAsset(account) {
				assets = assets + account.Balances.Current
			} else {
				liabilities = liabilities + account.Balances.Current
			}
			accountCache[account.AccountID] = true
		}
	}

	networth := assets - liabilities

	if err := db.SetNetworth(username, networth, assets, liabilities); err != nil {
		log.Printf("Problem setting networth: %+v\n", err)
		return err
	}

	return nil
}

// List of account type and subtype https://plaid.com/docs/#accounts
func isAsset(account Account) bool {
	switch account.Type {
	case "brokerage", "depository":
		return true
	case "loan", "mortgage":
		return false
	}

	switch account.Subtype {
	case "credit card", "line of credit":
		return false
	case "paypal", "403B", "cash management", "cd", "hsa", "keogh", "money market", "mutual fund", "prepaid", "recurring", "rewards", "safe deposit", "sarsep":
		return true
	}

	if account.Type == "other" && account.Subtype == "other" {
		if strings.ToLower(account.OfficialName) == "investment" {
			return true
		}
	}

	return false
}
