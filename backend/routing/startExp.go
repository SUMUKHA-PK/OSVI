// This file controls the start of the experiment.
// The API end point hit by the frontend to trigger
// any command on the RT experiment exists here.

package routing

import (
	"strconv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/SUMUKHA-PK/OSVI/backend/responses"
)

// Trigger sends the command intended by
// the user of the website to the RT machine.
func Trigger(w http.ResponseWriter, r *http.Request) {
	log.Println("Relay request initiated.")
	ServerData.ConnectionMap[r.RemoteAddr]++
	ServerData.Count++
	enableCors(&w)
	// Parse the incoming request
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	if err != nil {
		log.Printf("Bad request from client in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newReq TriggerRequest
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set a client time out of 30s per video
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(newReq.CameraIP)
	if err != nil {
		log.Printf("Couldn't connect to camera in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Create the file
	absPath, _ := os.UserHomeDir()
	out, err := os.Create(absPath + "/OSVIVideos/" + strconv.FormatInt(time.Now().UnixNano(),10)+".mp4")
	if err != nil {
		log.Printf("Couldn't create file in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	
	// If there is a trigger(er) in the system, the new
	// client must only be able to view and not trigger.
	if ServerData.Count > 1 {
		log.Printf("Triggering client exists, current client can only view\n")
		http.Error(w, "Experiment is already triggered, can only be viewed now\n", http.StatusConflict)
		return
	}

	// Forward the request to the remote RPi @ url
	outData := &responses.TriggerRPiRequest{newReq.RequestType, r.RemoteAddr}
	payload, err := json.Marshal(outData)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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


// TriggerGet does stuff
func TriggerGet(w http.ResponseWriter, r *http.Request){

	flags := 0
	if ServerData.Count > 0{
		flags = 0
	}else{
		flags = 1
	}
	fmt.Println("Active Users demanded")
	enableCors(&w)
	outD := &responses.ActiveRespose{200, "Request complete",flags}
	outJSON, err := json.Marshal(outD)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/activeUsers.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))
}	