package test

import (
	"testing"

	randomdata "github.com/Pallinder/go-randomdata"
	seed "github.com/jinmatt/go-seed-rand"
	"github.com/omniboost/go-quickbooks"
	"github.com/omniboost/go-quickbooks/sdk/consts"
	"github.com/tylerb/is"
)

func TestCreateAccount(t *testing.T) {
	is := is.New(t)

	qbo := quickbooks.NewClient(RealmID, AccessToken, true)

	account := quickbooks.Account{}
	account.Name = randomdata.SillyName() + seed.RandomKey(7)
	account.AccountType = consts.QBAccountIncomeType

	newAccount, err := qbo.CreateAccount(account)
	is.NotErr(err)
	is.NotNil(newAccount.Account.ID)
	is.Equal(account.Name, newAccount.Account.Name)
}
