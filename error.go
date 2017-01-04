package main

import "github.com/fatih/color"
import "os"

func handleError(err error) {

	color.Red("error: %s", err)
	os.Exit(1)
}

func handleWarning(warning string) {

	color.Yellow("Warning: %s", warning)

}
