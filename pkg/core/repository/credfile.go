package repository

import (
	"gopkg.in/yaml.v2"
)

// CredsHolder holds the credentials of merchants.
// This struct mimics a database.
type CredsHolder struct {
	Credentials map[string]string `yaml:"credentials"`
}

// NewCredsHolder creates a new CredsHolder.
func NewCredsHolder() CredsHolder {
	ch := CredsHolder{Credentials: make(map[string]string)}
	return ch
}

// Load loads data into the CredsHolder.
func (ch *CredsHolder) Load(data []byte) error {
	err := yaml.Unmarshal([]byte(data), &ch)
	return err
}

// ValidateUser checks whether username and password supplied match any in the store.
func (ch *CredsHolder) ValidateUser(username string, password string) (valid bool) {
	if pw, ok := ch.Credentials[username]; ok {
		if password == pw {
			return true
		}
	}
	return false
}
