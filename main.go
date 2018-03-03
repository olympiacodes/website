package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bellinghamcodes/website/internal/instagram"
	"github.com/bellinghamcodes/website/internal/meetup"
	"github.com/codegangsta/cli"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type handleFunc func(w http.ResponseWriter, req *http.Request)

var version = "Unknown"

func main() {
	app := cli.NewApp()
	app.Version = version
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
		cli.IntFlag{
			Name:   "instagram-fetch-interval",
			EnvVar: "INSTAGRAM_FETCH_INTERVAL",
			Usage:  "fetch interval in minutes for instagram photos",
			Value:  30,
		},
		cli.IntFlag{
			Name:   "instagram-count",
			EnvVar: "INSTAGRAM_COUNT",
			Usage:  "maximum number of photos to show from instagram",
			Value:  9,
		},
		cli.StringFlag{
			Name:   "facebook",
			EnvVar: "FACEBOOK_PAGE",
			Usage:  "facebook page to link to in footer",
		},
		cli.StringFlag{
			Name:   "coc-github-repo",
			EnvVar: "CODE_OF_CONDUCT_GITHUB_REPO",
			Usage:  "github repository to fetch code of conduct from",
			Value:  "bellinghamcodes/code-of-conduct",
		},
		cli.IntFlag{
			Name:   "coc-fetch-interval",
			EnvVar: "CODE_OF_CONDUCT_FETCH_INTERVAL",
			Usage:  "fetch interval in minutes for code of conduct",
			Value:  90,
		},
	}

	app.Action = run

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	groupName := c.String("meetup")
	instagramUser := c.String("instagram")
	interval := time.Duration(c.Int("meetup-fetch-interval"))
	homePageServer := &HomePageServer{
		GroupName:         c.String("organization-name"),
		TwitterUsername:   c.String("twitter"),
		InstagramUsername: instagramUser,
		FacebookPage:      c.String("facebook"),
		MeetupGroupName:   groupName,
	}

	eventsChan := make(chan []meetup.Event)
	go meetupLoop(groupName, eventsChan, time.Minute*interval)

	imagesChan := make(chan []Image)
	imagesInterval := time.Duration(c.Int("instagram-fetch-interval"))
	max := c.Int("instagram-count")
	go instagramLoop(instagramUser, imagesChan, max, time.Minute*imagesInterval)

	cocServer := &HTMLServer{
		GroupName: c.String("organization-name"),
		Title:     "Community Code of Conduct",
	}
	cocChan := make(chan template.HTML)
	cocInterval := time.Duration(c.Int("coc-fetch-interval"))
	repo := c.String("coc-github-repo")
	go codeOfConductLoop(repo, time.Minute*cocInterval, cocChan)

	go func() {
		for {
			select {
			case events := <-eventsChan:
				homePageServer.Events = events
			case images := <-imagesChan:
				homePageServer.Images = images
			case coc := <-cocChan:
				cocServer.Content = coc
			}
		}
	}()

	return serve(c, homePageServer, cocServer)
}

func serve(c *cli.Context, homePageServer http.Handler, codeOfConductServer http.Handler) error {
	http.HandleFunc("/request-invite", inviteRequestHandler(c))
	http.HandleFunc("/status", statusHandler(c))
	http.Handle("/code-of-conduct", codeOfConductServer)
	http.Handle("/", homePageServer)

	addr := fmt.Sprintf("%s:%d", c.String("host"), c.Int("port"))
	log.Printf("Starting bellingham.codes v%s on %s\n", version, addr)
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

func instagramLoop(username string, c chan []Image, max int, interval time.Duration) {
	client := instagram.Client{}

	for {
		var images []Image
		count := 0
		media, err := client.MediaForUser(username)
		if err != nil {
			log.Printf("Instagram fetch error: %s\n", err)
			goto SLEEP
		}

		for _, m := range media {
			// Skip non-images
			if m.Type != instagram.ImageMediaType {
				continue
			}
			images = append(images, Image{
				Src:  m.Thumbnail,
				Link: "https://www.instagram.com/" + username + "/",
				Alt:  m.Caption,
			})
			count++

			// Limit based on max
			if count >= max {
				break
			}
		}

		log.Printf("Successfully fetched %d images from Instagram", len(images))
		c <- images

	SLEEP:
		time.Sleep(interval)
	}
}

// codeOfConductLoop fetches the content for the Code of Conduct page regularly
// (as defined by the provided interval).
func codeOfConductLoop(repo string, interval time.Duration, c chan template.HTML) {

	client := http.Client{Timeout: time.Second * 15}

	for {
		var data []byte
		resp, err := client.Get("https://raw.githubusercontent.com/" + repo + "/master/README.md")
		if err != nil {
			log.Printf("Error fetching code of conduct: %s\n", err)
			goto SLEEP
		}

		defer resp.Body.Close()
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error fetching code of conduct: %s\n", err)
			goto SLEEP
		}

		log.Printf("Successfully fetched code of conduct (%d bytes) ", len(data))
		data = blackfriday.Run(data, blackfriday.WithExtensions(blackfriday.Autolink|blackfriday.AutoHeadingIDs|blackfriday.HeadingIDs))
		c <- template.HTML(data)

	SLEEP:
		time.Sleep(interval)

	}

}
