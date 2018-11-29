package test

import (
	"testing"

	quickbooks "github.com/omniboost/go-quickbooks"
	"github.com/tylerb/is"
)

func TestDiscovery(t *testing.T) {
	is := is.New(t)
	discovery, err := quickbooks.NewDiscovery(true)
	is.NotErr(err)
	is.NotNil(discovery.AuthorizationEndpoint)
}
