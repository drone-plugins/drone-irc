package main

type Params struct {
	Channel   string `json:"channel"`
	Recipient string `json:"recipient"`
	Prefix    string `json:"prefix"`
	Nick      string `json:"nick"`
	Server    struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Password string `json:"password"`
		TLS      bool   `json:"tls"`
	} `json:"server"`
}
