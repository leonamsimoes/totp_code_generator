package service

import (
	"bufio"
	"os"
	"strings"

	"github.com/pquerna/otp/totp"
	"github.com/totp_code_generator/domain"
)

// ReadInput reads user input from stdin and validates it, returning TOTP options and a CLIError.
func ReadInput() (totp.GenerateOpts, domain.CLIError) {
	Print(domain.Response{
		PrintOption: domain.BeginMessage,
	})

	Print(domain.Response{
		PrintOption: domain.InputSecretMessage,
	})

	reader := bufio.NewReader(os.Stdin)
	ln, err := reader.ReadString('\n')
	if err != nil {
		return totp.GenerateOpts{}, domain.CLIError{
			Message: err,
		}
	}

	opts, cliErr := validateInputs(ln)
	if cliErr.Message != nil {
		return totp.GenerateOpts{}, cliErr
	}

	return opts, domain.CLIError{}
}

// validateInputs validates the input line and returns TOTP options and a CLIError.
func validateInputs(ln string) (totp.GenerateOpts, domain.CLIError) {
	secret, cliErr := validateSecret(ln)
	if cliErr.Message != nil {
		return totp.GenerateOpts{}, cliErr
	}

	return totp.GenerateOpts{
		Secret: secret,
	}, domain.CLIError{}
}

// validateSecret checks if the secret is empty and returns the secret as bytes and a CLIError.
func validateSecret(secret string) ([]byte, domain.CLIError) {
	newSecret, empty := isEmpty(secret)
	if !empty {
		return nil, domain.CLIError{
			Message: domain.ErrorMessage_InvalidSecret,
			Code:    domain.ErrorCode_InvalidSecret,
		}
	}

	return []byte(newSecret), domain.CLIError{}
}

// isEmpty trims the string and checks if it is empty.
func isEmpty(s string) (string, bool) {
	s = strings.TrimSpace(s)
	return s, s == ""
}
