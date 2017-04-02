package app

import (
	"github.com/Sirupsen/logrus"
	"github.com/deshboard/boilerplate-model-service/apis/boilerplate"
	"github.com/jmoiron/sqlx"
	context "golang.org/x/net/context"
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

// Method is supposed to do something
func (s *Service) Method(ctx context.Context, r *boilerplate.BoilerplateRequest) (*boilerplate.BoilerplateResponse, error) {
	return &boilerplate.BoilerplateResponse{}, nil
}

// StreamingMethod is supposed to do something else
func (s *Service) StreamingMethod(r *boilerplate.BoilerplateRequest, stream boilerplate.Boilerplate_StreamingMethodServer) error {
	for {
		if err := stream.Send(&boilerplate.BoilerplateResponse{}); err != nil {
			return err
		}
	}
}
