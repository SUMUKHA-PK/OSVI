package routing

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/OSVI/backend/responses"
)

// ExperimentComplete is for the RPi to
// tell that it is free
func ExperimentComplete(w http.ResponseWriter, r *http.Request) {
	log.Println("Experiment complete response received")
	ServerData.Count = 0
	// Currently the Count works as a flag.
	// We can incorporate the use of the Map
	// later on I guess.

	outD := &responses.HTTPStatusOK{200, "Request complete"}
	outJSON, err := json.Marshal(outD)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))
}
