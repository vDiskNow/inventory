package models

// Email reperesents an e-mail address
type Email struct {
	Address string
}

// EmailLabel eperesents an e-mail address and it's type
type EmailLabel struct {
	Email
	Type string
}
