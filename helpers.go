// Helpers
//
// Misc. helper functions

package main

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, body interface{}, status int) {
	bytes, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)
}
