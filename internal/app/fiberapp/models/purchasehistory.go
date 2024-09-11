package models

type PurchaseHistoryEntry struct {
	Amount int    `json:"amount"`
	Item   string `json:"item"`
}

type PurchaseHistory struct {
	Id      string                 `json:"id"`
	History []PurchaseHistoryEntry `json:"history"`
}
