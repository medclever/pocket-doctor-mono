package languages

import (
	"context"
	"editor/app/types"
	"fmt"

	"github.com/urfave/cli/v3"
)

func listCommand(app types.App) *cli.Command {
	return &cli.Command{
		Name: "ls",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			for _, language := range app.LanguagesGetList().GetLanguages() {
				fmt.Println(language.View())
			}
			return nil
		},
	}
}
