// package domain - Core business logic, entities, interfaces

package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnError(t *testing.T) {
	tests := []struct {
		name     string
		arg      Code
		expected CLIError
	}{
		{
			name: errorMessage_Unexpected,
			arg:  0,
			expected: CLIError{
				Message: errorMessage_Unexpected,
				Code:    0,
			},
		},
		{
			name: errorMessage_InvalidArguments,
			arg:  1,
			expected: CLIError{
				Message: errorMessage_InvalidArguments,
				Code:    1,
			},
		},
		{
			name: errorMessage_InvalidSecret,
			arg:  2,
			expected: CLIError{
				Message: errorMessage_InvalidSecret,
				Code:    2,
			},
		},
		{
			name: errorMessage_InvalidNDigits,
			arg:  3,
			expected: CLIError{
				Message: errorMessage_InvalidNDigits,
				Code:    3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReturnError(tt.arg)
			assert.Equal(t, tt.expected, result, "should tha the same CLIError")
		})
	}
}
