package languages

import (
	"context"
	"editor/app"
	"errors"

	"github.com/urfave/cli/v3"
)

func addCommand(app *app.App) *cli.Command {
	return &cli.Command{
		Name:        "add",
		Description: "add <language code> <name>",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() != 2 {
				return errors.New("Yor should receive lang_code and lang_name")
			}
			app.LanguagesAdd(cmd.Args().Get(0), cmd.Args().Get(1))
			return app.Persist()
		},
	}
}
