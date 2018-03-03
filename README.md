# bellingham.codes [![Go Report Card](https://goreportcard.com/badge/github.com/bellinghamcodes/website)](https://goreportcard.com/report/github.com/bellinghamcodes/website) [![Build Status](https://travis-ci.org/bellinghamcodes/website.svg?branch=master)](https://travis-ci.org/bellinghamcodes/website)

## About bellingham.codes
bellingham.codes is a social group dedicated to growing the local developer community.

We are committed to providing a friendly, safe and welcoming environment for experienced and aspiring technologists, regardless of gender, gender identity and expression, sexual orientation, disability, physical appearance, body size, race, age or religion.

## About 
The primary purpose of this site is to automate the process for members to join the bellingham.codes Slack team.

## Building and Running
The website is built using [Go][go] with dependencies managed through the [dep][dep] tool. Once you have your go environment setup and dep installed you can get dependencies by running:
```sh
dep ensure
```

Once you have dependencies installed you can build for your current platform by running:
```sh
make build
```

To run for development purposes run:
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

|   Environment Variable   |                                       Description                                        |    Default Value     |
|--------------------------|------------------------------------------------------------------------------------------|----------------------|
| `$ORGANIZATION_NAME`     | Name of the organization                                                                 | `"bellingham.codes"` |
| `$TWITTER_USERNAME`      | Twitter user to link to in site footer.                                                  | `""`                 |
| `$INSTAGRAM_USERNAME`    | Instagram user to link to in site footer.                                                | `""`                 |
| `$FACEBOOK_PAGE`         | Facebook page to link to in site footer.                                                 | `""`                 |
| `$SLACK_TEAM`            | Slack team name, as found in the slack URL.                                              | `""`                 |
| `$SLACK_TOKEN`           | Access token for your slack team. It can be generated at https://api.slack.com/web#auth. | `""`                 |
| `$MAILCHIMP_TOKEN`       | The API token for your MailChimp account.                                                | `""`                 |
| `$MAILCHIMP_LIST`        | The ID of the MailChimp list.                                                            | `""`                 |
| `$MEETUP_NAME`           | Meetup.com group URL name to fetch upcoming events from.                                 | `""`                 |
| `$MEETUP_FETCH_INTERVAL` | Interval, in minutes, to fetch upcoming event information from Meetup.com                | `30`                 |
| `$CODE_OF_CONDUCT_GITHUB_REPO` | Github repository to fetch the Community Code of Conduct from                      | `"bellinghamcodes/code-of-conduct"` |
| `$CODE_OF_CONDUCT_FETCH_INTERVAL` | Interval, in minutes, to fetch Community Code of Conduct from Github            | `30`                                |

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
    tantalic/bellinghamcodes-website:1.6.1
```

## Running in Production
The production site runs on a [Kubernetes][k8s] cluster. To deploy on Kubernetes copy the `kubernetes/secrets.example.yaml` to `kubernetes/secrets.yaml` and complete the required values. Then run:

```sh
cd kubernetes/
./apply.sh
```

[go]: http://www.golang.org
[dep]: https://github.com/golang/dep
[k8s]: http://kubernetes.io