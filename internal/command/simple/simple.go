package simple

import (
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/urfave/cli/v2"

	"github.com/takumin/go-hcl-playground/internal/config"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "variable",
			Aliases:     []string{"v"},
			Usage:       "variable",
			EnvVars:     []string{"VARIABLE"},
			Value:       cfg.Variable,
			Destination: &cfg.Variable,
		},
	}...)
	return &cli.Command{
		Name:    "simple",
		Aliases: []string{"s"},
		Usage:   "simple",
		Flags:   flags,
		Action:  action(cfg),
	}
}

type Config struct {
	LogLevel string `hcl:"log_level"`
}

func action(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		var config Config
		if err := hclsimple.DecodeFile("config.hcl", nil, &config); err != nil {
			return err
		}
		log.Printf("%#v", config)
		return nil
	}
}
