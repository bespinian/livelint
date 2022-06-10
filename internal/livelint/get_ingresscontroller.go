package livelint

import (
	"context"
	"errors"
	"fmt"
	"strings"

	apiv1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var errNotFound = errors.New("not found")

func (n *Livelint) getIngressControllerPods(ingress netv1.Ingress, ingressClasses map[string]netv1.IngressClass) ([]apiv1.Pod, error) {
	ingressControllerPods := []apiv1.Pod{}
	allPods, err := n.getPods("", make(map[string]string))
	if err != nil {
		return nil, fmt.Errorf("error getting all pods in cluster for finding ingress controller of ingress %s: %w", ingress.Name, err)
	}

	ingressClassName := getIngressClassName(ingress)
	var ingressClass netv1.IngressClass
	if ingressClassName == "" {
		defaultIngressClass, found := getDefaultIngressClass(ingressClasses)
		if !found {
			return nil, fmt.Errorf("no default ingress class found when searching controller for ingress %s: %w", ingress.Name, err)
		}
		ingressClass = defaultIngressClass
	} else {
		ingressClass = ingressClasses[ingressClassName]
	}

	for _, pod := range allPods {
		if isControllerPod(pod, ingressClass.Spec.Controller) {
			labelSelector, err := n.getOwningControllerLabelSelector(pod)
			if err != nil {
				return nil, fmt.Errorf("unable to get label selector for controller pods of ingress %s: %w", ingress.Name, err)
			}
			ingressControllerPods, err = n.getPods("", labelSelector)
			if err != nil {
				return nil, fmt.Errorf("unable to get controller pods of ingress %s: %w", ingress.Name, err)
			}
			break
		}
	}
	return ingressControllerPods, nil
}

func isControllerPod(pod apiv1.Pod, controllerName string) bool {
	for _, container := range pod.Spec.Containers {
		for _, arg := range container.Args {
			if strings.Contains(arg, controllerName) {
				return true
			}
		}
		for _, envVariable := range container.Env {
			if strings.Contains(envVariable.Value, controllerName) {
				return true
			}
		}
	}
	return false
}

func (n *Livelint) getOwningControllerLabelSelector(pod apiv1.Pod) (map[string]string, error) {
	var labelSelector map[string]string
	for _, ownerRef := range pod.OwnerReferences {
		if *ownerRef.Controller {
			switch ownerRef.Kind {
			case "ReplicaSet":
				replicaSet, err := n.k8s.AppsV1().ReplicaSets(pod.Namespace).Get(context.Background(), ownerRef.Name, metav1.GetOptions{})
				if err != nil {
					return nil, fmt.Errorf("error getting owning ReplicaSet's label selector for pod %s: %w", pod.Name, err)
				}
				labelSelector = replicaSet.Spec.Selector.MatchLabels
				return labelSelector, nil
			case "DaemonSet":
				daemonSet, err := n.k8s.AppsV1().DaemonSets(pod.Namespace).Get(context.Background(), ownerRef.Name, metav1.GetOptions{})
				if err != nil {
					return nil, fmt.Errorf("error getting owning DaemonSet's label selector for pod %s: %w", pod.Name, err)
				}
				labelSelector = daemonSet.Spec.Selector.MatchLabels
				return labelSelector, nil
			case "StatefulSet":
				statefulSet, err := n.k8s.AppsV1().StatefulSets(pod.Namespace).Get(context.Background(), ownerRef.Name, metav1.GetOptions{})
				if err != nil {
					return nil, fmt.Errorf("error getting owning StatefulSet's label selector for pod %s: %w", pod.Name, err)
				}
				labelSelector = statefulSet.Spec.Selector.MatchLabels
				return labelSelector, nil
			default:
				return nil, fmt.Errorf("could not get owning controller's label selector for pod %s: %w", pod.Name, errNotFound)
			}
		}
	}
	return nil, fmt.Errorf("could not get owning controller's label selector for pod %s: %w", pod.Name, errNotFound)
}
