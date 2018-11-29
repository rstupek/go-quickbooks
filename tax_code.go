package quickbooks

type TaxCode struct {
	Name                string              `json:"Name"`
	Description         string              `json:"Description"`
	Active              bool                `json:"Active"`
	Taxable             bool                `json:"Taxable"`
	TaxGroup            bool                `json:"TaxGroup"`
	SalesTaxRateList    SalesTaxRateList    `json:"SalesTaxRateList"`
	PurchaseTaxRateList PurchaseTaxRateList `json:"PurchaseTaxRateList"`
	Domain              string              `json:"domain"`
	Sparse              bool                `json:"sparse"`
	ID                  string              `json:"Id"`
	SyncToken           string              `json:"SyncToken"`
	MetaData            struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData"`
}

type SalesTaxRateList struct {
	TaxRateDetail []TaxRateDetail `json:"TaxRateDetail"`
}

type PurchaseTaxRateList struct {
	TaxRateDetail []TaxRateDetail `json:"TaxRateDetail"`
}

type TaxRateRef struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type TaxRateDetail struct {
	TaxRateRef        TaxRateRef `json:"TaxRateRef"`
	TaxTypeApplicable string     `json:"TaxTypeApplicable"`
	TaxOrder          int        `json:"TaxOrder"`
}
