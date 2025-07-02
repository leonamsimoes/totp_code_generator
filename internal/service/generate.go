// Command to generate a TOTP code
package service

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateTOTP(input totp.GenerateOpts) (*otp.Key, error) {
	key, err := totp.Generate(input)
	if err != nil {

		return &otp.Key{}, err
	}

	return key, nil
}
