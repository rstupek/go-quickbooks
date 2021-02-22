package quickbooks

import (
	"encoding/json"
	"fmt"
)

// CreditMemoObject the complete quickbooks creditmemo object type
type CreditMemoObject struct {
	CreditMemo CreditMemo `json:"CreditMemo"`
	Time       string     `json:"time"`
}

// CreditMemo quickbooks creditmemo type
type CreditMemo struct {
	ID           string           `json:"Id,omitempty"`
	Domain       string           `json:"domain,omitempty"`
	Sparse       bool             `json:"sparse,omitempty"`
	SyncToken    string           `json:"SyncToken,omitempty"`
	CustomField  *[]CustomField   `json:"CustomField,omitempty"`
	DocNumber    string           `json:"DocNumber,omitempty"`
	TxnDate      string           `json:"TxnDate,omitempty"`
	LinkedTxn    *[]LinkedTxn     `json:"LinkedTxn,omitempty"`
	Line         []CreditMemoLine `json:"Line"`
	TxnTaxDetail *struct {
		TxnTaxCodeRef *TaxCodeRef `json:"TxnTaxCodeRef,omitempty"`
		TotalTax      float64     `json:"TotalTax"`
		TaxLine       []TaxLine   `json:"TaxLine,omitempty"`
	} `json:"TxnTaxDetail,omitempty"`
	CustomerRef  *CustomerRef `json:"CustomerRef"`
	CustomerMemo *struct {
		Value string `json:"value"`
	} `json:"CustomerMemo,omitempty"`
	BillAddr     *Address `json:"BillAddr"`
	ShipAddr     *Address `json:"ShipAddr"`
	SalesTermRef *struct {
		Value string `json:"value"`
	} `json:"SalesTermRef,omitempty"`
	DueDate               string  `json:"DueDate,omitempty"`
	TotalAmt              float64 `json:"TotalAmt,omitempty"`
	ApplyTaxAfterDiscount bool    `json:"ApplyTaxAfterDiscount,omitempty"`
	PrintStatus           string  `json:"PrintStatus,omitempty"`
	EmailStatus           string  `json:"EmailStatus,omitempty"`
	BillEmail             *struct {
		Address string `json:"Address"`
	} `json:"BillEmail,omitempty"`
	Balance  float64 `json:"Balance,omitempty"`
	MetaData *struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData,omitempty"`
}

// CreditMemoLine quickbooks creditmemo line item object
type CreditMemoLine struct {
	ID                  string               `json:"Id,omitempty"`
	LineNum             int                  `json:"LineNum,omitempty"`
	Description         string               `json:"Description,omitempty"`
	Amount              float64              `json:"Amount"`
	DetailType          string               `json:"DetailType"`
	SalesItemLineDetail *SalesItemLineDetail `json:"SalesItemLineDetail,omitempty"`
	SubTotalLineDetail  interface{}          `json:"SubTotalLineDetail,omitempty"`
}

// CreateCreditMemo creates an creditmemo on quickbooks
func (q *Quickbooks) CreateCreditMemo(creditmemo CreditMemo) (*CreditMemoObject, error) {
	endpoint := fmt.Sprintf("/company/%s/creditmemo", q.RealmID)

	res, err := q.makePostRequest(endpoint, creditmemo)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	newCreditMemo := CreditMemoObject{}
	err = json.NewDecoder(res.Body).Decode(&newCreditMemo)
	if err != nil {
		return nil, err
	}

	return &newCreditMemo, nil
}
