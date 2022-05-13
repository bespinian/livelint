package livelint

import (
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	apiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

//go:generate moq -out mock_kubernetes.go -pkg livelint . kubernetesInterface
type kubernetesInterface = kubernetes.Interface

//go:generate moq -out mock_apiv1.go -pkg livelint . apiv1Interface
type apiv1Interface = apiv1.CoreV1Interface

//go:generate moq -out mock_apiv1_event.go -pkg livelint . apiv1EventInterface
type apiv1EventInterface = apiv1.EventInterface

//go:generate moq -out mock_apiv1_pvc.go -pkg livelint . apiv1PVCInterface
type apiv1PVCInterface = apiv1.PersistentVolumeClaimInterface

//go:generate moq -out mock_appsv1.go -pkg livelint . appsv1Interface
type appsv1Interface = appsv1.AppsV1Interface

//go:generate moq -out mock_replicasets.go -pkg livelint . replicaSetInterface
type replicaSetInterface = appsv1.ReplicaSetInterface

//go:generate moq -out mock_apiv1_service.go -pkg livelint . apiv1ServiceInterface
type apiv1ServiceInterface = apiv1.ServiceInterface

//go:generate moq -out mock_apiv1_endpoints.go -pkg livelint . apiv1EndpointsInterface
type apiv1EndpointsInterface = apiv1.EndpointsInterface

//go:generate moq -out mock_apiv1_pod.go -pkg livelint . apiv1PodInterface
type apiv1PodInterface = apiv1.PodInterface

//go:generate moq -out mock_apiv1_node.go -pkg livelint . apiv1NodeInterface
type apiv1NodeInterface = apiv1.NodeInterface
