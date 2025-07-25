package cmd

import (
	"github.com/pquerna/otp/totp"
	"github.com/totp_code_generator/domain"
	"github.com/totp_code_generator/internal/service"
)

// CreateNewCode creates a new TOTP code and returns a response.
func CreateNewCode(input totp.GenerateOpts) (response domain.Response) {
	key, err := service.GenerateTOTP(input)
	if err != nil {
		response = domain.Response{
			OTPKey: key,
			PrintOptions: domain.PrintOptions{
				domain.ErrorMessage,
			},
			Error: domain.CLIError{
				Message: err,
			},
		}

		return response
	}

	service.Print(domain.Response{
		OTPKey: key,
		PrintOptions: domain.PrintOptions{
			domain.OTPCodeMessage,
			domain.EndMessage,
		},
	})

	return response
}
