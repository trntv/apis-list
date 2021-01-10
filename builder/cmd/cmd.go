package main

import (
	"github.com/apis-list/apis-list/builder"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "build",
		Action: func(c *cli.Context) error {
			list, err := builder.ReadList(c.Args().First())
			if err != nil {
				return err
			}

			dir, err := os.Getwd()
			if err != nil {
				return err
			}

			return builder.Render(list, dir)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
