package messages

import (
	"strings"

	"github.com/urfave/cli/v3"
)

func FullMessage(cmd *cli.Command) string {
	message := []string{}
	for i := range cmd.Args().Len() {
		message = append(message, cmd.Args().Get(i))
	}
	return strings.Join(message, " ")
}
