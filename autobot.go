package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/zephyyrr/autobot/config"
)

var (
	configuration Config
	processes     sync.WaitGroup
	configfile    = flag.String("f", "Config filename", "autobot.json")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.Parse()
	f, err := os.Open(configfile)
	if err != nil {
		log.Fatalln("Unable to read config file.", f)
	}
	config.LoadConfig(f)
}

func main() {
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	conn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln("Error listening to port.", err)
	}

	go http.Serve(conn, http.DefaultServeMux)

	select {
	case <-interrupt:
	}

	if err := conn.Close(); err != nil {

	}

	shutdown()
}

func rollOut() {

}

func shutdown() {

	// Allows clean shutdowns.
	processes.Wait()
}
