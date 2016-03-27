package models

// Order reperesents an order
type Order struct {
	ID        string
	PatientID string
	Items     []OrderItem
}

// OrderItem reperesents an orderItem
type OrderItem struct {
	Product InventoryItem
}
