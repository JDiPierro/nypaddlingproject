package api

import (
	"github.com/gorilla/mux"
	"github.com/ricoberger/go-vue-starter/pkg/api/response"
	"github.com/ricoberger/go-vue-starter/pkg/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (a *API) getDeploymentsHandler(w http.ResponseWriter, r *http.Request) {
	deploys, err := a.db.GetDeployments()
	if err != nil {
		logrus.WithError(err).Error("Error getting deployments")
		response.Errorf(w, r, nil, http.StatusInternalServerError, "Internal Server Error")
	}
	response.Write(w, r, deploys)
	return
}

func (a *API) mostRecentDeployForAppHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["app_id"]
	envID := vars["env_id"]

	deploy, err := a.db.GetLatestDeployForAppEnv(appID, envID)
	if err != nil {
		logrus.WithError(err).Error("Error getting latest deploy for app/env")
		response.Errorf(w, r, nil, http.StatusInternalServerError, "Internal Server Error")
	}
	response.Write(w, r, deploy)
	return
}

type EnvironmentOverviewResponse struct {
	Applications []*model.Application `json:"applications"`
	Deployments  []*model.Deployment  `json:"deployments"`
}

func (a *API) environmentOverview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	envID := vars["env_id"]

	// Get all deploys by environment
	deploys, err := a.db.GetDeploymentsForEnv(envID)
	if err != nil {
		logrus.WithError(err).Error("Error getting deploys for env")
		response.Errorf(w, r, nil, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Get all app_ids from all loaded deploys
	appIdMap := make(map[string]bool)
	for _, deploy := range deploys {
		appIdMap[deploy.AppID] = true
	}
	appIDs := make([]string, 0)
	for appID := range appIdMap {
		appIDs = append(appIDs, appID)
	}
	if len(appIDs) == 0 {
		logrus.WithError(err).Error("No AppIDs found for env")
		response.Errorf(w, r, nil, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Load Applications
	logrus.Debug("About to bulk get ", len(appIDs), " apps")
	apps, err := a.db.BulkGetApplications(appIDs)
	if err != nil {
		logrus.WithError(err).Error("Error bulk loading apps")
		response.Errorf(w, r, nil, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	logrus.Debug("Loaded ", len(apps), " apps")

	// Respond
	overview := &EnvironmentOverviewResponse{
		Applications: apps,
		Deployments:  deploys,
	}
	response.Write(w, r, overview)
	return
}
