// package commands - Command to show a TOTP code
package commands

import (
	"fmt"

	"github.com/pquerna/otp"
)

func PrintTOTP(k *otp.Key) {
	fmt.Printf("\nCODE: %s\n", k.AccountName())
}

func PrintBegin() {
	fmt.Printf(".:: BEGIN TOTP ::.\n")
}

func PrintEnd() {
	fmt.Printf("\n.:: END TOTP ::.")
}

func PrintError(errorCode int, err error) {
	var msg string
	switch errorCode {
	case 001:
		msg = "invalid inputs"
	case 002:
		msg = "secret is invalid"
	case 003:
		msg = "number of digits is invalid"

	default:
		msg = fmt.Sprintf("unexpected error: %s", err.Error())
	}

	fmt.Print(msg)
}
