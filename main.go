package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/bespinian/livelint/internal/livelint"
	"github.com/urfave/cli/v2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := os.Getenv(clientcmd.RecommendedConfigPathEnvVar)
	if kubeconfig == "" {
		home := homeDir()
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
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules,
		configOverrides)

	namespace, _, err := kubeConfig.Namespace()
	if err != nil {
		log.Fatal(fmt.Errorf("error creating getting namespace from k8s config: %w", err))
	}

	ll := livelint.New(k8s, config)

	app := &cli.App{
		Name:  "livelint",
		Usage: "debug k8s workload",

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
						Usage:   "if the tool should be verbose about what it's doing",
					},
				},

				Action: func(c *cli.Context) error {
					args := c.Args()
					go func() {
						ll.Start()
					}()
					runChecksErr := ll.RunChecks(c.String("namespace"), args.Get(0), c.Bool("verbose"))
					if runChecksErr != nil {
						log.Fatal(fmt.Errorf("error running checks: %w", runChecksErr))
					}
					time.Sleep(1 * time.Second)
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func homeDir() string {
	h := os.Getenv("HOME")
	if h == "" {
		h = os.Getenv("USERPROFILE") // windows
	}
	return h
}
