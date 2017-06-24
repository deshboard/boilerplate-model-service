package app

import (
	"database/sql"

	"github.com/go-kit/kit/log"
)

// Service implements the Protocol Buffer RPC server
type Service struct {
	db *sql.DB

	logger log.Logger
}

// NewService creates a new service object
func NewService(db *sql.DB, logger log.Logger) *Service {
	return &Service{
		db: db,

		logger: logger,
	}
}
