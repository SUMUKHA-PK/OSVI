// This file controls the start of the experiment.
// The API end point hit by the frontend to trigger
// any command on the RT experiment exists here.

package routing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/SUMUKHA-PK/OSVI/backend/responses"
)

// Trigger sends the command intended by
// the user of the website to the RT machine.
func Trigger(w http.ResponseWriter, r *http.Request) {
	log.Println("Relay request initiated.")

	// enableCors(&w)
	// Parse the incoming request
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	if err != nil {
		log.Printf("Bad request from client in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(string(body))
	var newReq TriggerRequest
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(ServerData.ConnectionMap)
	fmt.Println(ServerData.Count)
	// If there is a trigger(er) in the system, the new
	// client must only be able to view and not trigger.
	if ServerData.Count > 1 {
		log.Printf("Triggering client exists, current client can only view\n")
		http.Error(w, "Experiment is already triggered, can only be viewed now\n", http.StatusConflict)
		return
	}

	// Forward the request to the remote RPi @ url
	outData := &responses.TriggerRPiRequest{newReq.RequestType, ":55555"}
	payload, err := json.Marshal(outData)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Sending " + string(payload))

	// Crafting the request to send to the RPi
	req, err := http.NewRequest("POST", URL, strings.NewReader(string(payload)))
	if err != nil {
		log.Printf("Bad request in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	log.Println("Sent data to server")

	// Reading the response from the RPi
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Bad response in routing/startExp.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer res.Body.Close()

	// Reject bad transactions
	if res.StatusCode != 200 {
		log.Printf("Bad request from exp in routing/startExp.go. Wanted 200, received : %v\n", res.StatusCode)
		http.Error(w, "Bad request from exp in routing/startExp.go. Wanted 200, received : "+string(res.StatusCode), http.StatusBadRequest)
		return
	}

	log.Println(res.Status)
	log.Println(res.StatusCode)
	// body, err = ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Printf("Bad request in routing/RPiForward.go")
	// 	log.Println(err)
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	log.Println(ServerData.ConnectionMap)

	outD := &responses.TriggerResponse{200, res.Status}
	outJSON, err := json.Marshal(outD)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
