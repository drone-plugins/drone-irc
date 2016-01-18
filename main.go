package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
	"github.com/thoj/go-ircevent"
)

var (
	buildDate string
)

func main() {
	fmt.Printf("Drone IRC Plugin built at %s\n", buildDate)

	system := drone.System{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("system", &system)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.Channel) == 0 && len(vargs.Recipient) == 0 {
		fmt.Println("Please provide a channel or recipient")

		os.Exit(1)
		return
	}

	if len(vargs.Prefix) == 0 {
		vargs.Prefix = "build"
	}

	if vargs.Server.Port == 0 {
		vargs.Server.Port = 6667
	}

	if len(vargs.Nick) == 0 {
		r := rand.New(rand.NewSource(99))
		vargs.Nick = fmt.Sprintf("drone%d", r.Int31())
	}

	client := irc.IRC(
		vargs.Nick,
		vargs.Nick)

	if client == nil {
		fmt.Println("Failed to create IRC Client: Invalid nick?")

		os.Exit(1)
		return
	}

	client.Password = vargs.Server.Password
	client.UseTLS = vargs.Server.TLS

	err := client.Connect(
		net.JoinHostPort(
			vargs.Server.Host,
			strconv.Itoa(vargs.Server.Port)))

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
		return
	}

	go func() {
		if err := <-client.ErrorChan(); err != nil {
			fmt.Println(err)

			os.Exit(1)
			return
		}
	}()

	client.AddCallback("001", func(_ *irc.Event) {
		var destination string

		if len(vargs.Recipient) != 0 {
			destination = vargs.Recipient
		} else {
			if strings.HasPrefix(vargs.Channel, "#") {
				destination = vargs.Channel
			} else {
				destination = "#" + vargs.Channel
			}
		}

		if strings.HasPrefix(destination, "#") {
			client.Join(destination)
		}

		client.Privmsgf(
			destination,
			"[%s %s/%s#%s] %s on %s by %s (%s/%s/%v)",
			vargs.Prefix,
			repo.Owner,
			repo.Name,
			build.Commit[:8],
			build.Status,
			build.Branch,
			build.Author,
			system.Link,
			repo.FullName,
			build.Number)

		if strings.HasPrefix(destination, "#") {
			client.Part(destination)
		}

		client.Quit()
	})

	client.Loop()
}
