package main

import (
	"fmt"
	"os"
	"terrakube/cmd"
	"time"

	"github.com/fatih/color"
)

var version = fmt.Sprintf("DEV-%s", time.Now().Format("20060102T150405"))

func main() {
	if err := cmd.NewRootCmd(version).Execute(); err != nil {
		color.HiRed(err.Error())
		os.Exit(1)
	}
}
