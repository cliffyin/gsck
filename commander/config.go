package commander

import (
	"fmt"
	"os"

	"github.com/EvanLi/gsck/command"
	"github.com/EvanLi/gsck/config"
	"github.com/codegangsta/cli"
)

func init() {
	command.RegisterCommand(cli.Command{
		Name:   "config",
		Usage:  "Get/Set Defaults [ $HOME/.gsckconfig ]",
		Action: configAction,
	})
}

// CONFIG Action (gsck config ...)
func configAction(c *cli.Context) {
	if len(c.Args()) == 0 {
		cli.ShowCommandHelp(c, "config")
		os.Exit(1)
	}
	key := c.Args()[0]
	if len(c.Args()) == 1 {
		fmt.Println(config.GetString(key))
	} else if len(c.Args()) == 2 {
		value := c.Args()[1]
		config.Set(key, value)
	}
}
