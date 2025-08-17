// package service - Core business logic, entities, interfaces
package service

import (
	"github.com/totp_code_generator/domain"
)

// ReturnError based on the Code information returns the CLIError
func ReturnError(err error) (c domain.Code) {
	if err != nil {
		c = domain.ErrorFrom[err]
	}

	return c
}
