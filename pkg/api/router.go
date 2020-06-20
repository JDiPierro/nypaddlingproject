package api

import (
	"net/http"
)

func (a *API) InitRouter() {
	// Endpoint for browser preflight requests
	a.Router.Methods("OPTIONS").HandlerFunc(a.corsMiddleware(a.preflightHandler))

	// Endpoint for healthcheck
	a.Router.HandleFunc("/api/v1/health", a.publicMiddleware(a.healthHandler)).Methods("GET")

	// Account related endpoints
	/*a.Router.HandleFunc("/api/v1/auth", a.publicMiddleware(a.userLoginHandler)).Methods("POST")
	a.Router.HandleFunc("/api/v1/account", a.publicMiddleware(a.userSignupHandler)).Methods("POST")
	a.Router.HandleFunc("/api/v1/account", a.authenticatedMiddleware(a.userUpdateProfileHandler)).Methods("PUT")
	a.Router.HandleFunc("/api/v1/account", a.authenticatedMiddleware(a.userProfileHandler)).Methods("GET")
	a.Router.HandleFunc("/api/v1/account/email/{id}/{token}", a.publicMiddleware(a.userVerifyHandler)).Methods("GET")
	a.Router.HandleFunc("/api/v1/account/email", a.publicMiddleware(a.userResendVerificationMail)).Methods("POST")
	a.Router.HandleFunc("/api/v1/account/password", a.publicMiddleware(a.forgotPasswordHandler)).Methods("POST")
	a.Router.HandleFunc("/api/v1/account/password", a.publicMiddleware(a.resetPasswordHandler)).Methods("PUT")*/

}

func (a *API) authenticatedMiddleware(apiHandler http.HandlerFunc) http.HandlerFunc {
	return a.publicMiddleware(apiHandler)
	// Disabling authentication for now...
	//return a.corsMiddleware(a.logMiddleware(a.jwtMiddleware(apiHandler)))
}

func (a *API) publicMiddleware(apiHandler http.HandlerFunc) http.HandlerFunc {
	return a.corsMiddleware(a.logMiddleware(apiHandler))
}
