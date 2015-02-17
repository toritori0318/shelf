package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "shelf"
	app.Version = Version
	app.Usage = "Shelf is for easy putting in and out of data."
	app.Author = "toritori0318"
	app.Email = ""
	app.Commands = Commands

	app.Run(os.Args)
}
