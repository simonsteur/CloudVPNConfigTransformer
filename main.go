package main

import (
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cvct"
	app.Usage = "transform generated vpn configuration from cloud providers to represent your environment"
	app.Version = 0.1
	app.Action = func(c *cli.Context) error {

	}
}
