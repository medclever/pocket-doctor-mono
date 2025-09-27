package messages

import (
	"editor/app/types"

	"github.com/urfave/cli/v3"
)

func New(app types.App) *cli.Command {
	return &cli.Command{
		Name: "messages",
		Commands: []*cli.Command{
			listCommand(app),
			addCommand(app),
			translateCommand(app),
		},
	}
}
