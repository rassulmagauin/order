package env

import (
	_ "github.com/kelseyhightower/envconfig"
)

type Env struct {
	HostNumber  string   `envconfig:"HOST_NUMBER"`
	Users       []string `envconfig:"USERS"`
	DatabaseUrl string   `envconfig:"DATABASE_URL"`
}
