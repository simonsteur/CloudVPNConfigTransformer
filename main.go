package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {

	var OS string

	app := cli.NewApp()
	app.Name = "cvct"
	app.Version = "0.1"
	app.Usage = "transform generated vpn configuration from cloud providers to represent your environment"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "netscreen",
			Usage:       "transform Juniper Netscreen configuration",
			Destination: &OS,
		},
		cli.StringFlag{
			Name:        "junos",
			Usage:       "transform Juniper JunOS configuration",
			Destination: &OS,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Action = func(c *cli.Context) error {
		if OS == "netscreen" {
			fmt.Println("this would be the netscreen stuff yea")
		}
		if OS == "junos" {
			fmt.Println("this would be the junos stuff yea")

		}
		return nil
	}

	app.Run(os.Args)
}
