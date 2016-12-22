package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {

	var netscreen bool
	var junos bool

	app := cli.NewApp()
	app.Name = "cvct"
	app.Version = "0.1"
	app.Usage = "transform generated vpn configuration from cloud providers to represent your environment"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "netscreen",
			Usage:       "transform Juniper Netscreen configuration",
			Destination: &netscreen,
		},
		cli.BoolFlag{
			Name:        "junos",
			Usage:       "transform Juniper JunOS configuration",
			Destination: &junos,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Action = func(c *cli.Context) error {
		if junos == true && netscreen == true {
			fmt.Println("please only select one language")
			os.Exit(1)
		}
		if netscreen == true {
			fmt.Println("this would be the netscreen stuff yea")
		}
		if junos == true {
			fmt.Println("this would be the junos stuff yea")

		}
		return nil
	}

	app.Run(os.Args)
}
