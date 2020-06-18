package webhook

import (
	"github.com/ricoberger/go-vue-starter/pkg/db"
	"github.com/ricoberger/go-vue-starter/pkg/model"
	"time"
)

type Service struct {
	db db.DB
}

type Request struct {
	Application string `json:"application"`
	Env         string `json:"env"`
	Version     string `json:"version"`
	User        string `json:"user"`
	Branch      string `json:"branch"`
	Message     string `json:"message"`
}

func NewService(db db.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) HandleWebhook(req *Request) error {
	// Find or Create App by name
	app, err := s.findOrCreateApplication(req.Application)
	if err != nil {
		return err
	}
	// Find or Create Env by name
	env, err := s.findOrCreateEnvironment(req.Env)
	if err != nil {
		return err
	}
	// Create Deployment
	deploy := &model.Deployment{
		AppID:     app.ID,
		EnvID:     env.ID,
		User:      req.User,
		Version:   req.Version,
		Branch:    req.Branch,
		Message:   req.Message,
		Timestamp: time.Now().Format(time.RFC850),
	}
	err = s.db.CreateDeployment(deploy)

	return err
}

func (s *Service) findOrCreateApplication(appName string) (*model.Application, error) {
	app, err := s.db.GetApplicationByName(appName)
	if err != nil {
		return nil, err
	}
	if app != nil {
		// App found, return it
		return app, nil
	}

	// App not found, build a new one.
	app = &model.Application{
		Name: appName,
	}
	err = s.db.CreateApplication(app)
	return app, err
}

func (s *Service) findOrCreateEnvironment(envName string) (*model.Environment, error) {
	env, err := s.db.GetEnvironmentByName(envName)
	if err != nil {
		return nil, err
	}
	if env != nil {
		// Env found, return it
		return env, nil
	}

	// Env not found, build a new one.
	env = &model.Environment{
		Name: envName,
	}
	err = s.db.CreateEnvironment(env)
	return env, err
}
