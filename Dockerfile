FROM scratch
MAINTAINER Kevin Stock <kevinstock@tantalic.com>

ADD certs/ca-certificates.crt /etc/ssl/certs/
ADD build/bellinghamcodes-linux-amd64 bellinghamcodes

ENV PORT 80
# ENV SLACK_TEAM
# ENV SLACK_TOKEN
# ENV MAILCHIMP_TOKEN
# ENV MAILCHIMP_LIST

EXPOSE 80

ENTRYPOINT ["/bellinghamcodes"]
