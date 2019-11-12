package routing

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/OSVI/backend/responses"
)

// GetViews does
func GetViews(w http.ResponseWriter, r *http.Request) {
	log.Println("Viewers demanded")
	enableCors(&w)
	if ServerData.Count > 1 {
		Flag = -1
	}
	outD := &responses.ViewersRes{200, "Viewer Data", ServerData.Count, Flag}
	outJSON, err := json.Marshal(outD)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))
}
