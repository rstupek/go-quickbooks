package quickbooks

import (
	"golang.org/x/oauth2"
)

const (
	AccountingScope          = "com.intuit.quickbooks.accounting"
	PaymentsScope            = "com.intuit.quickbooks.payment"
	PayrollScope             = "com.intuit.quickbooks.payroll"
	PayrollTimetrackingScope = "com.intuit.quickbooks.payroll.timetracking"
	PayrollBenefitsScope     = "com.intuit.quickbooks.payroll.benefits"
)

type Oauth2Config struct {
	oauth2.Config
}

func NewOauth2ConfigProduction() (*Oauth2Config, error) {
	return NewOauth2Config(false)
}

func NewOauth2ConfigSandbox() (*Oauth2Config, error) {
	return NewOauth2Config(true)
}

func NewOauth2Config(isSandbox bool) (*Oauth2Config, error) {
	discovery, err := NewDiscovery(isSandbox)
	if err != nil {
		return nil, err
	}

	config := &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes: []string{
				AccountingScope,
				PaymentsScope,
				PayrollScope,
				PayrollTimetrackingScope,
				PayrollBenefitsScope,
			},
			Endpoint: oauth2.Endpoint{
				AuthURL:  discovery.AuthorizationEndpoint,
				TokenURL: discovery.TokenEndpoint,
			},
		},
	}

	return config, nil
}
