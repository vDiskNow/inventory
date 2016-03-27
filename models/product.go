package models

// InventoryItem reperesents a product in inventory
type InventoryItem struct {
	ID          string
	Name        string
	Description string
	PriceUSD    float64
	Category    string
	ImageURL    string
	Quantity    int
}
