package models

type Profile struct {
	Id            string          `json:"id"`
	Name          string          `json:"name"`
	Email         string          `json:"email"`
	PastPurchases PurchaseHistory `json:"purchase_history"`
}
