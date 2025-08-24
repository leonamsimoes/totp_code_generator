// package service - Command to show a TOTP code
package service

import (
	"fmt"
	"log"

	"github.com/totp_code_generator/domain"
)

// Print shows the message in the console based on the response.
func Print(resp domain.Response) {
	for _, option := range resp.PrintOptions {
		msg, exist := domain.Messages[option]
		if !exist {
			err, exist := domain.FromError[option]
			if !exist {
				log.Fatal(domain.ErrorMessage_Unexpected.Error())
			}
			msg = err.Error()
		}

		if option == domain.OTPCodeResponse {
			printResponse(resp)
			return
		}

		fmt.Print(msg)
	}
}

func printResponse(resp domain.Response) {
	fmt.Print(resp.OTPKey)
}
