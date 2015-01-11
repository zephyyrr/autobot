package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"

	c "github.com/zephyyrr/autobot/config"
)

var (
	config     c.Config
	processes  sync.WaitGroup
	configfile = flag.String("f", "autobot.toml", "Config filename")
	debug      = flag.Bool("d", false, "Set to true for debug printing.")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.Parse()
	f, err := os.Open(*configfile)
	if err != nil {
		log.Fatalf("Unable to read open config file %s. %s", *configfile, err)
	}
	config, err = c.LoadConfig(f)
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

	go http.Serve(conn, nil)

	select {
	case <-interrupt:
	}

	if err := conn.Close(); err != nil {

	}

	shutdown()
}

var (
	UnknownActionError = errors.New("Unknown Action")
)

func rollOut(actions []c.Action) (stdout []byte, err error) {
	// Pull from git
	// Issue build command (read from config)
	// ???
	// Profit!

	for _, action := range actions {
		log.Printf("%s: %s", action.Type, action.Payload)
		switch action.Type {
		case c.Command:
			s := strings.Split(action.Payload, " ")
			cmd := exec.Command(s[0], s[1:]...)
			stdout, err = HandleCommand(cmd)
			if err != nil {
				return stdout, err
			}

		case c.Install:
			cmd := exec.Command("go", "install", action.Payload)
			stdout, err = HandleCommand(cmd)
			if err != nil {
				return stdout, err
			}

		case c.Test:
			cmd := exec.Command("go", "test", action.Payload)
			stdout, err = HandleCommand(cmd)
			if err != nil {
				return stdout, err
			}
		default:
			return []byte{}, UnknownActionError
		}
	}
	return

}

func HandleCommand(cmd *exec.Cmd) (stdout []byte, err error) {
	stdout, stderr := cmd.Output()
	if err != nil {
		log.Printf("Error output of last command: \n%s", stderr)
		return
	}

	if *debug {
		fmt.Println(string(stdout))
	}

	if stderr != nil {
		err = stderr
	}
	return
}

func shutdown() {
	// Allows clean shutdowns.
	processes.Wait()
}

func hookAddress(req *http.Request) string {
	return fmt.Sprintf("%s%s", req.Host, "/hook")
}
