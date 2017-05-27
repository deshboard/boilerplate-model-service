package app

import (
	"database/sql"

	"github.com/Sirupsen/logrus"
)

// Service implements the Protocol Buffer RPC server
type Service struct {
	db *sql.DB

	logger logrus.FieldLogger
}

// NewService creates a new service object
func NewService(db *sql.DB, logger logrus.FieldLogger) *Service {
	return &Service{
		db: db,

		logger: logger,
	}
}
