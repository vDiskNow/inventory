package models

import "fmt"

// PersonName reperesents the name of a person
type PersonName struct {
	Given         string `json:"given"`
	MiddleInitial string `json:"middleInitial"`
	SirName       string `json:"sirName"`
}

// USPerson reperesents a person living in the U.S.A.
type USPerson struct {
	ID        string
	Name      PersonName        `json:"name"`
	Addresses []*USAddressLabel `json:"addresses"`
	Phones    []*PhoneLabel     `json:"phones"`
	Emails    []*EmailLabel     `json:"emails"`
}

// GetFullName returns a string formatted with full display name
func (usp *USPerson) GetFullName() string {
	return fmt.Sprintf("%s %s. %s", usp.Name.Given, usp.Name.MiddleInitial, usp.Name.SirName)
}
