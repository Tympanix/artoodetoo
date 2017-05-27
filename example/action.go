package example

import (
	"log"

	"github.com/Tympanix/automato/unit"
)

// EmailAction mimcs sending an email as an action
type EmailAction struct {
	Receiver string `io:"input"`
	Subject  string `io:"input"`
	Message  string `io:"input"`
}

func init() {
	unit.Register(&EmailAction{})
}

// Describe describes what an email action does
func (a *EmailAction) Describe() string {
	return "An example action which mimics sending an email"
}

// Execute sends the email
func (a *EmailAction) Execute() {
	log.Printf("New Mail:\nTo: <%s>\nSubject: %s\nMessage: %s\n", a.Receiver, a.Subject, a.Message)
}
