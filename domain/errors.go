// package domain - Core business logic, entities, interfaces
package domain

const (
	// ErrorCode list of errors
	ErrorCode_Unexpected Code = iota
	ErrorCode_InvalidArguments
	ErrorCode_InvalidSecret
	ErrorCode_InvalidNDigits

	errorMessage_Unexpected       = "something unexpected happens"
	errorMessage_InvalidArguments = "inputs are invalid"
	errorMessage_InvalidSecret    = "secret is invalid"
	errorMessage_InvalidNDigits   = "number of digits is invalid"
)

var errorInformation = map[Code]string{
	ErrorCode_Unexpected:       errorMessage_Unexpected,
	ErrorCode_InvalidArguments: errorMessage_InvalidArguments,
	ErrorCode_InvalidSecret:    errorMessage_InvalidSecret,
	ErrorCode_InvalidNDigits:   errorMessage_InvalidNDigits,
}

type (
	// Code help to define the specific error
	Code int

	// CLIError define the complete error information
	CLIError struct {
		Message string
		Code    Code
	}
)

// ReturnError based on the Code information returns the CLIError
func ReturnError(cd Code) CLIError {
	msg, ok := errorInformation[cd]
	if !ok {
		return CLIError{
			Message: errorInformation[ErrorCode_Unexpected],
			Code:    ErrorCode_Unexpected,
		}
	}

	return CLIError{
		Message: msg,
		Code:    cd,
	}
}
