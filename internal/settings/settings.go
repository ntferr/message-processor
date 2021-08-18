package settings

import (
	"log"

	env "github.com/Netflix/go-env"
)

type Settings struct {
	Port string `env:"PORT"`
	Host string `env:"HOST"`

	AWS struct {
	}
}

var settings Settings

func init() {
	_, err := env.UnmarshalFromEnviron(&settings)
	if err != nil {
		log.Fatal(err)
	}
}
