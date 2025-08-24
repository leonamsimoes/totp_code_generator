package service

import (
	"flag"

	"github.com/totp_code_generator/domain"
)

const (
	defaultLength   int    = 6
	defaultDuration int    = 30
	defaultIssuer   string = "gmail"
)

// NewInput returns the inputs flags set
func NewInput(argLength, argDuration int, argHelper, argIssuer, argAccount, argSecret string) domain.Input {
	return domain.Input{
		FlagHelper:   argHelper,
		FlagIssuer:   argIssuer,
		FlagAccount:  argAccount,
		FlagSecret:   argSecret,
		FlagLength:   argLength,
		FlagDuration: argDuration,
	}
}

// FlagStringAssist helps to expect the whole flag or the alias
func FlagStringAssist(strFlag string) (str *string, err error) {
	var (
		defaultValue string
		msg          string
	)
	switch strFlag {
	case domain.HelperFlag:
		defaultValue = ""
		msg = domain.HelperInformation

	case domain.IssuerFlag:
		defaultValue = defaultIssuer
		msg = domain.IssuerInformation

	case domain.AccountFlag:
		defaultValue = defaultIssuer
		msg = domain.AccountInformation

	case domain.SecretFlag:
		defaultValue = defaultIssuer
		msg = domain.SecretInformation

	default:
		return str, domain.ErrorMessage_InvalidArguments
	}

	return flag.String(strFlag, defaultValue, msg), nil
}

// FlagIntAssist helps to expect the whole flag or the alias
func FlagIntAssist(strFlag string) (i *int, err error) {
	var (
		defaultValue int
		msg          string
	)

	switch strFlag {
	case domain.LengthFlag:
		defaultValue = defaultLength
		msg = domain.HelperInformation

	case domain.DurationFlag:
		defaultValue = defaultDuration
		msg = domain.IssuerInformation

	default:
		return i, domain.ErrorMessage_InvalidArguments
	}

	return flag.Int(strFlag, defaultValue, msg), nil
}
