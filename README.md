# drone-irc
Drone plugin for sending IRC messages.

# Testing

To test this works, you can run this command on your command-line shell:

```bash
./drone-irc <<EOF
{
    "repo" : {
        "host": "github.com",
        "owner": "foo",
        "name": "bar",
        "self_url": "http://my.drone.io/foo/bar"
    },
    "commit" : {
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
		  "host": "some.irc.server",
	},
	"nick": "test-drone"
    }
}
EOF
```