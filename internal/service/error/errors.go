// package service - Core business logic, entities, interfaces
package service

import (
	"errors"

	"github.com/totp_code_generator/domain"
)

// ReturnError based on the Code information returns the CLIError
func ReturnError(err error) domain.CLIError {
	if err == nil {
		return domain.CLIError{}
	}

	for iErr, cliError := range domain.ErrorInformation {
		if errors.Is(errors.Unwrap(err), iErr) {
			return cliError
		}
	}

	return domain.CLIError{
		Message: domain.ErrorMessage_Unexpected,
		Code:    domain.ErrorCode_Unexpected,
	}
}
