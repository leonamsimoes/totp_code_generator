package domain

import (
	"strings"
)

// Input related with the available flags
type Input struct {
	FlagHelper   string
	FlagIssuer   string
	FlagAccount  string
	FlagSecret   string
	FlagLength   int
	FlagDuration int
}

// ValidateInputs reads user input from stdin and validates it.
func (i Input) ValidateInputs() (resp Response) {
	err := validateString(i.FlagAccount)
	if err != nil {
		resp.PrintOptions = append(resp.PrintOptions, errorCode_InvalidAccount_ExternalError)
	}

	err = validateString(i.FlagIssuer)
	if err != nil {
		resp.PrintOptions = append(resp.PrintOptions, errorCode_InvalidIssuer_ExternalError)
	}

	err = validateLength(i.FlagLength)
	if err != nil {
		resp.PrintOptions = append(resp.PrintOptions, errorCode_InvalidSecretLength_ExternalError)
	}

	resp.Error = err

	return
}

func validateString(str string) error {
	if isEmpty(str) {
		return ErrorMessage_InvalidArguments
	}

	return nil
}

func validateLength(length int) error {
	if length == 0 {
		return ErrorMessage_InvalidArguments
	}

	return nil
}

// isEmpty trims the string and checks if it is empty.
func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
