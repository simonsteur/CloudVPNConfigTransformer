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
	file          string
	t1            string
	t2            string
	zone          string
	external_if   string
	cidr          string
	external_zone string

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
					Name:        "tunnel_interface_1, if1",
					Usage:       "use to specify which tunnel interface should be used for Tunnel 1",
					Destination: &t1,
				},
				cli.StringFlag{
					Name:        "tunnel_interface_2, íf2",
					Usage:       "use to specify which tunnel interface should be used for Tunnel 2",
					Destination: &t2,
				},
				cli.StringFlag{
					Name:        "zone",
					Usage:       "use to specify which zone these interfaces should be bound to",
					Destination: &zone,
				},
				cli.StringFlag{
					Name:        "outgoing_interface, external_if, eif",
					Usage:       "use to specify the outgoing interface for the vpn",
					Destination: &external_if,
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
				cli.StringFlag{
					Name:        "file, f",
					Usage:       "use to specify the location of the file holding the generated vpn configuration",
					Destination: &file,
				},
				cli.StringFlag{
					Name:        "tunnel_interface_1, if1",
					Usage:       "use to specify which tunnel interface should be used for Tunnel 1",
					Destination: &t1,
				},
				cli.StringFlag{
					Name:        "tunnel_interface_2, íf2",
					Usage:       "use to specify which tunnel interface should be used for Tunnel 2",
					Destination: &t2,
				},
				cli.StringFlag{
					Name:        "zone",
					Usage:       "use to specify which zone these interfaces should be bound to",
					Destination: &zone,
				},
				cli.StringFlag{
					Name:        "external_zone, untrust_zone",
					Usage:       "The zone which protects the external interfaces",
					Destination: &zone,
				},
				cli.StringFlag{
					Name:        "outgoing_interface, external_if, eif",
					Usage:       "use to specify the external interface for the vpn",
					Destination: &external_if,
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
			Action: TransformJunOSConfig,
		},
	}

	app.Run(os.Args)
}
