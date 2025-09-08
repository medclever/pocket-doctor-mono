package main

import (
	"context"
	"editor/app"
	"editor/cmd/messages"
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
		},
	}
	if err := c.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
