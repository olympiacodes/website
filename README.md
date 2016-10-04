# bellingham.codes [![Code Climate](https://codeclimate.com/github/bellinghamcodes/website/badges/gpa.svg)](https://codeclimate.com/github/bellinghamcodes/website) [![Build Status](https://travis-ci.org/bellinghamcodes/website.svg?branch=master)](https://travis-ci.org/bellinghamcodes/website)

## About bellingham.codes
bellingham.codes is a social group dedicated to growing the local developer community.

We are committed to providing a friendly, safe and welcoming environment for experienced and aspiring technologists, regardless of gender, gender identity and expression, sexual orientation, disability, physical appearance, body size, race, age or religion.

## About 
The primary purpose of this site is to automate the process for members to join the bellingham.codes Slack team.

## Building and Running
The website is built using [Go][go] with dependencies managed through [Glide][glide]. Once you have your go environment setup and glide installed you can get dependencies by running:
```sh
glide install
```

Once you have dependencies installed you can build for your current platform by running:
```sh
make
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

| Environment Variable |                                       Description                                        | Default Value |
|----------------------|------------------------------------------------------------------------------------------|---------------|
| `$SLACK_TEAM`        | Slack team name, as found in the slack URL.                                              | `""`          |
| `$SLACK_TOKEN`       | Access token for your slack team. It can be generated at https://api.slack.com/web#auth. | `""`          |
| `$MAILCHIMP_TOKEN`   | The API token for your MailChimp account.                                                | `""`          |
| `$MAILCHIMP_LIST`    | The ID of the MailChimp list.                                                            | `""`          |

For example:
```sh
docker run \
    --detach \
    --publish 8888:80 \
    --env "SLACK_TEAM=bellinghamcodes" \
    --env "SLACK_TOKEN=XXXX-XXXXXXXXXXX-XXXXXXXXXXX-XXXXXXXXXXX-XXXXXXXXXX" \
    --env "MAILCHIMP_TOKEN=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX-us1" \
    --env "MAILCHIMP_LIST=XXXXXXXXXX" \
    tantalic/bellinghamcodes-website:latest
```

[go]: http://www.golang.org
[glide]: https://glide.sh
