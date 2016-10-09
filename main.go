package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bellinghamcodes/website/internal/meetup"
	"github.com/codegangsta/cli"
)

type handleFunc func(w http.ResponseWriter, req *http.Request)

func main() {
	app := cli.NewApp()
	app.Version = "1.3.0"
	app.Usage = "bellingam.codes website"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port",
			Value:  3000,
			EnvVar: "PORT",
			Usage:  "tcp port to listen on",
		},
		cli.StringFlag{
			Name:   "host",
			Value:  "",
			EnvVar: "HOST",
			Usage:  "ip address/host to listen on",
		},
		cli.StringFlag{
			Name:   "slack-team",
			EnvVar: "SLACK_TEAM",
			Usage:  "slack team name, as found in the slack URL",
		},
		cli.StringFlag{
			Name:   "slack-token",
			EnvVar: "SLACK_TOKEN",
			Usage:  "access token for your slack team",
		},
		cli.StringFlag{
			Name:   "mailchimp-token",
			EnvVar: "MAILCHIMP_TOKEN",
			Usage:  "api token for your mailchimp account",
		},
		cli.StringFlag{
			Name:   "mailchimp-list",
			EnvVar: "MAILCHIMP_LIST",
			Usage:  "id of the mailchimp list",
		},
		cli.StringFlag{
			Name:   "meetup",
			EnvVar: "MEETUP_NAME",
			Usage:  "url name of the meetup.com group",
		},
		cli.IntFlag{
			Name:   "meetup-fetch-interval",
			EnvVar: "MEETUP_FETCH_INTERVAL",
			Usage:  "fetch interval in minutes for meetup event information",
			Value:  30,
		},
		cli.StringFlag{
			Name:   "organization-name",
			Value:  "bellingham.codes",
			EnvVar: "ORGANIZATION_NAME",
			Usage:  "name of the organization",
		},
		cli.StringFlag{
			Name:   "twitter",
			EnvVar: "TWITTER_USERNAME",
			Usage:  "twitter account to link to in footer",
		},
		cli.StringFlag{
			Name:   "instagram",
			EnvVar: "INSTAGRAM_USERNAME",
			Usage:  "instagram account to link to in footer",
		},
		cli.StringFlag{
			Name:   "facebook",
			EnvVar: "FACEBOOK_PAGE",
			Usage:  "facebook page to link to in footer",
		},
	}

	app.Action = run

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	groupName := c.String("meetup")
	interval := time.Duration(c.Int("meetup-fetch-interval"))
	homePageServer := &HomePageServer{
		GroupName:         c.String("organization-name"),
		TwitterUsername:   c.String("twitter"),
		InstagramUsername: c.String("instagram"),
		FacebookPage:      c.String("facebook"),
	}

	eventsChan := make(chan []meetup.Event)
	go meetupLoop(groupName, eventsChan, time.Minute*interval)
	go func() {
		for {
			homePageServer.Events = <-eventsChan
		}
	}()

	return serve(c, homePageServer)
}

func serve(c *cli.Context, homePageServer http.Handler) error {
	http.HandleFunc("/request-invite", inviteRequestHandler(c))
	http.HandleFunc("/status", statusHandler(c))
	http.Handle("/", homePageServer)

	addr := fmt.Sprintf("%s:%d", c.String("host"), c.Int("port"))
	log.Printf("Starting server on %s\n", addr)
	return http.ListenAndServe(addr, nil)
}

func meetupLoop(groupName string, c chan []meetup.Event, interval time.Duration) {
	client := meetup.Client{
		GroupURLName: groupName,
	}

	for {
		events, err := client.FetchEvents()
		if err != nil {
			log.Printf("Meetup.com fetch error: %s\n", err)
			goto SLEEP
		}

		log.Printf("Successfully fetched %d events from Meetup.com", len(events))
		c <- events

	SLEEP:
		time.Sleep(interval)
	}
}
