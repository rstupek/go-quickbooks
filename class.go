package quickbooks

type Class struct {
	Name               string `json:"Name"`
	SubClass           bool   `json:"SubClass"`
	FullyQualifiedName string `json:"FullyQualifiedName"`
	Active             bool   `json:"Active"`
	Domain             string `json:"domain"`
	Sparse             bool   `json:"sparse"`
	ID                 string `json:"Id"`
	SyncToken          string `json:"SyncToken"`
	MetaData           struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData"`
}
