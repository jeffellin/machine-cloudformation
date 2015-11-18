package main

import (
	"github.com/jeffellin/machine-cloudformation"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(amazoncf.NewDriver("", ""))
}
