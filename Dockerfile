FROM golang:1.10.3 AS builder

# Install dep to use for dependency management
ENV DEP_VERSION 0.4.1
RUN curl -o /usr/local/bin/dep -L https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 && \
    chmod a+x /usr/local/bin/dep

WORKDIR /go/src/github.com/bellinghamcodes/website

# Install dependencies
COPY Gopkg.* ./
RUN dep ensure -vendor-only

# Build go binary
COPY . .
ARG VERSION=unkown
RUN go build -a -tags netgo -ldflags "-w -X main.version=${VERSION}" -o bellinghamcodes


FROM scratch

ARG BUILD_DATE
ARG NAME="bellingham.codes website"
ARG VERSION=unkown
ARG VCS_REF=unkown
ARG VCS_URL=https://github.com/bellinghamcodes/website
ARG URL=http://bellingham.codes

ADD certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/bellinghamcodes/website/bellinghamcodes /bellinghamcodes

# The following environment varialbes can be used to configure this image:
#
# ENV ORGANIZATION_NAME
# ENV SLACK_TEAM
# ENV SLACK_TOKEN
# ENV MAILCHIMP_TOKEN
# ENV MAILCHIMP_LIST
# ENV MEETUP_NAME
# ENV MEETUP_FETCH_INTERVAL
# ENV CODE_OF_CONDUCT_GITHUB_REPO
# ENV CODE_OF_CONDUCT_FETCH_INTERVAL
ENV PORT 80


EXPOSE 80

LABEL maintainer="kevinstock@tantalic.com" \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.name=$NAME \
    org.label-schema.url=$URL \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vcs-url=$VCS_URL \
    org.label-schema.version=$VERSION \
    org.label-schema.schema-version="1.0"

ENTRYPOINT ["/bellinghamcodes"]
