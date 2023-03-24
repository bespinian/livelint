package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bespinian/livelint/internal/livelint"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var buildversion, builddate, githash string

var errNoHome = errors.New("error finding your HOME directory")

func main() {
	kubeconfig := os.Getenv(clientcmd.RecommendedConfigPathEnvVar)
	if kubeconfig == "" {
		home := homedir.HomeDir()
		if home == "" {
			exitWithErr(errNoHome)
		}
		kubeconfig = filepath.Join(home, clientcmd.RecommendedHomeDir, clientcmd.RecommendedFileName)
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		exitWithErr(fmt.Errorf("error building k8s config from flags: %w", err))
	}

	k8s, err := kubernetes.NewForConfig(config)
	if err != nil {
		exitWithErr(fmt.Errorf("error creating k8s client set: %w", err))
	}
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	namespace, _, err := kubeConfig.Namespace()
	if err != nil {
		exitWithErr(fmt.Errorf("error getting namespace from k8s config: %w", err))
	}

	bubbletea := livelint.NewBubbleTeaInterface(tea.NewProgram(livelint.InitialModel()))
	ll := livelint.New(k8s, config, bubbletea)

	app := &cli.App{
		Name:    "livelint",
		Usage:   "debug k8s workloads",
		Version: buildversion,

		Commands: []*cli.Command{
			{
				Name:    "check",
				Aliases: []string{"c"},
				Usage:   "checks a Kubernetes deployment for issues",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "namespace",
						Aliases: []string{"n"},
						Value:   namespace,
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
					args := c.Args()

					go func() {
						defer bubbletea.Quit()
						runErr := ll.RunChecks(c.String("namespace"), args.Get(0), c.Bool("verbose"))
						if runErr != nil {
							bubbletea.DisplayError(runErr)
						}
					}()

					err = bubbletea.Start()
					if err != nil {
						return fmt.Errorf("failed to start UI: %w", err)
					}

					return nil
				},
			},
		},

		Metadata: map[string]interface{}{
			"build-date": builddate,
			"git-hash":   githash,
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func exitWithErr(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "failed to start livelint: %s", err)
	os.Exit(1)
}
