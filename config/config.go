package config

import (
	"flag"
	"io"

	"github.com/BurntSushi/toml"
)

var configfile = flag.String("f", "Config filename", "autobot.json")

type Config struct {
}

func LoadConfig(r io.Reader) (c Config, err error) {
	_, err = toml.DecodeReader(r, &c)
	return
}
