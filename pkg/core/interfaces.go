package core

import "context"

type Repository interface {
	ValidateUser(username string, password string) (valid bool)
}

// ShutDowner represents anything that can be shutdown like an HTTP server.
type ShutDowner interface {
	ShutDown(ctx context.Context) error
}
