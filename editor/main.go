package main

import (
	"context"
	"editor/app"
	"editor/cmd/languages"
	"editor/cmd/messages"
	"editor/cmd/ui"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := app.New()
	app.Init()
	c := cli.Command{
		Commands: []*cli.Command{
			messages.New(app),
			languages.New(app),
			ui.New(app),
		},
	}
	if err := c.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
