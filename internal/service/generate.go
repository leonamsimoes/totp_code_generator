// Command to generate a TOTP code
package service

import (
	"errors"
	"fmt"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// GenerateTOTP generates a TOTP key using the provided options.
func GenerateTOTP(input totp.GenerateOpts) (*otp.Key, error) {
	key, err := totp.Generate(input)
	if err != nil {
		return &otp.Key{}, errors.Join(err, fmt.Errorf("could not generate the TOTP code"))
	}

	return key, nil
}
