package main

import (
	"log"

	"github.com/drone/drone-plugin-go/plugin"
	"github.com/thoj/go-ircevent"
)

func main() {
	repo := plugin.Repo{}
	commit := plugin.Commit{}
	
	var params struct {
		Nick    string `json:"nick"`
		
		Server  struct {
			Host string `json:"host"`
			Port int `json:"port"`
			Password string `json:"password"`
		} `json:"server"`
		
		Channel string `json:"channel"`
	}

	plugin.Param("repo", &repo)
	plugin.Param("commit", &commit)
	plugin.Param("vargs", &params)
	
	if err := plugin.Parse(); err != nil {
		log.Fatal(err)
	}
	
	if params.Server.Port == 0 {
		params.Server.Port = 6667
	}

	client := irc.IRC(params.Nick, params.Nick)
	if client == nil {
		log.Fatal("Failed to make IRC Client: Invalid nick?")
	}

	client.Password = params.Server.Password
	
	err := client.Connect(fmt.Fprintf("%s:%d", params.Server.Host, params.Server.Port))
	if err != nil {
		log.Fatal(err)
	}

	client.AddCallback("001", func(_ *irc.Event) {
		client.Noticef(params.Channel, "[Drone %s/%s] %s (%s/%d)", repo.Owner, repo.Name, commit.State, repo.Self, commit.Sequence)
		client.Quit()
	})

	client.Loop()
}
