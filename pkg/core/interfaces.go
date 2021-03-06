package core

import "context"

// Repository represents a database holding credentials
type Repository interface {
	ValidateUser(username string, password string) (valid bool)
}

// ShutDowner represents anything that can be shutdown like an HTTP server.
type ShutDowner interface {
	ShutDown(ctx context.Context) error
}
