package server

import (
	"encoding/json"
	"net/http"
)

func setJsonHeader(w http.ResponseWriter, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
}

func jsonResponse(w http.ResponseWriter, resp interface{}, httpStatusCode int) {
	setJsonHeader(w, httpStatusCode)
	jsonResp, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}
