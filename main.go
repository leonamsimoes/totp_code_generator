package main

import (
	"flag"
	"os"

	"github.com/totp_code_generator/cmd"
	"github.com/totp_code_generator/domain"
	"github.com/totp_code_generator/internal/service"
)

const (
	defaultLength   int    = 6
	defaultDuration int    = 30
	defaultIssuer   string = "gmail"
)

var defaultStatus int

func init() {
	service.Print(domain.Response{
		PrintOptions: []domain.Code{
			domain.LineSeparator,
			domain.LineBreaker,
			domain.LineSeparator,
			domain.LineBreaker,
			domain.BeginMessage,
			domain.LineBreaker,
		},
	})
}

// main is the entry point for the TOTP code generator application.
func main() {
	var (
		status   = defaultStatus
		response = domain.Response{}
	)

	defer func() {
		response.PrintOptions = append(
			response.PrintOptions,
			domain.LineBreaker,
			domain.EndMessage,
			domain.LineBreaker,
			domain.LineSeparator,
			domain.LineBreaker,
			domain.LineSeparator,
			domain.LineBreaker,
		)

		service.Print(response)
		os.Exit(status)
	}()

	inputs := readInput()

	response = inputs.ValidateInputs()
	if response.Error != nil {
		status = response.PrintOptions[len(response.PrintOptions)+1].Int()
		return
	}

	response = cmd.CreateNewCode(inputs)
}

func readInput() domain.Input {
	flagHelper := flag.String(domain.HelperFlag, "", domain.HelperInformation)

	flagIssuer := flag.String(domain.IssuerFlag, defaultIssuer, domain.IssuerInformation)

	flagAccount := flag.String(domain.AccountFlag, "", domain.AccountInformation)

	flagLength := flag.Int(domain.LengthFlag, defaultLength, domain.LengthInformation)

	flagDuration := flag.Int(domain.DurationFlag, defaultDuration, domain.DurationInformation)

	flagSecret := flag.String(domain.SecretFlag, "", domain.SecretInformation)

	flag.Parse()

	return service.NewInput(
		*flagLength,
		*flagDuration,
		*flagHelper,
		*flagIssuer,
		*flagAccount,
		*flagSecret,
	)
}
