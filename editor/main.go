package main

import (
	"context"
	"editor/cmd/messages"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	c := cli.Command{
		Commands: []*cli.Command{
			messages.New(),
		},
	}
	if err := c.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
