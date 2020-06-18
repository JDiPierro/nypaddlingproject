package api

import (
	"github.com/ricoberger/go-vue-starter/pkg/api/response"
	"github.com/sirupsen/logrus"
	"net/http"
)


func (a *API) getEnvironmentsHandler(w http.ResponseWriter, r *http.Request) {
	envs, err := a.db.GetEnvironments()
	if err != nil {
		logrus.WithError(err).Error("Error getting environments")
		response.Errorf(w, r, nil, http.StatusInternalServerError, "Internal Server Error")
	}
	response.Write(w, r, envs)
	return
}
