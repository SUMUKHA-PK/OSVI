package routing

import (
	"net/http"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"

	"github.com/gorilla/mux"
)

// SetupRouting sets up the routing for the server
func SetupRouting(r *mux.Router) *mux.Router {
	r.HandleFunc("/", server.HealthCheckHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/sendRequestToRPi", ForwardToRPi).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/getItemsFromCart", GetItemsFromCart).Methods(http.MethodPost, http.MethodOptions)
	return r
}
