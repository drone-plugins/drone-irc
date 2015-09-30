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
        "full_name": "foo/bar"
    },
    "system": {
        "link_url": "http://drone.mycompany.com"
    },
    "build" : {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "commit": "9f2849d5",
        "branch": "master",
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
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
    "build" : {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "head_commit": {
            "sha": "9f2849d5",
            "branch": "master",
            "message": "Update the Readme",
            "author": {
                "login": "johnsmith",
                "email": "john.smith@gmail.com"
            }
        }
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
