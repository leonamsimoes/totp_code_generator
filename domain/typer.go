package domain

import "github.com/pquerna/otp"

type PrintOption int

const (
	// Response
	BeginMessage PrintOption = 1 + iota
	EndMessage
	OTPCodeMessage
	ErrorMessage

	// Request
	InputSecretMessage
)

type Response struct {
	PrintOption
	OTPKey *otp.Key
	Error  CLIError
}
