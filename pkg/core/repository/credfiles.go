package repository

import (
	"gopkg.in/yaml.v2"
)

// CredsHolder holds the credentials of merchants.
// This struct mimics a database.
type CredsHolder struct {
	Credentials map[string]string `yaml:"credentials"`
}

func NewCredsHolder() CredsHolder {
	return CredsHolder{}
}

func (ch *CredsHolder) Load(data []byte) error {
	err := yaml.Unmarshal([]byte(data), &ch)
	if err != nil {
		return err
	}
	return nil
}

func (ch *CredsHolder) ValidateUser(username string, password string) (valid bool) {
	if pw, ok := ch.Credentials[username]; ok {
		if password == pw {
			return true
		}
	}
	return false
}
