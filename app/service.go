package app

import (
	"database/sql"

	"github.com/goph/stdlib/errors"
	"github.com/goph/stdlib/log"
)

// ServiceOption sets options in the Service.
type ServiceOption func(s *Service)

// DB returns a ServiceOption that sets the DB object for the service.
func DB(db *sql.DB) ServiceOption {
	return func(s *Service) {
		s.db = db
	}
}

// Logger returns a ServiceOption that sets the logger for the service.
func Logger(l log.Logger) ServiceOption {
	return func(s *Service) {
		s.logger = l
	}
}

// ErrorHandler returns a ServiceOption that sets the error handler for the service.
func ErrorHandler(l errors.Handler) ServiceOption {
	return func(s *Service) {
		s.errorHandler = l
	}
}

// Service contains the main controller logic.
type Service struct {
	db *sql.DB

	logger       log.Logger
	errorHandler errors.Handler
}

// NewService creates a new service object.
func NewService(opts ...ServiceOption) *Service {
	s := new(Service)

	for _, opt := range opts {
		opt(s)
	}

	// Default logger
	if s.logger == nil {
		s.logger = log.NewNopLogger()
	}

	// Default error handler
	if s.errorHandler == nil {
		s.errorHandler = errors.NewNopHandler()
	}

	return s
}
