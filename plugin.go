package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
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
	client.Debug = p.Config.IRCDebug
	client.UseSASL = p.Config.IRCSASL
	if p.Config.IRCSASL {
		client.SASLLogin = p.Config.Nick
		client.SASLPassword = p.Config.SASLPassword
	}

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
		txt, err := RenderTrim(p.Config.Template, p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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
