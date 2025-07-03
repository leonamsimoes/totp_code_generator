// package service - Core business logic, entities, interfaces

package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/totp_code_generator/domain"
)

func TestReturnError(t *testing.T) {
	tests := []struct {
		name     string
		arg      error
		expected domain.CLIError
	}{
		{
			name:     "Unexpected - nil",
			arg:      nil,
			expected: domain.CLIError{},
		},
		{
			name: "Unexpected - new",
			arg:  fmt.Errorf("error error"),
			expected: domain.CLIError{
				Code: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReturnError(tt.arg)
			assert.Equal(t, tt.expected.Code, result.Code, "should the same error code")
		})
	}
}
