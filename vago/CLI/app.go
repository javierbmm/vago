package CLI

import (
	"log"
	"os"
	"vago/vago/generator"
	"vago/vago/input"
	"vago/vago/server"

	"github.com/urfave/cli/v2"
)

func Run() {
	var configFile string
	var port int
	var noLog, noTime bool

	app := &cli.App{
		Name:  "vago",
		Usage: "A minimalistic Static Site Generator to create wonderful websites from Markdown files.\n",

		Commands: []*cli.Command{
			{
				Name:  "build",
				Usage: " Build/Generate web content from markdown files on indicated folder (default: config.yaml).",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "config",
						Value:       "config.yaml",
						Aliases:     []string{"c"},
						Usage:       "Load configuration from `FILE`.",
						Destination: &configFile,
					},
					&cli.BoolFlag{
						Name:        "no-log",
						Value:       false,
						Aliases:     []string{"nl"},
						Usage:       "Don't display logs for every page.",
						Destination: &noLog,
					},
					&cli.BoolFlag{
						Name:        "no-time",
						Value:       false,
						Aliases:     []string{"ct"},
						Usage:       "Don't display timestamps when logging.",
						Destination: &noTime,
					},
				},
				Action: func(*cli.Context) error {
					config := input.ReadYAML(configFile).AsIOPath()
					generator.Build(config, noLog, noTime)
					return nil
				},
			},
			{
				Name:  "serve",
				Usage: "Start serving generated content via specific port (default: 8080)",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "port",
						Value:       8080,
						Aliases:     []string{"p"},
						Usage:       "Serve files on port `PORT`",
						Destination: &port,
					},
					&cli.StringFlag{
						Name:        "config",
						Value:       "config.yaml",
						Aliases:     []string{"c"},
						Usage:       "Load configuration from `FILE`.",
						Destination: &configFile,
					},
				},
				Action: func(*cli.Context) error {
					config := input.ReadYAML(configFile).AsIOPath()
					server.Serve(port, config.OutFolder)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
