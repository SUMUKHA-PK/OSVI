package routing

// URL is the IP address of the RPi
// in the local network.
var URL string

// TriggerRequest is the request format
// of the Trigger function
type TriggerRequest struct {
	RequestType string
}
