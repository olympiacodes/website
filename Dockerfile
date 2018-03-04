FROM golang:1.10.0 AS builder

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

ADD certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/bellinghamcodes/website/bellinghamcodes /bellinghamcodes

ENV PORT 80
# ENV ORGANIZATION_NAME
# ENV SLACK_TEAM
# ENV SLACK_TOKEN
# ENV MAILCHIMP_TOKEN
# ENV MAILCHIMP_LIST
# ENV MEETUP_NAME
# ENV MEETUP_FETCH_INTERVAL

EXPOSE 80

LABEL maintainer="kevinstock@tantalic.com"
ENTRYPOINT ["/bellinghamcodes"]
