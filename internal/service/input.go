package service

import (
	"github.com/totp_code_generator/domain"
)

// NewInput returns the inputs flags set
func NewInput(flagLength, flagDuration int, flagHelper, flagIssuer, flagAccount, flagSecret string) domain.Input {
	return domain.Input{
		FlagHelper:   flagHelper,
		FlagIssuer:   flagIssuer,
		FlagAccount:  flagAccount,
		FlagSecret:   flagSecret,
		FlagLength:   flagLength,
		FlagDuration: flagDuration,
	}
}
