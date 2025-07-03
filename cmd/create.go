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
			OTPKey:      key,
			PrintOption: domain.ErrorMessage,
			Error: domain.CLIError{
				Message: err,
			},
		}

		return response
	}

	response = domain.Response{
		OTPKey:      key,
		PrintOption: domain.OTPCodeMessage,
	}

	service.Print(domain.Response{
		PrintOption: domain.EndMessage,
	})

	return response
}
