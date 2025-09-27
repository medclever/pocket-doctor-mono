package languages

import (
	"editor/app/types"

	"github.com/urfave/cli/v3"
)

func New(app types.App) *cli.Command {
	return &cli.Command{
		Name: "languages",
		Commands: []*cli.Command{
			listCommand(app),
			addCommand(app),
		},
	}
}
