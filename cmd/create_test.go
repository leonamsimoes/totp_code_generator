package cmd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/totp_code_generator/domain"
)

func TestCreateNewCode(t *testing.T) {
	tests := []struct {
		name        string
		input       domain.Input
		expected    domain.Response
		newCode     bool
		errExpected assert.ErrorAssertionFunc
	}{
		{
			name: "successful",
			input: domain.Input{
				FlagAccount: "foo@bar.com",
				FlagIssuer:  "FOOBAR12345",
				FlagLength:  6,
			},
			errExpected: assert.NoError,
			newCode:     false,
		},
		{
			name: "successful - new code",
			input: domain.Input{
				FlagAccount:  "foo@bar.com",
				FlagIssuer:   "FOOBAR12345",
				FlagLength:   6,
				FlagDuration: 1,
			},
			errExpected: assert.NoError,
			newCode:     true,
		},
		{
			name:        "fail - without data",
			input:       domain.Input{},
			errExpected: assert.NoError,
			newCode:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := CreateNewCode(tt.input)
			ok := tt.errExpected(t, response.Error, "should be equal")
			if ok {
				return
			}

			assert.NotEmpty(t, response.OTPKey, "OTP key should be equal")

			if tt.newCode {
				time.Sleep(time.Duration(tt.input.FlagDuration + 5))
				newResponse := CreateNewCode(tt.input)
				assert.NotEqualf(t, response, newResponse, "OTP key should not be equal")
			}
		})
	}
}
