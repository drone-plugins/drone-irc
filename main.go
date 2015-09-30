package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/drone/drone-plugin-go/plugin"
	"github.com/thoj/go-ircevent"
)

type Arguments struct {
	Prefix string `json:"prefix"`
	Nick   string `json:"nick"`

	Server struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Password string `json:"password"`
		TLS      bool   `json:"tls"`
	} `json:"server"`

	Channel string `json:"channel"`
}

func main() {
	repo := plugin.Repo{}
	build := plugin.Build{}
	system := plugin.System{}
	vargs := Arguments{}

	plugin.Param("build", &build)
	plugin.Param("repo", &repo)
	plugin.Param("system", &system)
	plugin.Param("vargs", &vargs)

	if err := plugin.Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if vargs.Nick == "" {
		r := rand.New(rand.NewSource(99))
		vargs.Nick = fmt.Sprintf("drone%d", r.Int31())
	}

	if vargs.Prefix == "" {
		vargs.Prefix = "drone"
	}

	if vargs.Server.Port == 0 {
		vargs.Server.Port = 6667
	}

	client := irc.IRC(vargs.Nick, vargs.Nick)

	if client == nil {
		fmt.Println("Failed to make IRC Client: Invalid nick?")
		os.Exit(1)
	}

	client.Password = vargs.Server.Password
	client.UseTLS = vargs.Server.TLS

	if err := client.Connect(fmt.Sprintf("%s:%d", vargs.Server.Host, vargs.Server.Port)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func() {
		if err := <-client.ErrorChan(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	client.AddCallback("001", func(_ *irc.Event) {
		if strings.HasPrefix(vargs.Channel, "#") {
			client.Join(vargs.Channel)
		}

		client.Privmsgf(
			vargs.Channel,
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

		if strings.HasPrefix(vargs.Channel, "#") {
			client.Part(vargs.Channel)
		}

		client.Quit()
	})

	client.Loop()
}
