# drone-irc

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-irc/status.svg)](http://beta.drone.io/drone-plugins/drone-irc)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-irc/coverage.svg)](https://aircover.co/drone-plugins/drone-irc)
[![](https://badge.imagelayers.io/plugins/drone-irc:latest.svg)](https://imagelayers.io/?images=plugins/drone-irc:latest 'Get your own badge on imagelayers.io')

Drone plugin to send build status notifications via IRC

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-irc <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "channel": "development",
        "nick": "test-drone",
        "server": {
            "port": 6697,
            "host": "irc.foobar.com",
            "password": "pa$$word",
            "tls": true
        }
    }
}
EOF
```

## Docker

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i plugins/drone-irc <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "channel": "development",
        "nick": "test-drone",
        "server": {
            "port": 6697,
            "host": "irc.foobar.com",
            "password": "pa$$word",
            "tls": true
        }
    }
}
EOF
```
