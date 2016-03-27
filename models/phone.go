package models

import "fmt"

// PhoneNumber reperesents a Telephone number
type PhoneNumber struct {
	CountryCode string `json:"countryCode"`       // Country code for the phone... U.S.A = 1
	Number      string `json:"number"`            // Phone number
	Carrier     string `json:"carrier,omitempty"` // i.e. Verizon, AT&T...
}

// PhoneLabel reperesents a phone and it's type
type PhoneLabel struct {
	PhoneNumber
	Type string
}

// GetE164 returns an E164 string for dialing
func (pn *PhoneNumber) GetE164() string {
	return fmt.Sprintf("+%s%s", pn.CountryCode, pn.Number)
}
