package main

import (
	"github.com/pquerna/otp/totp"
	"github.com/totp_code_generator/internal/commands"
)

func main() {
	commands.PrintBegin()

	key, err := commands.GenerateTOTP(totp.GenerateOpts{})
	if err != nil {

	}

	commands.PrintTOTP(key)

	commands.PrintEnd()
}
