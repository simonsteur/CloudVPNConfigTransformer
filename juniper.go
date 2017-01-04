package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/urfave/cli"
)

import "strings"

// TransformScreenOSConfig transforms the ScreenOS configuration to the desired values
func TransformScreenOSConfig(c *cli.Context) error {

	input, err := ioutil.ReadFile(file)
	if err != nil {
		handleError(err)
	}

	values := make(map[string]string)
	values["tunnel_interface_1"] = t1
	values["tunnel_interface_2"] = t1
	values["internal_zone"] = zone
	values["cidr"] = cidr
	values["external_interface"] = external_if
	checkValueWarnings(values)

	r := strings.NewReplacer(
		"tunnel.1", t1,
		"tunnel.2", t2,
		"Trust", zone,
		"10.0.0.0/16", cidr,
		"ethernet0/0", external_if,
		"# set route", "set route",
	)

	VpnConfig := r.Replace(string(input))

	if comments == true {
		regex := regexp.MustCompile("(?m)[\r\n]+^.*#.*$")
		output := regex.ReplaceAllString(VpnConfig, "")
		fmt.Println(output)
	} else {
		fmt.Println(VpnConfig)
	}

	return nil

}

// TransformJunOSConfig transforms the JunOS configuration to the desired values
func TransformJunOSConfig(c *cli.Context) error {
	input, err := ioutil.ReadFile(file)
	if err != nil {
		handleError(err)
	}

	values := make(map[string]string)
	values["tunnel_interface_1"] = t1
	values["tunnel_interface_2"] = t1
	values["internal_zone"] = zone
	values["external_zone"] = zone
	values["cidr"] = cidr
	values["external_interface"] = external_if
	checkValueWarnings(values)

	r := strings.NewReplacer(
		"st0.1", t1,
		"st0.2", t2,
		"Trust", zone,
		"Untrust", external_zone,
		"10.0.0.0/16", cidr,
		"ge-0/0/0.0", external_if,
		"# set routing-options", "set routing-options",
	)

	VpnConfig := r.Replace(string(input))

	if comments == true {
		regex := regexp.MustCompile("(?m)[\r\n]+^.*#.*$")
		output := regex.ReplaceAllString(VpnConfig, "")
		fmt.Println(output)
	} else {
		fmt.Println(VpnConfig)
	}

	return nil
}
