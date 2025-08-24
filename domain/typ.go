package domain

// Lines design
const (
	LineSeparator Code = iota - 10
	LineBreaker
)

// Printed messages
const (
	BeginMessage Code = iota
	EndMessage
	OTPCodeMessage
	InputSecretMessage
	InputAccountMessage
	OTPCodeResponse
)

// ErrorCode list of errors code
const (
	ErrorCode_Unexpected Code = iota + 1000
	ErrorCode_InvalidSecret
	errorCode_InvalidArguments
	errorCode_InvalidNDigits
	errorCode_InvalidSecret_ExternalError
	errorCode_InvalidSecretLength_ExternalError
	errorCode_Unexpected_ExternalError
	errorCode_InvalidAccount_ExternalError
	errorCode_InvalidIssuer_ExternalError
)

// Flags flag expected
const (
	HelperFlag   = "help"
	IssuerFlag   = "issuer"
	AccountFlag  = "account"
	LengthFlag   = "length"
	DurationFlag = "duration"
	SecretFlag   = "secret"

	HelperFlagAlias   = "-h"
	IssuerFlagAlias   = "-i"
	AccountFlagAlias  = "-a"
	LengthFlagAlias   = "-l"
	DurationFlagAlias = "-d"
	SecretFlagAlias   = "-s"
)

type (
	// Response represents a message or result to be printed or returned.
	Response struct {
		PrintOptions []Code
		OTPKey       string
	}

	// Code helps to select the message to be printed
	Code int
)

// Int returns the code into the int format
func (c Code) Int() int {
	return int(c)
}
