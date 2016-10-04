package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
)

type handleFunc func(w http.ResponseWriter, req *http.Request)

func main() {
	app := cli.NewApp()
	app.Version = "1.1.1"
	app.Usage = "hack bellingham website"

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
	}

	app.Action = serve

	app.Run(os.Args)
}

func serve(c *cli.Context) {
	http.HandleFunc("/request-invite", inviteRequestHandler(c))
	http.HandleFunc("/status", statusHandler(c))
	http.Handle("/", http.FileServer(assetFS()))

	addr := fmt.Sprintf("%s:%d", c.String("host"), c.Int("port"))
	http.ListenAndServe(addr, nil)
}
