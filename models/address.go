package models

// USAddress reperesents a U.S. Postal address
type USAddress struct {
	StreetNumber   string `json:"streetNumber"`             // ONLY the house/building number/letter
	StreeDirection string `json:"streeDirection,omitempty"` // i.e. North, South...
	StreetName     string `json:"streetName"`               // The name of the street
	StreetLabel    string `json:"streetLabel,omitempty"`    // i.e. Road, Street, Court...
	UnitNumber     string `json:"unitNumber,omitempty"`     // i.e. Apartment number, Suite number... or address line two in normal addressing terms
	City           string `json:"city"`                     // City name
	State          string `json:"state"`                    // State as a two letter abbreviation
	Zip            string `json:"zip"`                      // 5 digit zip
	Zip4           string `json:"zip4,omitempty"`           // 4 digit new Zip plus 4
}

// USAddressLabel reperesents a U.S.A. postal address and it's type
type USAddressLabel struct {
	USAddress
	Type string // i.e. Work, Home, Vacation
}
