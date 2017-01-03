package main

import (
	"os"

	"github.com/urfave/cli"
)

var (

	// providers
	aws   bool
	azure bool

	// config values
	file string
	t1   string
	t2   string
	zone string
	oif  string
	cidr string

	// generic
	comments bool
)

func main() {

	app := cli.NewApp()
	app.Name = "cvct"
	app.Version = "0.1"
	app.Usage = "transform generated vpn configuration from cloud providers to represent your environment"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:  "screenos",
			Usage: "use when dealing with ScreenOS configuration",
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
					Name:        "zone",
					Usage:       "use to specify which zone these interfaces should be bound to",
					Destination: &zone,
				},
				cli.StringFlag{
					Name:        "outgoing_interface, outgoingif, oif",
					Usage:       "use to specify the outgoing interface for the vpn",
					Destination: &oif,
				},
				cli.StringFlag{
					Name:        "cidr",
					Usage:       "use to specify the cidr network address to use for routing over the vpn tunnel",
					Destination: &cidr,
				},
				cli.BoolFlag{
					Name:        "nocomments, nc",
					Usage:       "Use if you want all comments stripped from the output",
					Destination: &comments,
				},
			},
			Action: TransformScreenOSConfig,
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
			Action: TransformJunOSConfig,
		},
	}

	app.Run(os.Args)
}
