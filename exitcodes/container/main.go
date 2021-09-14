package main

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/bespinian/k8s-app-benchmarks/exitcodes/pkg/ebpf"
	"github.com/bespinian/k8s-app-benchmarks/exitcodes/pkg/exitcodes"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if errors.Is(err, rest.ErrNotInCluster) {
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	} else if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	k8s, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	e := exitcodes.New(k8s)
	ebpf.Run(e.HandleExec, e.HandleExit, e.HandleDone)
}
