// Status
//
// Provides a HTTP handler for reporting the status of
// the application in JSON format.

package main

import (
	"net/http"

	"github.com/codegangsta/cli"
)

type statusResponse struct {
	Version string `json:"version"`
}

func statusHandler(c *cli.Context) handleFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		response := statusResponse{
			Version: c.App.Version,
		}
		jsonResponse(w, response, http.StatusOK)
	}
}
