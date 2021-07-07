package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bespinian/k8s-app-benchmarks/livelint/pkg/livelint"
	"github.com/urfave/cli/v2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	checkArgsCount = 2
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

	ll := livelint.New(k8s)

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
						Value:   "default",
						Usage:   "the source namespace",
					},
				},

				Action: func(c *cli.Context) error {
					args := c.Args()

					if args.Len() != checkArgsCount {
						return errors.New("usage: livelint check NAMESPACE DEPLOYMENT") //nolint:goerr113
					}
					err = ll.Check(c.String("namespace"), args.Get(0))
					if err != nil {
						return fmt.Errorf("error creating topo: %w", err)
					}

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
