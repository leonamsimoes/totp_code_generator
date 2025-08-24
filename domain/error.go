// package domain - Core business logic, entities, interfaces
package domain

import (
	"github.com/pquerna/otp"
)

// Error details
var (
	// FromError maps error values to CLIError details, get the code and return the error.
	FromError = map[Code]error{
		ErrorCode_Unexpected:       ErrorMessage_Unexpected,
		errorCode_InvalidArguments: ErrorMessage_InvalidArguments,
		ErrorCode_InvalidSecret:    ErrorMessage_InvalidSecret,
		errorCode_InvalidNDigits:   errorMessage_InvalidNDigits,

		// externals
		errorCode_InvalidSecret_ExternalError:       otp.ErrValidateSecretInvalidBase32,
		errorCode_InvalidAccount_ExternalError:      otp.ErrGenerateMissingAccountName,
		errorCode_InvalidIssuer_ExternalError:       otp.ErrGenerateMissingIssuer,
		errorCode_InvalidSecretLength_ExternalError: otp.ErrValidateInputInvalidLength,
	}

	// ErrorFrom maps error values to CLIError details, get the error and return the code.
	ErrorFrom = map[error]Code{
		ErrorMessage_Unexpected:       ErrorCode_Unexpected,
		ErrorMessage_InvalidArguments: errorCode_InvalidArguments,
		ErrorMessage_InvalidSecret:    ErrorCode_InvalidSecret,
		errorMessage_InvalidNDigits:   errorCode_InvalidNDigits,

		// externals
		otp.ErrValidateSecretInvalidBase32: errorCode_InvalidSecret_ExternalError,
		otp.ErrGenerateMissingAccountName:  errorCode_InvalidAccount_ExternalError,
		otp.ErrGenerateMissingIssuer:       errorCode_InvalidIssuer_ExternalError,
		otp.ErrValidateInputInvalidLength:  errorCode_InvalidSecretLength_ExternalError,
	}
)
