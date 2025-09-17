package languages

import (
	"editor/app"

	"github.com/urfave/cli/v3"
)

func New(app *app.App) *cli.Command {
	return &cli.Command{
		Name: "languages",
		Commands: []*cli.Command{
			listCommand(app),
			addCommand(app),
		},
	}
}
