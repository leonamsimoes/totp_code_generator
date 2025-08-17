// package service - Core business logic, entities, interfaces

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/totp_code_generator/domain"
)

func TestReturnError(t *testing.T) {
	tests := []struct {
		name     string
		arg      error
		expected domain.Code
	}{
		{
			name:     "Error - nil",
			arg:      nil,
			expected: 0,
		},
		{
			name:     "Error - Unexpected",
			arg:      domain.ErrorMessage_Unexpected,
			expected: domain.ErrorCode_Unexpected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReturnError(tt.arg)
			assert.Equal(t, tt.expected, result, "should the same error code")
		})
	}
}
