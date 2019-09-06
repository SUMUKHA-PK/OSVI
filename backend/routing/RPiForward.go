package routing

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// ForwardToRPi forwards a basic test request to the RPi
func ForwardToRPi(w http.ResponseWriter, r *http.Request) {

	log.Println("Test forward to RPi")

	url := "http://10.100.82.95:12346/"

	payload := strings.NewReader("{\"Url\": \"" + url + "\"}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Printf("Bad request in routing/RPiForward.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.Header.Add("Prediction-Key", "5b6a7662b4b24116a58d16683e0606b0")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Bad request in routing/RPiForward.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Sent data to server")
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Bad request in routing/RPiForward.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Got data from server")
	fmt.Println(string(body))
}
