package commands

import (
	"fmt"

	"github.com/bglebrun/rita/config"
	"github.com/bglebrun/rita/database"
	"github.com/urfave/cli"
)

func init() {

	databases := cli.Command{
		Name:  "show-databases",
		Usage: "Print the databases currently stored",
		Action: func(c *cli.Context) error {
			conf := config.InitConfig("")
			dbm := database.NewMetaDBHandle(conf)
			for _, name := range dbm.GetDatabases() {
				fmt.Println(name)
			}
			return nil
		},
	}

	bootstrapCommands(databases)
}
