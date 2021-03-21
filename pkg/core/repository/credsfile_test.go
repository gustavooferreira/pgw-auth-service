package repository_test

import (
	"testing"

	"github.com/gustavooferreira/pgw-auth-service/pkg/core/repository"
	"github.com/stretchr/testify/assert"
)

func TestExtractParentURL(t *testing.T) {
	tests := map[string]struct {
		username       string
		password       string
		expectedOutput bool
	}{
		"empty user":            {username: "", password: "", expectedOutput: false},
		"user with no password": {username: "bill", password: "", expectedOutput: false},
		"user creds mismatch":   {username: "bill", password: "pass3", expectedOutput: false},
		"no user in db":         {username: "joe", password: "pass123", expectedOutput: false},
		"user/pass match":       {username: "john", password: "pass3", expectedOutput: true},
	}

	ch := createCredsHolder()

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := ch.ValidateUser(test.username, test.password)
			assert.Equal(t, test.expectedOutput, value)
		})
	}
}

func createCredsHolder() *repository.CredsHolder {
	ch := repository.NewCredsHolder()

	ch.Credentials["bill"] = "pass1"
	ch.Credentials["adam"] = "pass2"
	ch.Credentials["john"] = "pass3"

	return &ch
}
