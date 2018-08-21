package resp

import (
	"encoding/json"
	"net/http"
)

//RespondJSON makes the response with payload as json format
func RespondJSON(w *http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		(*w).Write([]byte(err.Error()))
		return
	}
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(status)
	(*w).Write([]byte(response))
}

//RespondError makes the error response with payload as json format
func RespondError(w *http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}

//RespondSuccess returns RespondError with given parameters if error is not null. Otherwise returns JSON data of action:"success"
func RespondSuccess(w *http.ResponseWriter, status int, err error, action string) {
	if err != nil {
		RespondError(w, status, err.Error())
		return
	}
	RespondJSON(w, status, map[string]string{action: "success"})
}
