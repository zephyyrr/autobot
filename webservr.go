package main

import (
	"fmt"
	"net/http"

	c "github.com/zephyyrr/autobot/config"
)

func init() {
	http.HandleFunc("/hook", handleHook)
	http.HandleFunc("/", handleRoot)
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<html><body>")
	fmt.Fprintln(w, "Welcome to Autobot!<br>")
	fmt.Fprintln(w, "Autobot is designed to be interfaced with Github's Webhooks API.<br>")
	fmt.Fprintf(w, "To use, setup your projects Github repository to use the Webhook %s.<br>\n", hookAddress(req))
	fmt.Fprintln(w, "</body></html>")
}

func handleHook(w http.ResponseWriter, req *http.Request) {
	processes.Add(1)
	defer processes.Done()

	event := req.Header.Get("Sch-Github-Event")

	switch event {
	case c.Ping:
	default:
		//Unknown event
		//Log and exit
	}

	if err := rollOut(config.Events[event]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

}
