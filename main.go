package main

import "github.com/szakalboss/hosting-bimmer/cmd"

func main() {
	app := cmd.New()
	app.RunAndExitOnError()
}
