package app

import (
	"github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
)

// Service implements the Protocol Buffer RPC server
type Service struct {
	db *sqlx.DB

	logger logrus.FieldLogger
}

// NewService creates a new service object
func NewService(db *sqlx.DB, logger logrus.FieldLogger) *Service {
	return &Service{
		db: db,

		logger: logger,
	}
}
