package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func doNetscreen(c *cli.Context) error {

	fmt.Println("this would be the netscreen stuff yea")

	if aws == true {
		fmt.Println("this would also be for aws yea")
	}
	return nil
}
