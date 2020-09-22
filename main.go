package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "discuz-to-markdown",
		Usage: "export discuz article to markdown",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config, c",
				Usage: "Load configuration from `FILE`",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Init configuration for export",
				Action:  initConfiguration,
			},
			{
				Name:    "export",
				Aliases: []string{"e"},
				Usage:   "Export article data",
				Action:  exportArticle,
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func initConfiguration(c *cli.Context) error {
	return nil
}

func exportArticle(c *cli.Context) error {
	return nil
}
