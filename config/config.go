package config

import (
	"io"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Address string
	Port    int

	Events map[string][]Action
}

const (
	Deployment = "deployment"
	Push       = "push"
	Ping       = "ping"
	Release    = "release"
)

type Type string

const (
	Command Type = "command" // Calls <Command>
	Install      = "install" // Calls go install <Package>
	Test         = "test"    //Calls go test <Package>
)

type Action struct {
	Type Type
	//Filter Filter

	Payload string
}

type Filter struct {
	Owner      string
	Repository string
	//Branch      string // Only if the event regards this branch. Not sure how to implement at this time
	Environment string
}

func LoadConfig(r io.Reader) (c Config, err error) {
	_, err = toml.DecodeReader(r, &c)
	return
}

func DefautlConfig() (c Config) {
	return Config{
		Address: "",
		Port:    8080,
		Events: map[string][]Action{
			Deployment: []Action{
				Action{Type: Test, Payload: "github.com/zephyyrr/autobot"},
				Action{Type: Install, Payload: "github.com/zephyyrr/autobot"},
				Action{Type: Command, Payload: "echo 'test'"},
			},
			Push: []Action{
				Action{Type: Install, Payload: "github.com/zephyyrr/autobot"},
				Action{Type: Install, Payload: "github.com/zephyyrr/goda"},
			},
		},
	}
}
