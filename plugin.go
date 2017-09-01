package main

import (
	"fmt"
	"os"
	"net"
	"strconv"
	"strings"
	"math/rand"
	"github.com/thoj/go-ircevent"
)

type (
	Repo struct {
		Owner string `json:"owner"`
		Name  string `json:"name"`
	}

	Build struct {
		Tag     string `json:"tag"`
		Event   string `json:"event"`
		Number  int    `json:"number"`
		Commit  string `json:"commit"`
		Ref     string `json:"ref"`
		Branch  string `json:"branch"`
		Author  string `json:"author"`
		Message string `json:"message"`
		Status  string `json:"status"`
		Link    string `json:"link"`
		Started int64  `json:"started"`
		Created int64  `json:"created"`
	}

	Config struct {
		Prefix string
		Nick string
		Channel string
		Recipient string
		IRCHost string
		IRCPort int
		IRCPassword string
		IRCEnableTLS bool
	}

	Job struct {
		Started int64 `json:"started"`
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

func (p Plugin) Exec() error {


	//system := drone.System{}
	//repo := drone.Repo{}
	//build := drone.Build{}
	//vargs := Params{}


	//plugin.Param("system", &system)
	//plugin.Param("repo", &repo)
	//plugin.Param("build", &build)
	//plugin.Param("vargs", &vargs)
	//plugin.MustParse()

	if len(p.Config.Channel) == 0 && len(p.Config.Recipient) == 0 {
		fmt.Println("Please provide a channel or recipient")
		return fmt.Errorf("Please provide a channel or recipient")
	}

	if len(p.Config.Nick) == 0 {
		r := rand.New(rand.NewSource(99))
		p.Config.Nick = fmt.Sprintf("drone%d", r.Int31())
	}

	client := irc.IRC(
		p.Config.Nick,
		p.Config.Nick)

	if client == nil {
		return fmt.Errorf("Failed to create IRC Client: Invalid nick?")
	}

	client.Password = p.Config.IRCPassword
	client.UseTLS = p.Config.IRCEnableTLS

	err := client.Connect(
		net.JoinHostPort(
			p.Config.IRCHost,
			strconv.Itoa(p.Config.IRCPort)))
	if err != nil {
		return err
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

		if len(p.Config.Recipient) != 0 {
			destination = p.Config.Recipient
		} else {
			if strings.HasPrefix(p.Config.Channel, "#") {
				destination = p.Config.Channel
			} else {
				destination = "#" + p.Config.Channel
			}
		}

		if strings.HasPrefix(destination, "#") {
			client.Join(destination)
		}

		client.Privmsgf(
			destination,
			"[%s %s/%s#%s] %s on %s by %s (%s)",
			p.Config.Prefix,
			p.Repo.Owner,
			p.Repo.Name,
			p.Build.Commit[:8],
			p.Build.Status,
			p.Build.Branch,
			p.Build.Author,
			p.Build.Link,)

		if strings.HasPrefix(destination, "#") {
			client.Part(destination)
		}

		client.Quit()
	})
	client.Loop()
	return nil
}