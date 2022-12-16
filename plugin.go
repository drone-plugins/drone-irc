package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/drone/drone-template-lib/template"
	"github.com/pkg/errors"
	irc "github.com/thoj/go-ircevent"
)

type (
	Repo struct {
		Owner string
		Name  string
	}

	Build struct {
		Tag     string
		Event   string
		Number  int
		Commit  string
		Ref     string
		Branch  string
		Author  string
		Message string
		Status  string
		Link    string
		Started int64
		Created int64
	}

	Config struct {
		Prefix       string
		Nick         string
		Channel      string
		Recipient    string
		IRCHost      string
		IRCPort      int
		IRCPassword  string
		IRCEnableTLS bool
		IRCDebug     bool
		IRCSASL      bool
		SASLPassword string
		Template     string
	}

	Job struct {
		Started int64
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

func (p Plugin) Exec() error {
	if len(p.Config.Channel) == 0 && len(p.Config.Recipient) == 0 {
		return errors.New("Please provide a channel or recipient")
	}

	if len(p.Config.Nick) == 0 {
		r := rand.New(rand.NewSource(99))
		p.Config.Nick = fmt.Sprintf("drone%d", r.Int31())
	}

	client := irc.IRC(p.Config.Nick, p.Config.Nick)

	if client == nil {
		return errors.New("Failed to create IRC Client: Invalid nick?")
	}

	client.Password = p.Config.IRCPassword
	client.UseTLS = p.Config.IRCEnableTLS
	client.Debug = p.Config.IRCDebug
	client.UseSASL = p.Config.IRCSASL

	if p.Config.IRCSASL {
		client.SASLLogin = p.Config.Nick
		client.SASLPassword = p.Config.SASLPassword
	}

	err := client.Connect(net.JoinHostPort(p.Config.IRCHost, strconv.Itoa(p.Config.IRCPort)))

	if err != nil {
		return errors.Wrap(err, "failed to connect to server")
	}

	go func() {
		if err := <-client.ErrorChan(); err != nil {
			_ = errors.Wrap(err, "received an error from server")

			os.Exit(1)
			return
		}
	}()

	client.AddCallback("001", func(event *irc.Event) {
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

		txt, err := template.RenderTrim(p.Config.Template, p)

		if err != nil {
			_ = errors.Wrap(err, "failed to render template")

			os.Exit(1)
			return
		}

		client.Privmsg(destination, txt)

		if strings.HasPrefix(destination, "#") {
			client.Part(destination)
		}

		client.Quit()
	})

	client.Loop()

	return nil
}
