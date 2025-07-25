// package service - Command to show a TOTP code
package service

import (
	"fmt"

	"github.com/pquerna/otp"
	"github.com/totp_code_generator/domain"
)

// Print shows the messages in the console based on the response.
func Print(resps ...domain.Response) {
	for _, resp := range resps {
		for _, option := range resp.PrintOptions {
			var msg string

			switch option {
			case 1:
				msg = printBegin()
			case 2:
				msg = printEnd()
			case 3:
				msg = printTOTP(resp.OTPKey)
			case 4:
				msg = printError(resp.Error)

			default:
				msg = printError(domain.CLIError{
					Message: fmt.Errorf("not expected, %s", resp.Error.Message),
				})
			}

			fmt.Print(msg)
		}
	}
}

func printTOTP(k *otp.Key) string {
	return fmt.Sprintf("\nCODE: %s\n", k.AccountName())
}

func printBegin() string {
	return ".:: BEGIN TOTP ::.\n"
}

func printEnd() string {
	return ".:: END TOTP ::."
}

func printError(err domain.CLIError) (msg string) {
	switch err.Code {
	case 1:
		msg = "invalid inputs"

	case 2:
		msg = "secret is invalid"

	case 3:
		msg = "number of digits is invalid"

	default:
		msg = fmt.Sprintf("internal error: %s", err.Message)
	}

	return msg
}
