package main

import (
	"os"

	"github.com/urfave/cli"
)

var (
	aws   bool
	azure bool

	file string
	t1   string
	t2   string
	zone string
	cidr string
)

func main() {

	app := cli.NewApp()
	app.Name = "cvct"
	app.Version = "0.1"
	app.Usage = "transform generated vpn configuration from cloud providers to represent your environment"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:  "netscreen",
			Usage: "use when dealing with netscreen configuration",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "aws",
					Usage:       "use when dealing with aws generated configuration",
					Destination: &aws,
				},
				cli.StringFlag{
					Name:        "file, f",
					Usage:       "use to specify the location of the file holding the generated vpn configuration",
					Destination: &file,
				},
				cli.StringFlag{
					Name:        "tunnel1, t1",
					Usage:       "use to specify which tunnel interface should be used for Tunnel 1",
					Destination: &t1,
				},
				cli.StringFlag{
					Name:        "tunnel2, t2",
					Usage:       "use to specify which tunnel interface should be used for Tunnel 2",
					Destination: &t2,
				},
				cli.StringFlag{
					Name:        "zone, z",
					Usage:       "use to specify which zone these interfaces should be bound to",
					Destination: &zone,
				},
				cli.StringFlag{
					Name:        "outgoing_interface, outgoingif, oif",
					Usage:       "use to specify the outgoing interface for the vpn",
					Destination: &zone,
				},
				cli.StringFlag{
					Name:        "cidr",
					Usage:       "use to specify the cidr network address to use for routing over the vpn tunnel",
					Destination: &cidr,
				},
			},
			Action: doNetscreen,
		},
		{
			Name:  "junos",
			Usage: "use when dealing with junos configuration",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "aws",
					Usage:       "use when dealing with aws generated configuration",
					Destination: &aws,
				},
			},
			Action: doJunos,
		},
	}

	app.Run(os.Args)
}
