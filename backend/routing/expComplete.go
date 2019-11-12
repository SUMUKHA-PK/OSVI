package routing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/OSVI/backend/responses"
)

// ExperimentComplete is for the RPi to
// tell that it is free
func ExperimentComplete(w http.ResponseWriter, r *http.Request) {
	log.Println("Experiment complete response received")

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	if err != nil {
		log.Printf("Bad request from client in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newReq ExperimentCompleteResponse
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ServerData.ConnectionMap[newReq.IP]--
	delete(ServerData.ConnectionMap, newReq.IP)
	ServerData.Count--
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
