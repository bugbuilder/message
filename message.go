package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang/glog"
)

// Message represent the message from the host
type Message struct {
	Host   string   `json:"host"`
	Quotes []string `json:"quotes"`
}

func builder(c *Config) Message {
	host, _ := os.Hostname()

	glog.V(2).Infoln("Building message")

	var quotes []string
	for t, colleagues := range c.Teams {
		cs := strings.Join(colleagues, ",")
		quotes = append(quotes, fmt.Sprintf("%s, %s: %s", c.Message, t, cs))
	}

	return Message{
		host,
		quotes,
	}
}
