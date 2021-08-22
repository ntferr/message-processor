package settings

import (
	"log"

	env "github.com/Netflix/go-env"
)

type Settings struct {
	Port string `env:"PORT"`
	Host string `env:"HOST"`

	AWS struct {
		URL        string `env:"QUEUE_URL,required"`
		Region     string `env:"AWS_REGION,required"`
		MaxRetries int    `env:"MAX_RETRIES,required"`
	}
}

var settings Settings

func init() {
	_, err := env.UnmarshalFromEnviron(&settings)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSettings() Settings {
	return settings
}
