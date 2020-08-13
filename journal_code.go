package quickbooks

import "time"

type JournalCode struct {
	SyncToken string    `json:"SyncToken"`
	Domain    string    `json:"domain"`
	Name      string    `json:"Name"`
	Sparse    bool      `json:"sparse"`
	Time      time.Time `json:"time"`
	Active    bool      `json:"Active"`
	MetaData  struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData"`
	Type        string `json:"Type"`
	ID          string `json:"Id"`
	Description string `json:"Description"`
}
