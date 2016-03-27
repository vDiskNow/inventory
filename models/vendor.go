package models

// USVendorCompany reperesents a US based company
type USVendorCompany struct {
	ID          string
	Name        string
	URL         string
	LocationIDs []string
}

// USVendorLocation reperesents a US based company's location(s)
type USVendorLocation struct {
	ID        string
	CompanyID string
	Name      string
	Address   USAddress
	ShipTo    USAddress
	BillTo    USAddress
}

// USVendorEmployee reperesents an employee of a US based company
type USVendorEmployee struct {
	USPerson
	CompanyID string
	Title     string
}
