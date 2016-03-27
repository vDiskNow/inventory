package models

import "time"

// USPatient reperesents a US based patient
type USPatient struct {
	USPerson
	VerificationCode string
	Issued           time.Time
	ValidUntil       time.Time
	VerifyURL        string
	VerifyPhone      *PhoneNumber
}
