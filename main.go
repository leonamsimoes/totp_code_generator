package main

import (
	"flag"
	"os"

	"github.com/totp_code_generator/cmd"
	"github.com/totp_code_generator/domain"
	"github.com/totp_code_generator/internal/service"
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

	inputs, err := readInput()
	if err != nil {
		status = response.PrintOptions[len(response.PrintOptions)+1].Int()
		return
	}

	response, err = inputs.ValidateInputs()
	if err != nil {
		status = response.PrintOptions[len(response.PrintOptions)+1].Int()
		return
	}

	response, err = cmd.CreateNewCode(inputs)
	if err != nil {
		status = response.PrintOptions[len(response.PrintOptions)+1].Int()
		return
	}

	service.Print(response)
}

func readInput() (empty domain.Input, err error) {
	flagHelper, err := service.FlagStringAssist(domain.HelperFlag)
	if err != nil {
		return empty, err
	}

	flagIssuer, err := service.FlagStringAssist(domain.IssuerFlag)
	if err != nil {
		return empty, err
	}

	flagAccount, err := service.FlagStringAssist(domain.AccountFlag)
	if err != nil {
		return empty, err
	}

	flagSecret, err := service.FlagStringAssist(domain.SecretFlag)
	if err != nil {
		return empty, err
	}

	flagLength, err := service.FlagIntAssist(domain.LengthFlag)
	if err != nil {
		return empty, err
	}

	flagDuration, err := service.FlagIntAssist(domain.DurationFlag)
	if err != nil {
		return empty, err
	}

	flag.Parse()

	return service.NewInput(
		*flagLength,
		*flagDuration,
		*flagHelper,
		*flagIssuer,
		*flagAccount,
		*flagSecret,
	), nil
}
