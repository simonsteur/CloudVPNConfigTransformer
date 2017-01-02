package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func doJunos(c *cli.Context) error {

	fmt.Println("this would be the Junos stuff yea")

	if aws == true {
		fmt.Println("this would also be for aws yea")
	}
	return nil
}
