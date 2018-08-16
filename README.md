# bellingham.codes [![Go Report Card](https://goreportcard.com/badge/github.com/bellinghamcodes/website)](https://goreportcard.com/report/github.com/bellinghamcodes/website) [![Build Status](https://travis-ci.org/bellinghamcodes/website.svg?branch=master)](https://travis-ci.org/bellinghamcodes/website)

bellingham.codes is a social group dedicated to growing the local developer community in the greater Bellingham area.

We are committed to providing a friendly, safe and welcoming environment for experienced and aspiring technologists, regardless of gender, gender identity and expression, sexual orientation, disability, physical appearance, body size, race, age or religion. See our [Community Code of Conduct][coc].

## About this Site

This primary purpose of this site is to make it simple for users to join the bellingham.codes Slack workspace and mailing list. In addition to this, the site also lists upcoming events (from meetup.com), recent photos from Instagram, and provides links to our social media channels.

## Building and Running

The website is built using [Go][go]. However, to simplify the contribution and development process for non-Go developers the site can be developed without a local Go installation. Instead the only requirement is to have a working installation of [Docker][docker]. Once you have Docker installed you can use the following commands for
development:

Install dependencies by running:

```sh
make dep
```

To run for development purposes run:

```sh
make dev
```

To create the code-generated files (for a production build and prior to making
a commit) run:

```sh
make generate
```

To run in production mode (always test in production mode prior to making
commits and pull requests!) run:

```sh
make run
```

## Running with Docker

To build the docker image run:

```sh
make docker
```

The docker image generated will expose port 80 running the website.

Additional configuration options are controlled through the following environment variables:

| Environment Variable              | Description                                                                              | Default Value                       |
| --------------------------------- | ---------------------------------------------------------------------------------------- | ----------------------------------- |
| `$ORGANIZATION_NAME`              | Name of the organization                                                                 | `"bellingham.codes"`                |
| `$TWITTER_USERNAME`               | Twitter user to link to in site footer.                                                  | `""`                                |
| `$INSTAGRAM_USERNAME`             | Instagram user to link to in site footer.                                                | `""`                                |
| `$FACEBOOK_PAGE`                  | Facebook page to link to in site footer.                                                 | `""`                                |
| `$SLACK_TEAM`                     | Slack team name, as found in the slack URL.                                              | `""`                                |
| `$SLACK_TOKEN`                    | Access token for your slack team. It can be generated at https://api.slack.com/web#auth. | `""`                                |
| `$MAILCHIMP_TOKEN`                | The API token for your MailChimp account.                                                | `""`                                |
| `$MAILCHIMP_LIST`                 | The ID of the MailChimp list.                                                            | `""`                                |
| `$MEETUP_NAME`                    | Meetup.com group URL name to fetch upcoming events from.                                 | `""`                                |
| `$MEETUP_FETCH_INTERVAL`          | Interval, in minutes, to fetch upcoming event information from Meetup.com                | `30`                                |
| `$CODE_OF_CONDUCT_GITHUB_REPO`    | Github repository to fetch the Community Code of Conduct from                            | `"bellinghamcodes/code-of-conduct"` |
| `$CODE_OF_CONDUCT_FETCH_INTERVAL` | Interval, in minutes, to fetch Community Code of Conduct from Github                     | `30`                                |

For example:

```sh
docker run \
    --detach \
    --publish 8888:80 \
    --env "SLACK_TEAM=bellinghamcodes" \
    --env "SLACK_TOKEN=XXXX-XXXXXXXXXXX-XXXXXXXXXXX-XXXXXXXXXXX-XXXXXXXXXX" \
    --env "MAILCHIMP_TOKEN=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX-us1" \
    --env "MAILCHIMP_LIST=XXXXXXXXXX" \
    --env "ORGANIZATION_NAME=bellingham.codes"
    --env "TWITTER_USERNAME=bellinghamcodes" \
    --env "INSTAGRAM_USERNAME=bellinghamcodes" \
    --env "FACEBOOK_PAGE=bellinghamcodes" \
    tantalic/bellinghamcodes-website:1.10.0
```

## Running in Production

The production site runs on a [Kubernetes][k8s] cluster. To deploy on Kubernetes copy the `kubernetes/secrets.example.yaml` to `kubernetes/secrets.yaml` and complete the required values. Then run:

```sh
cd kubernetes/
./apply.sh
```

[coc]: http://bellingham.codes/code-of-conduct
[go]: http://www.golang.org
[docker]: https://www.docker.com/products/docker-desktop
[k8s]: http://kubernetes.io
