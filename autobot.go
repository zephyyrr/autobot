package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/zephyyrr/autobot/config"
)

var (
	configuration config.Config
	processes     sync.WaitGroup
	configfile    = flag.String("f", "autobot.toml", "Config filename")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.Parse()
	f, err := os.Open(*configfile)
	if err != nil {
		log.Fatalf("Unable to read open config file %s. %s", *configfile, err)
	}
	configuration, err = config.LoadConfig(f)
	if err != nil {
		log.Fatalln("Unable to parse config file.", err)
	}
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
	// Pull from git
	// Issue build command (read from config)
	// ???
	// Profit!
}

func shutdown() {
	// Allows clean shutdowns.
	processes.Wait()
}
