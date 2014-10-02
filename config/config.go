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
	Shell  Type = "shell"  // Calls <Command>
	Go          = "go"     // Calls go install <Package>
	Docker      = "docker" //Calls go test <Package>
)

type Action struct {
	Type Type
	Command
	//Filter Filter

	Parameter string
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
				Action{Type: Go, Command, Parameter: "github.com/zephyyrr/autobot"},
				Action{Type: Install, Command: "install", Parameter: "github.com/zephyyrr/autobot"},
				Action{Type: Shell, Command: "echo", Parameter: "test"},
			},
			Push: []Action{
				Action{Type: Go, Command: Install, Parameter: "github.com/zephyyrr/autobot"},
				Action{Type: Go, Command: Install, Parameter: "github.com/zephyyrr/goda"},
			},
		},
	}
}

func WriteConfig(w io.Writer, c Config) (err error) {
	enc := toml.NewEncoder(w)
	err = enc.Encode(c)
	return
}
