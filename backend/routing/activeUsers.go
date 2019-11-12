package routing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/SUMUKHA-PK/OSVI/backend/responses"
)

// ActiveUsers gets all users
func ActiveUsersPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Active Users changed")
	enableCors(&w)
	// Parse the incoming request
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	if err != nil {
		log.Printf("Bad request from client in routcing/activeUser.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newReq ActiveRequest
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in routing/activeUser.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result int
	result = NumActiveUsers
	if newReq.Data == 1 {
		result++
	} else {
		result--
	}

	NumActiveUsers = result

	outD := &responses.ActiveRespose{200, "Request complete",result}
	outJSON, err := json.Marshal(outD)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/activeUsers.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))
}

// ActiveUsersGet does stuff
func ActiveUsersGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Active Users demanded")
	enableCors(&w)
	outD := &responses.ActiveRespose{200, "Request complete",NumActiveUsers}
	outJSON, err := json.Marshal(outD)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/activeUsers.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))
}