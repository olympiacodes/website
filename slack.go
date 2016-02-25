// Slack
//
// Provides a HTTP handler for signing up for a Slack invite
// via a JSON request.

package main

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/cli"
	"github.com/nlopes/slack"
)

type invitationRequest struct {
	Email string `json:"email"`
}

type invitationResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode,omitempty"`
}

func slackInviteRequestHandler(c *cli.Context) handleFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var invitation invitationRequest
		err := decoder.Decode(&invitation)
		if err != nil {
			jsonResponse(w, invitationResponse{
				Message:   "There was an error processing your invitation. Please try again later.",
				ErrorCode: 1,
			}, http.StatusInternalServerError)
			return
		}

		err = inviteToSlack(c.String("slack-token"), c.String("slack-team"), "", "", invitation.Email)
		if err != nil {
			jsonResponse(w, invitationResponse{
				Message:   "There was an error processing your invitation. Please try again later.",
				ErrorCode: 2,
			}, http.StatusInternalServerError)
			return
		}

		jsonResponse(w, invitationResponse{
			Message: "Your invitation has been sent. Please check your email.",
		}, http.StatusOK)
	}
}

func inviteToSlack(slackToken, teamName, firstName, lastName, email string) error {
	client := slack.New(slackToken)
	return client.InviteToTeam(teamName, firstName, lastName, email)
}
