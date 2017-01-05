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
	file         string
	t1           string
	t2           string
	zone         string
	externalif   string
	cidr         string
	externalzone string

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
					Usage:       "The generated configuration file",
					Destination: &file,
				},
				cli.StringFlag{
					Name:        "tunnel-interface-1, if1",
					Usage:       "Tunnel interface for Tunnel 1",
					Destination: &t1,
				},
				cli.StringFlag{
					Name:        "tunnel-interface-2, íf2",
					Usage:       "Tunnel interface for Tunnel 2",
					Destination: &t2,
				},
				cli.StringFlag{
					Name:        "zone, internal-zone",
					Usage:       "The zone which the tunnel interfaces should be bound to",
					Destination: &zone,
				},
				cli.StringFlag{
					Name:        "outgoing-interface, external-if, eif",
					Usage:       "The external interface for your tunnel",
					Destination: &externalif,
				},
				cli.StringFlag{
					Name:        "cidr",
					Usage:       "The network address you wish to route over the tunnel",
					Destination: &cidr,
				},
				cli.BoolFlag{
					Name:        "nocomments, nc",
					Usage:       "Strips all comments from output",
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
					Usage:       "The generated configuration file",
					Destination: &file,
				},
				cli.StringFlag{
					Name:        "tunnel-interface-1, if1",
					Usage:       "Tunnel interface for Tunnel 1",
					Destination: &t1,
				},
				cli.StringFlag{
					Name:        "tunnel-interface-2, íf2",
					Usage:       "Tunnel interface for Tunnel 2",
					Destination: &t2,
				},
				cli.StringFlag{
					Name:        "zone, internal-zone",
					Usage:       "The zone which the tunnel interfaces should be bound to",
					Destination: &zone,
				},
				cli.StringFlag{
					Name:        "external-zone, untrust-zone",
					Usage:       "The zone which protects the external interfaces",
					Destination: &externalzone,
				},
				cli.StringFlag{
					Name:        "outgoing-interface, external-if, eif",
					Usage:       "The external interface for your tunnel",
					Destination: &externalif,
				},
				cli.StringFlag{
					Name:        "cidr",
					Usage:       "The network address you wish to route over the tunnel",
					Destination: &cidr,
				},
				cli.BoolFlag{
					Name:        "nocomments, nc",
					Usage:       "Strips all comments from output",
					Destination: &comments,
				},
			},
			Action: TransformJunOSConfig,
		},
	}

	app.Run(os.Args)
}
