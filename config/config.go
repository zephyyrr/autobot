package config

import (
	"io"

	"github.com/BurntSushi/toml"
)

type Config struct {
}

func LoadConfig(r io.Reader) (c Config, err error) {
	_, err = toml.DecodeReader(r, &c)
	return
}
