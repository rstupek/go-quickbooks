package quickbooks

import (
	"encoding/json"
	"fmt"
)

// SalesReceiptObject the complete quickbooks sales receipt object type
type SalesReceiptObject struct {
	SalesReceipt SalesReceipt `json:"SalesReceipt"`
	Time         string       `json:"time"`
}

// SalesReceipt quickbooks sales receipt type
type SalesReceipt struct {
	ID               string `json:"Id,omitempty"`
	DocNumber        string `json:"DocNumber,omitempty"`
	SyncToken        string `json:"SyncToken,omitempty"`
	Domain           string `json:"domain,omitempty"`
	Balance          int    `json:"Balance,omitempty"`
	PaymentMethodRef struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"PaymentMethodRef"`
	BillAddr            *Address `json:"BillAddr"`
	DepositToAccountRef struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"DepositToAccountRef"`
	TxnDate     string  `json:"TxnDate"`
	TotalAmt    float64 `json:"TotalAmt"`
	CustomerRef struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"CustomerRef"`
	CustomerMemo struct {
		Value string `json:"value"`
	} `json:"CustomerMemo"`
	PrintStatus           string             `json:"PrintStatus,omitempty"`
	PaymentRefNum         string             `json:"PaymentRefNum"`
	EmailStatus           string             `json:"EmailStatus,omitempty"`
	Sparse                bool               `json:"sparse,omitempty"`
	Line                  []SalesReceiptLine `json:"Line"`
	ApplyTaxAfterDiscount bool               `json:"ApplyTaxAfterDiscount,omitempty"`
	CustomField           *[]CustomField     `json:"CustomField,omitempty"`
	TxnTaxDetail          *struct {
		TxnTaxCodeRef *TaxCodeRef `json:"TxnTaxCodeRef,omitempty"`
		TotalTax      float64     `json:"TotalTax"`
		TaxLine       []TaxLine   `json:"TaxLine,omitempty"`
	} `json:"TxnTaxDetail,omitempty"`
	MetaData *struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData,omitempty"`
}

// SalesReceiptLine quickbooks sales receipt line item object
type SalesReceiptLine struct {
	Description         string               `json:"Description,omitempty"`
	DetailType          string               `json:"DetailType"`
	SalesItemLineDetail *SalesItemLineDetail `json:"SalesItemLineDetail,omitempty"`
	LineNum             int                  `json:"LineNum,omitempty"`
	Amount              float64              `json:"Amount"`
	ID                  string               `json:"Id,omitempty"`
	SubTotalLineDetail  struct {
	} `json:"SubTotalLineDetail,omitempty"`
}

// CreateSalesReceipt creates an sales receipt on quickbooks
func (q *Quickbooks) CreateSalesReceipt(invoice SalesReceipt) (*SalesReceiptObject, error) {
	endpoint := fmt.Sprintf("/company/%s/salesreceipt", q.RealmID)

	res, err := q.makePostRequest(endpoint, invoice)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	newSalesReceipt := SalesReceiptObject{}
	err = json.NewDecoder(res.Body).Decode(&newSalesReceipt)
	if err != nil {
		return nil, err
	}

	return &newSalesReceipt, nil
}
