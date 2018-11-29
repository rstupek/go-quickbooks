package test

import (
	"testing"

	"github.com/omniboost/go-quickbooks"
	"github.com/tylerb/is"
)

func TestCompanyInfo(t *testing.T) {
	is := is.New(t)

	qbo := quickbooks.NewClient(RealmID, AccessToken, true)

	company, err := qbo.GetCompanyInfo()
	is.NotErr(err)
	is.NotNil(company.CompanyInfo.CompanyName)
}
