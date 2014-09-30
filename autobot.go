package main

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var processes sync.WaitGroup

func main() {
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	conn := net.Listen("tcp", ":19840")

	http.Serve(conn, http.DefaultServeMux)

	select {
	case <-interrupt:
	}

	shutdown()
}

func rollOut() {

}

func shutdown() {
	conn.Close()
	// Allows clean shutdowns.
	processes.Wait()
}
