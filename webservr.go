package main

import (
	"fmt"
	"log"
	"net/http"

	c "github.com/zephyyrr/autobot/config"
)

func init() {
	http.HandleFunc("/hook", handleHook)
	http.HandleFunc("/config", handleConfig)
	http.HandleFunc("/", handleRoot)
}

const rootView = `
<html>
<head>
<script>
function hook(method, callback) {
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange = callback;
	xmlhttp.open("POST", "/hook", true);
	xmlhttp.setRequestHeader("Sch-Github-Event", method);
	xmlhttp.send();
}

function pong() {
	console.log("Ping");
	hook("ping", null);
}

function push() {
	console.log("Push");
	hook("push", null);
}

function release() {
	console.log("Release");
	hook("release", null);
}

function deploy() {
	console.log("Deployment");
	hook("deployment", null);
}
</script>
</head>
<body>
Welcome to Autobot!<br>
Autobot is designed to be interfaced with Github's Webhooks API.<br>
To use, setup your projects Github repository to use <a href="/hook">this</a> as a webhook reciever.<br>
<ul>

	<li><a onclick="pong();" href="#">Test Ping</a></li>
	<li><a onclick="push();" href="#">Test Push</a></li>
	<li><a onclick="release();" href="#">Test Release</a></li>
	<li><a onclick="deploy();" href="#">Test Deployment</a></li>
</ul>
</body>
</html>
`

func handleRoot(w http.ResponseWriter, req *http.Request) {
	log.Println("Serving root")
	fmt.Fprintln(w, rootView)
}

func handleHook(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" && req.Method != "post" {
		log.Println("Redirecting to root")
		handleRoot(w, req)
		return
	}

	processes.Add(1)
	defer processes.Done()

	event := req.Header.Get("Sch-Github-Event")
	log.Printf("Hook (%s) called", event)

	if output, err := rollOut(config.Events[event]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(output)
	}
}

func handleConfig(w http.ResponseWriter, req *http.Request) {
	log.Println("Serving configfile")
	c.WriteConfig(w, c.DefautlConfig())
}
