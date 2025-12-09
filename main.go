package main

import "terrakube/cmd"

var version = "DEV"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
