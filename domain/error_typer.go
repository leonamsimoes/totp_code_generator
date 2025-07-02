// package domain - Core business logic, entities, interfaces
package domain

import (
	"errors"

	"github.com/pquerna/otp"
)

type (
	// Code help to define the specific error
	Code int

	// CLIError defines the complete error information for the CLI.
	CLIError struct {
		Message error
		Code    Code
	}
)

// ErrorCode list of errors code
const (
	ErrorCode_Unexpected Code = iota
	ErrorCode_InvalidSecret
	errorCode_InvalidArguments
	errorCode_InvalidNDigits
	errorCode_ExternalError
)

// Error details
var (
	ErrorMessage_Unexpected       = errors.New("something unexpected happens")
	ErrorMessage_InvalidSecret    = errors.New("secret is invalid")
	errorMessage_InvalidArguments = errors.New("inputs are invalid")
	errorMessage_InvalidNDigits   = errors.New("number of digits is invalid")
	errorMessage_InvalidAccounts  = errors.New("account is invalid")

	// ErrorInformation maps error values to CLIError details.
	ErrorInformation = map[error]CLIError{
		ErrorMessage_Unexpected: {
			Message: ErrorMessage_Unexpected,
			Code:    ErrorCode_Unexpected,
		},

		errorMessage_InvalidArguments: {
			Message: errorMessage_InvalidArguments,
			Code:    errorCode_InvalidArguments,
		},

		ErrorMessage_InvalidSecret: {
			Message: ErrorMessage_InvalidSecret,
			Code:    ErrorCode_InvalidSecret,
		},

		errorMessage_InvalidNDigits: {
			Message: errorMessage_InvalidNDigits,
			Code:    errorCode_InvalidNDigits,
		},

		// externals
		otp.ErrValidateSecretInvalidBase32: {
			Message: ErrorMessage_InvalidSecret,
			Code:    errorCode_ExternalError,
		},

		otp.ErrValidateInputInvalidLength: {
			Message: ErrorMessage_Unexpected,
			Code:    errorCode_ExternalError,
		},

		otp.ErrGenerateMissingIssuer: {
			Message: errorMessage_InvalidNDigits,
			Code:    errorCode_ExternalError,
		},

		otp.ErrGenerateMissingAccountName: {
			Message: errorMessage_InvalidAccounts,
			Code:    errorCode_ExternalError,
		},
	}
)
