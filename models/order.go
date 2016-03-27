package models

// Order reperesents an order
type Order struct {
	ID        string
	PatientID string
	Items     []*OrderItem
	Total     float64
}

// OrderItem reperesents an orderItem
type OrderItem struct {
	Product  InventoryItem
	Quantity int
	Total    float64
}

// TotalOrder return the order total
func (o *Order) TotalOrder() float64 {
	var total float64
	for _, item := range o.Items {
		total += item.Total
	}
	return total
}

// TotalItem returns the total for this line item
func (ot *OrderItem) TotalItem() float64 {
	return float64(ot.Quantity) * ot.Product.PriceUSD
}
