# drone-irc

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-irc/status.svg)](http://beta.drone.io/drone-plugins/drone-irc)
[![](https://badge.imagelayers.io/plugins/drone-irc:latest.svg)](https://imagelayers.io/?images=plugins/drone-irc:latest 'Get your own badge on imagelayers.io')

Drone plugin for sending build status notifications via IRC

## Usage

```sh
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
        "commit": "64908ed2414b771554fda6508dd56a0c43766831",
        "branch": "master",
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
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

Build the Docker container using `make`:

```sh
make deps build
docker build --rm=true -t plugins/drone-irc .
```

### Example

```sh
docker run -i plugins/drone-irc <<EOF
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
        "commit": "64908ed2414b771554fda6508dd56a0c43766831",
        "branch": "master",
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
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
