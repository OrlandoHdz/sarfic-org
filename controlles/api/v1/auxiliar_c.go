package v1

import (
	"encoding/json"
	"net/http"
)

//ResponseData structure
type ResponseData struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

//Message retorna un map data
func Message(success bool, message string) map[string]interface{} {
	return map[string]interface{}{"success": success, "message": message}
}

//Respond retorna la estructura
func Respond(w http.ResponseWriter, statuscode int, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(data)
}
