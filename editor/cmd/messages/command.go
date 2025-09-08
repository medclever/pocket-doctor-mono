package messages

import (
	"github.com/urfave/cli/v3"
)

func New() *cli.Command {
	return &cli.Command{
		Name: "messages",
		Commands: []*cli.Command{
			listCommand(),
		},
	}
}
