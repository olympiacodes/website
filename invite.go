// Invite
//
// Provides a HTTP handler for signing up for a Slack invite and/or
// a MailChimp mailing list signup via a JSON request.

package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/codegangsta/cli"
	"github.com/mnbbrown/mailchimp"
	"github.com/nlopes/slack"
)

type invitationRequest struct {
	Email              string `json:"email"`
	RequestSlack       bool   `json:"requestSlack"`
	RequestMailingList bool   `json:"requestMailingList"`
}

type invitationResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode,omitempty"`
}

func inviteRequestHandler(c *cli.Context) handleFunc {
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

		var (
			slackErr, chimpErr error
			wg                 sync.WaitGroup
		)

		if invitation.RequestSlack {
			wg.Add(1)
			go func() {
				defer wg.Done()
				slackErr = inviteToSlack(c.String("slack-token"), c.String("slack-team"), "", "", invitation.Email)
			}()
		}

		if invitation.RequestMailingList {
			wg.Add(1)
			go func() {
				defer wg.Done()
				chimpErr = subscribeToList(c.String("mailchimp-token"), invitation.Email, c.String("mailchimp-list"))
			}()
		}

		wg.Wait()

		if slackErr != nil || chimpErr != nil {
			var (
				message string
				code    int
			)

			if slackErr != nil && chimpErr != nil {
				message = "There was an error processing your invitation. Please try again later."
				code = 2
			} else if slackErr != nil {
				message = "There was an error processing your invitation to the Slack group. Please try again later."
				code = 3
			} else {
				message = "There was an adding you to the mailing list. Please try again later."
				code = 4
			}

			jsonResponse(w, invitationResponse{
				Message:   message,
				ErrorCode: code,
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

func subscribeToList(mailchimpApiKey, email, listId string) error {
	client, err := mailchimp.NewClient(mailchimpApiKey, nil)
	if err != nil {
		return err
	}

	_, err = client.Subscribe(email, listId)
	return err
}
