package routing

import (
	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
)

// URL is the IP address of the RPi
// in the local network.
var URL string

// ServerData contains all the data in the server
var ServerData server.Data

//Flag int to maintain first user or last
var Flag int

var NumActiveUsers int

// TriggerRequest is the request format
// of the Trigger function
type TriggerRequest struct {
	RequestType string
	CameraIP    string
}

// ExperimentCompleteResponse struct
type ExperimentCompleteResponse struct {
	IP string
}

type ActiveRequest struct {
	Data int
}
