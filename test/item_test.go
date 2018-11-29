package test

import (
	"testing"

	randomdata "github.com/Pallinder/go-randomdata"
	seed "github.com/jinmatt/go-seed-rand"
	"github.com/omniboost/go-quickbooks"
	"github.com/omniboost/go-quickbooks/sdk/consts"
	"github.com/tylerb/is"
)

func TestCreateItem(t *testing.T) {
	is := is.New(t)

	qbo := quickbooks.NewClient(RealmID, AccessToken, true)

	// create a chart of account for item
	account := quickbooks.Account{}
	account.Name = randomdata.SillyName() + seed.RandomKey(7)
	account.AccountType = consts.QBAccountIncomeType
	newAccount, err := qbo.CreateAccount(account)
	is.NotErr(err)

	item := quickbooks.Item{}
	item.Name = randomdata.SillyName() + seed.RandomKey(7)
	item.IncomeAccountRef = &quickbooks.AccountRef{
		Value: newAccount.Account.ID,
		Name:  newAccount.Account.Name,
	}
	item.Type = consts.QBItemServiceType

	newItem, err := qbo.CreateItem(item)
	is.NotErr(err)
	is.NotNil(newItem.Item.ID)
	is.Equal(item.Name, newItem.Item.Name)
}
