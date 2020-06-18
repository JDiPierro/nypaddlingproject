package api

import (
	"github.com/ricoberger/go-vue-starter/pkg/api/response"
	"github.com/sirupsen/logrus"
	"net/http"
)


func (a *API) getApplicationsHandler(w http.ResponseWriter, r *http.Request) {
	applications, err := a.db.GetApplications()
	if err != nil {
		logrus.WithError(err).Error("Error getting Applications")
		response.Errorf(w, r, nil, http.StatusInternalServerError, "Internal Server Error")
	}
	logrus.Info("Responding with ", len(applications), " apps")
	response.Write(w, r, applications)
	return
}
