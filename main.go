// Copyright (c) 2018, Vincent Landgraf
// SEE LICENSE file for details

package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var remoteHost = "google.de"
var interval = time.Second * 30
var online = true
var offlineSince time.Time

func main() {
	// configure using command line
	flag.StringVar(&remoteHost, "remote-host", remoteHost, "set the remote host to check against (using http)")
	flag.DurationVar(&interval, "duration", interval, "set duration check interval")
	flag.Parse()

	// start a timer for checking remote side
	t := time.NewTicker(interval)
	for {
		select {
		case <-t.C:
			go checkHTTP()
		}
	}
}

func checkHTTP() {
	timeout := interval

	r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/", remoteHost), strings.NewReader(""))
	if err != nil {
		fmt.Errorf("Error HTTP check failed: %s\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	r = r.WithContext(ctx)

	c := http.DefaultClient
	_, err = c.Do(r)
	if err != nil {
		if online {
			online = false
			offlineSince = time.Now()
		}
		return
	}
	if !online {
		online = true
		now := time.Now()
		fmt.Printf("%s\t%s\n", now.Format(time.RFC3339), now.Sub(offlineSince))
	}
}
