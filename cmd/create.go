package cmd

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/totp_code_generator/domain"
	"github.com/totp_code_generator/internal/service"
)

// CreateNewCode creates a new TOTP code and returns a response.
func CreateNewCode(ipt domain.Input) (resp domain.Response) {
	input := totp.GenerateOpts{
		Issuer:      ipt.FlagIssuer,
		AccountName: ipt.FlagAccount,
		Digits:      otp.Digits(ipt.FlagLength),
		Secret:      []byte(ipt.FlagSecret),
		Period:      uint(ipt.FlagDuration),
	}

	key, err := service.GenerateTOTP(input)
	if err != nil {
		resp = domain.Response{
			OTPKey: key,
			PrintOptions: []domain.Code{
				domain.ErrorCode_Unexpected,
			},
			Error: err,
		}

		return resp
	}

	service.Print(domain.Response{
		OTPKey: key,
		PrintOptions: []domain.Code{
			domain.OTPCodeMessage,
		},
	})

	service.Print(domain.Response{
		OTPKey: key,
		PrintOptions: []domain.Code{
			domain.OTPCodeResponse,
		},
	})

	return
}
