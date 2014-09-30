package main

import "net/http"

func init() {
	http.HandleFunc("/push", handlePush)
	http.HandleFunc("/deploy", handleDeploy)
	http.HandleFunc("/release", handleRelease)
}

func handlePush(w http.ResponseWriter, req *http.Request) {
	processes.Add(1)
	defer processes.Done()
	http.Error(w, "", http.StatusNotImplemented)
}

func handleDeploy(w http.ResponseWriter, req *http.Request) {
	processes.Add(1)
	defer processes.Done()
	http.Error(w, "", http.StatusNotImplemented)
}

func handleRelease(w http.ResponseWriter, req *http.Request) {
	processes.Add(1)
	defer processes.Done()
	http.Error(w, "", http.StatusNotImplemented)
}
