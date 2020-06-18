package api

import (
	"encoding/json"
	"github.com/ricoberger/go-vue-starter/pkg/api/response"
	"github.com/ricoberger/go-vue-starter/pkg/api/webhook"
	"github.com/sirupsen/logrus"
	"net/http"
)

/*
Example Request:

curl -v -XPOST -d '{"application": "streaming-stats-service", "env": "live", "version": "dad45c3e57d5de593ef0d9865f3c25710cdd8ef5", "user": "jdipierro"}' localhost:8080/api/v1/webhook
 */

func (a *API) webhookHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request
	payload := &webhook.Request{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		logrus.WithError(err).Error("Error decoding Deployment request")
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = a.services.webhook.HandleWebhook(payload)
	if err != nil {
		logrus.WithError(err).Error("Error handling deployment webhook")
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.Write(w, r, nil)
	return
}
