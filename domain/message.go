package domain

import (
	"errors"
	"fmt"
)

// ErrorMessages
var (
	ErrorMessage_Unexpected       = errors.New("something unexpected happens")
	ErrorMessage_InvalidSecret    = errors.New("secret is invalid")
	ErrorMessage_InvalidArguments = errors.New("inputs are invalid")
	errorMessage_InvalidNDigits   = errors.New("number of digits is invalid")
)

// Messages Simple message
var (
	Messages = map[Code]string{
		BeginMessage:        "BEGIN",
		EndMessage:          "END",
		OTPCodeMessage:      "Code: ",
		OTPCodeResponse:     "",
		InputSecretMessage:  "insert the secret",
		InputAccountMessage: "insert the account",
		LineSeparator:       "============================================",
		LineBreaker:         "\n",
	}
)

// HelperInformation help flag text information
var HelperInformation = fmt.Sprintf(
	"%s%s",
	"--helper (shows the command args \nrequired values: issuer, account. optional: length.)\n",
	"example: --account=foo@bar.com\n--issuer=FOOBAR123\n--length=8\n",
)

// IssuerInformation issuer flag text information
var IssuerInformation = fmt.Sprintf(
	"%s%s%s",
	"name of the issuing organization/company\n",
	"(eg, gmail, google, outlook, yahoo)\n",
	"(default gmail)\n",
)

// AccountInformation account flag text information
var AccountInformation = fmt.Sprintf(
	"%s%s",
	"name of the User's Account\n",
	"(eg, email address)\n",
)

// DurationInformation duration flag text information
var DurationInformation = fmt.Sprintf(
	"%s%s%s",
	"TTL code\n",
	"(eg, 50, 120... in seconds)\n",
	"defaults = 30\n",
)

// LengthInformation length flag text information
var LengthInformation = fmt.Sprintf(
	"%s%s%s",
	"digits to request\n",
	"defaults to 6\n",
	"(eg, 6 or 8 is the most common values)\n",
)

// SecretInformation secret flag text information
var SecretInformation = fmt.Sprintf(
	"%s%s%s",
	"the secret to be used to generate the OTP code\n",
	"eg, ulid code value\n",
	"default: random\n",
)
