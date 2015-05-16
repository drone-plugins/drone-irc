package main

import (
	"fmt"
	"os"

	"github.com/drone/drone-plugin-go/plugin"
	"github.com/thoj/go-ircevent"
)

func main() {
	repo := plugin.Repo{}
	commit := plugin.Commit{}

	var params struct {
		Nick string `json:"nick"`

		Server struct {
			Host     string `json:"host"`
			Port     int    `json:"port"`
			Password string `json:"password"`
			TLS      bool   `json:"tls"`
		} `json:"server"`

		Channel string `json:"channel"`
	}

	plugin.Param("repo", &repo)
	plugin.Param("commit", &commit)
	plugin.Param("vargs", &params)

	if err := plugin.Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if params.Server.Port == 0 {
		params.Server.Port = 6667
	}

	client := irc.IRC(params.Nick, params.Nick)
	if client == nil {
		fmt.Println("Failed to make IRC Client: Invalid nick?")
		os.Exit(1)
	}

	client.Password = params.Server.Password
	client.UseTLS = params.Server.TLS

	err := client.Connect(fmt.Sprintf("%s:%d", params.Server.Host, params.Server.Port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Listen for errors and immediately exit with
	// and error status.
	go func() {
		err := <-client.ErrorChan()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	client.AddCallback("001", func(_ *irc.Event) {
		client.Noticef(params.Channel, "[Drone %s/%s/%d] %s (%s/%d)", repo.Owner, repo.Name, commit.Sequence, commit.State, repo.Self, commit.Sequence)
		client.Quit()
	})

	client.Loop()
}
