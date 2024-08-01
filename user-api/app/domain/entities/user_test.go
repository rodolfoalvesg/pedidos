package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewUser tests the NewUser function.
func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "123.456.789-00", "john@g.com", "1133334444")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Name)
	assert.NotEmpty(t, user.CPF)
	assert.NotEmpty(t, user.Email)
}

// TestNewUserInvalid tests the NewUser function with invalid parameters.
func TestNewUserInvalid(t *testing.T) {
	user, err := NewUser("John Doe", "", "john@g.com", "1133334444")
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

// TestUserValidate tests the validate method of the User entities.
func TestUserValidate(t *testing.T) {
	tests := []struct {
		name     string
		user     User
		expected error
	}{
		{
			name: "valid user",
			user: User{
				Name:  "John Doe",
				CPF:   "123.456.789-00",
				Email: "john@g.com",
				Phone: "1133334444",
			},
			expected: nil,
		},
		{
			name: "missing name",
			user: User{
				Name:  "",
				CPF:   "123.456.789-00",
				Email: "john@g.com",
				Phone: "1133334444",
			},
			expected: ErrNameIsRequired,
		},
		{
			name: "missing cpf",
			user: User{
				Name:  "John Doe",
				CPF:   "",
				Email: "john@g.com",
				Phone: "1133334444",
			},
			expected: ErrCPFIsRequired,
		},
		{
			name: "missing email",
			user: User{
				Name:  "John Doe",
				CPF:   "123.456.789-00",
				Email: "",
				Phone: "1133334444",
			},
			expected: ErrEmailIsRequired,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.user.validate()
			assert.Equal(t, tt.expected, err)
		})
	}
}
