package app

import (
	"database/sql"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
	"github.com/goph/stdlib/errors"
)

// Service implements the RPC server.
type Service struct {
	db *sql.DB

	Logger       log.Logger
	ErrorHandler errors.Handler
}

// NewService creates a new service object.
func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,

		Logger:       log.NewNopLogger(),
		ErrorHandler: emperror.NewNullHandler(),
	}
}
