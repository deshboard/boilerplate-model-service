package app

import (
	"database/sql"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
)

// Service implements the Protocol Buffer RPC server
type Service struct {
	db *sql.DB

	logger       log.Logger
	errorHandler emperror.Handler
}

// NewService creates a new service object
func NewService(db *sql.DB, logger log.Logger, errorHandler emperror.Handler) *Service {
	return &Service{
		db: db,

		logger:       logger,
		errorHandler: errorHandler,
	}
}
