package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bespinian/livelint/internal/livelint"
	"github.com/urfave/cli/v2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var buildversion, builddate, githash string

func main() {
	kubeconfig := os.Getenv(clientcmd.RecommendedConfigPathEnvVar)
	if kubeconfig == "" {
		home := homedir.HomeDir()
		if home == "" {
			log.Fatal("error finding your HOME directory")
		}
		kubeconfig = filepath.Join(home, clientcmd.RecommendedHomeDir, clientcmd.RecommendedFileName)
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(fmt.Errorf("error building k8s config from flags: %w", err))
	}

	k8s, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(fmt.Errorf("error creating k8s client set: %w", err))
	}
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	namespace, _, err := kubeConfig.Namespace()
	if err != nil {
		log.Fatal(fmt.Errorf("error getting namespace from k8s config: %w", err))
	}

	ll := livelint.New(k8s, config)

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
						err = ll.RunChecks(c.String("namespace"), args.Get(0), c.Bool("verbose"))
						if err != nil {
							log.Fatal(fmt.Errorf("error running checks: %w", err))
						}
						defer ll.Quit()
					}()

					ll.Start()

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
