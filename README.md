# drone-irc

Drone plugin for sending IRC messages.

# Testing

To test this works, you can run this command on your command-line shell:

```bash
./drone-irc <<EOF
{
    "repo" : {
        "owner": "foo",
        "name": "bar",
        "self_url": "http://my.drone.io/foo/bar"
    },
    "commit" : {
        "sequence": 1
        "state": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "sha": "9f2849d5",
        "branch": "master",
        "pull_request": "800",
        "author": "john.smith@gmail.com",
        "message": "Update the Readme"
    },
    "vargs": {
        "channel": "#development",
        "server": {
          "port": 6697,
          "host": "irc.foobar.com",
          "password": "pa$$word",
          "tls": true
        },
        "nick": "test-drone"
    }
}
EOF
```
## Docker

Build the Docker container. Note that we need to use the `-netgo` tag so that
the binary is built without a CGO dependency:

```sh
CGO_ENABLED=0 go build -a -tags netgo
docker build --rm=true -t plugins/drone-irc .
```

Send an IRC notification:

```sh
docker run -i plugins/drone-irc <<EOF
{
    "repo" : {
        "owner": "foo",
        "name": "bar",
        "self_url": "http://my.drone.io/foo/bar"
    },
    "commit" : {
        "sequence": 1
        "state": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "sha": "9f2849d5",
        "branch": "master",
        "pull_request": "800",
        "author": "john.smith@gmail.com",
        "message": "Update the Readme"
    },
    "vargs": {
        "channel": "#development",
        "server": {
          "port": 6697,
          "host": "irc.foobar.com",
          "password": "pa$$word",
          "tls": true
        },
        "nick": "test-drone"
    }
}
EOF
```