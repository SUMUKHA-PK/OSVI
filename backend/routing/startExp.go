// This file controls the start of the experiment.
// The API end point hit by the frontend to trigger
// any command on the RT experiment exists here.

package routing

import (
	"encoding/json"
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

	// Parse the incoming request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/startExp.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newReq TriggerRequest
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Coudn't Unmarshal data in routing/startExp.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Forward the request to the remote RPi @ url

	outData := &responses.TriggerRPiRequest{newReq.RequestType, ":55555"}
	payload, err := json.Marshal(outData)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/groceries.go/AddItemsToCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Sending " + string(payload))

	// Crafting the request to send to the RPi
	req, err := http.NewRequest("POST", URL, strings.NewReader(string(payload)))
	if err != nil {
		log.Printf("Bad request in routing/startExp.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	log.Println("Sent data to server")

	// Reading the response from the RPi
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Bad response in routing/startExp.go: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer res.Body.Close()

	// Reject bad transactions
	if res.StatusCode != 200 {
		log.Printf("Bad request in routing/startExp.go. Wanted 200, received : %v", res.StatusCode)
		// outD =
		http.Error(w, "Bad request in routing/startExp.go. Wanted 200, received : "+string(res.StatusCode), http.StatusBadRequest)
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
		log.Printf("Can't Marshall to JSON in routing/startExp.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))
}
