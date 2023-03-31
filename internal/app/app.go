package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"github.com/bespinian/livelint/internal/livelint"
)

type App struct {
	k8s            *kubernetes.Clientset
	config         *rest.Config
	namespace      string
	ui             livelint.UserInteraction
	deploymentName string
}

var (
	errNoHome           = errors.New("error finding your HOME directory")
	errNoDeploymentName = errors.New("deployment name is missing")
)

func New(namespace string, deploymentName string, ui livelint.UserInteraction) (*App, error) {
	if deploymentName == "" {
		return nil, errNoDeploymentName
	}

	app := App{
		namespace:      namespace,
		ui:             ui,
		deploymentName: deploymentName,
	}

	kubeConfigPath, err := getKubecConfigPath()
	if err != nil {
		return nil, err
	}

	if app.namespace == "" {
		loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
		kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, nil)
		kubeConfigNamespace, _, err := kubeConfig.Namespace()
		if err != nil {
			return nil, fmt.Errorf("error getting namespace from k8s config: %w", err)
		}
		app.namespace = kubeConfigNamespace
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error building k8s config from flags: %w", err)
	}
	app.config = config

	k8s, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error creating k8s client set: %w", err)
	}
	app.k8s = k8s

	// bubbletea := livelint.NewBubbleTeaInterface(tea.NewProgram(livelint.InitialModel()))
	// ll := livelint.New(k8s, config, bubbletea)

	return &app, nil
}

func (app *App) Start() error {
	ll := livelint.New(app.k8s, app.config, app.ui)

	runErr := ll.RunChecks(app.namespace, app.deploymentName)
	if runErr != nil {
		app.ui.DisplayError(runErr)
	}
	return nil
}

func getKubecConfigPath() (string, error) {
	kubeconfig := os.Getenv(clientcmd.RecommendedConfigPathEnvVar)
	if kubeconfig == "" {
		home := homedir.HomeDir()
		if home == "" {
			return "", errNoHome
		}
		kubeconfig = filepath.Join(home, clientcmd.RecommendedHomeDir, clientcmd.RecommendedFileName)
	}

	return kubeconfig, nil
}
