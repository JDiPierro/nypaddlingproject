package api

import (
	"github.com/ricoberger/go-vue-starter/pkg/api/webhook"
	"net/http"

	"github.com/ricoberger/go-vue-starter/pkg/api/response"
	"github.com/ricoberger/go-vue-starter/pkg/db"
	"github.com/ricoberger/go-vue-starter/pkg/mail"

	"github.com/gorilla/mux"
)

// Config represents the API configuration
type Config struct {
	Domain        string `yaml:"domain"`
	SigningSecret string `yaml:"signing_secret"`
}

// API represents the structure of the API
type API struct {
	Router *mux.Router

	config *Config
	db     db.DB
	mail   *mail.Client
	services *Services
}

type Services struct {
	webhook *webhook.Service
}

// New returns the api settings
func New(config *Config, db db.DB, mail *mail.Client, router *mux.Router) (*API, error) {
	api := &API{
		config: config,
		db:     db,
		mail:   mail,

		services: &Services{
			webhook.NewService(db),
		},

		Router: router,
	}

	api.InitRouter()
	return api, nil
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	response.Write(w, r, struct {
		Status string `json:"status"`
	}{
		"ok",
	})

	return
}

func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
