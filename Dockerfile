# Docker image for Drone's slack notification plugin
#
#     docker build --rm=true -t plugins/drone-irc .

FROM library/golang:1.4

# copy the local package files to the container's workspace.
ADD . /go/src/github.com/drone-plugins/drone-irc/

# build the slack plugin inside the container.
RUN go get github.com/drone-plugins/drone-irc/... && \
    go install github.com/drone-plugins/drone-irc

# run the plugin when the container starts
ENTRYPOINT ["/go/bin/drone-irc"]

