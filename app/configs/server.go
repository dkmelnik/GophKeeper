package configs

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Server struct {
	ADDR     string `required:"true"`
	Secret   string `required:"true"`
	Exp      int    `required:"true"`
	MongoDSN string `required:"true" split_words:"true"`
	MongoDB  string `required:"true" split_words:"true"`

	ExpDuration time.Duration
}

func NewServer() (Server, error) {
	var c Server
	if err := envconfig.Process("server", &c); err != nil {
		return c, err
	}
	c.ExpDuration = time.Duration(c.Exp) * time.Hour

	return c, nil
}
