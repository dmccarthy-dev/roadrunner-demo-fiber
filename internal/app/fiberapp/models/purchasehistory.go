package models

type PurchaseHistoryEntry struct {
	Amount int    `json:"amount"`
	Item   string `json:"item"`
}

type PurchaseHistory struct {
	Id      string                 `json:"id,omitempty"`
	History []PurchaseHistoryEntry `json:"history,omitempty"`
}
