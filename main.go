package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bespinian/livelint/internal/app"
	"github.com/bespinian/livelint/internal/livelint"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

var buildversion, builddate, githash string

func main() {
	cli := &cli.App{
		Name:    "livelint",
		Usage:   "debug k8s workloads",
		Version: buildversion,

		Commands: []*cli.Command{
			{
				Name:        "check",
				Aliases:     []string{"c"},
				Usage:       "livelint check/c [--namespace/-n <namespace>] <deployment name>",
				UsageText:   "livelint check my-deployment\nlivelint check --namespace default my-deployment",
				Description: "Checks a Kubernetes deployment for issues",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "namespace",
						Aliases: []string{"n"},
						Value:   "",
						Usage:   "the source namespace",
					},
					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Value:   false,
						Usage:   "if livelint should be verbose about what it's doing",
					},
				},

				Action: func(c *cli.Context) error {
					bubbletea := livelint.NewBubbleTeaInterface(tea.NewProgram(livelint.InitialModel()))
					namespace := c.String("namespace")
					app, err := app.New(namespace, c.Args().First(), bubbletea)
					if err != nil {
						if err := cli.ShowSubcommandHelp(c) != nil; err {
							log.Println(err)
						}
						exitWithErr(err)
					}

					go func() {
						defer bubbletea.Quit()
						err = app.Start()
						if err != nil {
							exitWithErr(err)
						}
					}()

					_, err = ui.Run()
					if err != nil {
						return fmt.Errorf("failed to start UI: %w", err)
					}

					return nil
				},

				OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
					if isSubcommand {
						if err := cli.ShowSubcommandHelp(c) != nil; err {
							log.Println(err)
						}
					} else {
						if err := cli.ShowAppHelp(c) != nil; err {
							log.Println(err)
						}
					}
					return err
				},
			},
		},

		Metadata: map[string]interface{}{
			"build-date": builddate,
			"git-hash":   githash,
		},
	}

	err := cli.Run(os.Args)
	if err != nil {
		exitWithErr(err)
	}
}

func exitWithErr(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "\nfailed to start livelint: %s\n", err)
	os.Exit(1)
}
