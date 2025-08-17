// Command to generate a TOTP code
package service

import (
	"errors"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/totp_code_generator/domain"
)

// GenerateTOTP generates a TOTP key using the provided options.
func GenerateTOTP(input totp.GenerateOpts) (string, error) {
	now := time.Now().UTC()
	key, err := totp.GenerateCodeCustom(
		string(input.Secret),
		now,
		totp.ValidateOpts{
			Period: input.Period,
			Digits: input.Digits,
		})
	if err != nil {
		return "", errors.Join(err, domain.ErrorMessage_Unexpected)
	}

	return key, nil
}
