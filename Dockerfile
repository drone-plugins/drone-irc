# Docker image for the Drone IRC plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-irc
#     make deps build
#     docker build --rm=true -t plugins/drone-irc .

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-irc /bin/
ENTRYPOINT ["/bin/drone-irc"]
