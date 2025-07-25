package domain

import "github.com/pquerna/otp"

// PrintOption to help to select the message to be printed
type PrintOption int

// Printed messages
const (
	BeginMessage PrintOption = 1 + iota
	EndMessage
	OTPCodeMessage
	ErrorMessage

	InputSecretMessage
)

// Response represents a message or result to be printed or returned.
type Response struct {
	PrintOptions
	OTPKey *otp.Key
	Error  CLIError
}

// PrintOptions list of value to be show
type PrintOptions []PrintOption
