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

## Running (in Production)
Running the site will require the setting of a number of options. These options can be set via command line flags or environment variables. 

| Environment Variable |         Flag        |                                       Description                                        | Default Value |
|----------------------|---------------------|------------------------------------------------------------------------------------------|---------------|
| `$PORT`              | `--port`            | The TCP port to listen on.                                                               | `3000`        |
| `$HOST`              | `--host`            | The IP address/hostname to listen on.                                                    | All hosts     |
| `$SLACK_TEAM`        | `--slack-team`      | Slack team name, as found in the slack URL.                                              | `""`          |
| `$SLACK_TOKEN`       | `--slack-token`     | Access token for your slack team. It can be generated at https://api.slack.com/web#auth. | `""`          |
| `$MAILCHIMP_TOKEN`   | `--mailchimp-token` | The API token for your MailChimp account.                                                | `""`          |
| `$MAILCHIMP_LIST`    | `--mailchimp-list`  | The ID of the MailChimp list.                                                            | `""`          |

### systemd Configuration
The canonical way to run the site is through the [`systemd`][systemd] service manager to setup the environment, manage when the application is started, and monitor the process to keep it running. This can be done with a system file like the one below:

. The following 

```apacheconf
[Unit]
Description=bellingham.codes Website

[Service]
ExecStart=/usr/local/bin/bellinghamcodes
Restart=always
User=root
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=bellinghamcodes
Environment=PORT=80
Environment=SLACK_TEAM=bellinghamcodes
Environment=SLACK_TOKEN=XXXX-XXXXXXXXXXX-XXXXXXXXXXX-XXXXXXXXXXX-XXXXXXXXXX
Environment=MAILCHIMP_TOKEN=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX-us1
Environment=MAILCHIMP_LIST=XXXXXXXXXX

[Install]
WantedBy=multi-user.target
```


[go]: http://www.golang.org
[glide]: https://glide.sh
[systemd]: https://freedesktop.org/wiki/Software/systemd/
