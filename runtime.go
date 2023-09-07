package runtime

import (
	// Internal packages
	"github.com/Hammad887/chat-app/db"
	"github.com/Hammad887/chat-app/db/mysql"
	"github.com/Hammad887/chat-app/service"

	// External package
	wraperrors "github.com/pkg/errors"
)

// Runtime initializes values for entry point to our application
type Runtime struct {
	dbc     db.DataStore
	service service.Manager
}

// NewRuntime creates a new runtime
func NewRuntime() (*Runtime, error) {
	client, err := mysql.NewClient(db.Option{})
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to connect with database")
	}
	return &Runtime{service: service.NewService(&client)}, nil
}

// Service return  service layer object
func (rt *Runtime) Service() service.Manager {
	return rt.service
}
