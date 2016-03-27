package models

// Product reperesents a product
type InventoryItem struct {
	ID          string
	Name        string
	Description string
	Category    string
	ImageURL    string
	Quantity    int
}
