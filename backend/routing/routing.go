// Package routing has all the routing activities.
// It also has the static IP of the RPi client/server
// that triggers the experiment.
package routing

import (
	"net/http"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"

	"github.com/gorilla/mux"
)

// SetupRouting sets up the routing for the server
func SetupRouting(r *mux.Router) *mux.Router {
	// The static IP of the RPi/RT experiment
	URL = "http://10.100.82.95:12345/"

	// adding all the end points to the router
	r.HandleFunc("/", server.HealthCheckHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/sendRequestToRPi", ForwardToRPi).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/trigger", Trigger).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/experimentComplete", ExperimentComplete).Methods(http.MethodPost, http.MethodOptions)
	return r
}
