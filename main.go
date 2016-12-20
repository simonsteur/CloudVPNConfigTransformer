package main

import (
	"github.com/urfave/cli"
)

func main() {

    var os string

	app := cli.NewApp()
    app.Name = "cvct"
    app.Version = 0.1
    app.Usage = "transform generated vpn configuration from cloud providers to represent your environment"

    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "netscreen",
            value: "",
            Usage: "transform netscreen configuration",
            Destination: &os,
        }
    }
    app.Action = func(c *cli.Context) error {
    
	}
}
