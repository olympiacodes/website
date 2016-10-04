# HackBellingham.com [![Code Climate](https://codeclimate.com/github/hackbellingham/website/badges/gpa.svg)](https://codeclimate.com/github/hackbellingham/website) [![Build Status](https://travis-ci.org/hackbellingham/website.svg?branch=master)](https://travis-ci.org/hackbellingham/website)

## About Hack Bellingham
Hack Bellingham is a social group dedicated to growing the local developer community.

We are committed to providing a friendly, safe and welcoming environment for experienced and aspiring technologists, regardless of gender, gender identity and expression, sexual orientation, disability, physical appearance, body size, race, age or religion.

## About 
The primary purpose of this site is to automate the process for members to join the Hack Bellingham Slack team.

## Building and Running
The website is built using [Go][go] and is 'go gettable'. Once you have your go environment setup you can get dependencies by running:
```sh
go get
```

Once you have dependencies installed you can build for your current platform by running:
```sh
make
```

To run for development purposes run:
```sh
make run
```

To cross-compile to another platform run one of the following:
```sh
make linux
make freebsd
make osx
```

## Running (in Production)
Running the site will require the setting of a number of options. These options can be set via command line flags or environment variables. 

|        Environment Variable        |         Flag        |                                       Description                                        | Default Value |
|------------------------------------|---------------------|------------------------------------------------------------------------------------------|---------------|
| `$HACK_BELLINGHAM_PORT`            | `--port`            | The TCP port to listen on.                                                               | `3000`        |
| `$HACK_BELLINGHAM_HOST`            | `--host`            | The IP address/hostname to listen on.                                                    | All hosts     |
| `$HACK_BELLINGHAM_SLACK_TEAM`      | `--slack-team`      | Slack team name, as found in the slack URL.                                              | `""`          |
| `$HACK_BELLINGHAM_SLACK_TOKEN`     | `--slack-token`     | Access token for your slack team. It can be generated at https://api.slack.com/web#auth. | `""`          |
| `$HACK_BELLINGHAM_MAILCHIMP_TOKEN` | `--mailchimp-token` | The API token for your MailChimp account.                                                | `""`          |
| `$HACK_BELLINGHAM_MAILCHIMP_LIST`  | `--mailchimp-list`  | The ID of the MailChimp list.                                                            | `""`          |

### systemd Configuration
The canonical way to run the site is through the [`systemd`][systemd] service manager to setup the environment, manage when the application is started, and monitor the process to keep it running. This can be done with a system file like the one below:

. The following 

```apacheconf
[Unit]
Description=Hack Bellingham Website

[Service]
ExecStart=/usr/local/bin/hackbellingham
Restart=always
User=root
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=hackbellingham
Environment=HACK_BELLINGHAM_PORT=80
Environment=HACK_BELLINGHAM_SLACK_TEAM=hackbellingham
Environment=HACK_BELLINGHAM_SLACK_TOKEN=XXXX-XXXXXXXXXXX-XXXXXXXXXXX-XXXXXXXXXXX-XXXXXXXXXX
Environment=HACK_BELLINGHAM_MAILCHIMP_TOKEN=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX-us1
Environment=HACK_BELLINGHAM_MAILCHIMP_LIST=XXXXXXXXXX

[Install]
WantedBy=multi-user.target
```


[go]: http://www.golang.org
[systemd]: https://freedesktop.org/wiki/Software/systemd/
