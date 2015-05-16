# Docker image for an IRC  plugin
#
#     CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t plugins/drone-irc .

FROM gliderlabs/alpine:3.1
RUN apk-install ca-certificates
ADD drone-irc /bin/
ENTRYPOINT ["/bin/drone-irc"]
