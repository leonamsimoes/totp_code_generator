package main

import (
	"github.com/totp_code_generator/cmd"
	"github.com/totp_code_generator/domain"
	"github.com/totp_code_generator/internal/service"
)

// main is the entry point for the TOTP code generator application.
func main() {
	input, cliErr := service.ReadInput()
	if cliErr.Message != nil {
		service.Print(domain.Response{
			Error: cliErr,
		})
	}

	response := cmd.CreateNewCode(input)
	service.Print(response)

	service.Print(domain.Response{
		PrintOption: domain.EndMessage,
	})
}
