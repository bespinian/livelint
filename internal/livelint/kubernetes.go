package livelint

import (
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	apiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	netv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

//go:generate moq -out mock_kubernetes_test.go -pkg livelint_test . KubernetesInterface
type KubernetesInterface = kubernetes.Interface

//go:generate moq -out mock_apiv1_test.go -pkg livelint_test . Apiv1Interface
type Apiv1Interface = apiv1.CoreV1Interface

//go:generate moq -out mock_apiv1_event_test.go -pkg livelint_test . Apiv1EventInterface
type Apiv1EventInterface = apiv1.EventInterface

//go:generate moq -out mock_apiv1_pvc_test.go -pkg livelint_test . Apiv1PVCInterface
type Apiv1PVCInterface = apiv1.PersistentVolumeClaimInterface

//go:generate moq -out mock_appsv1_test.go -pkg livelint_test . Appsv1Interface
type Appsv1Interface = appsv1.AppsV1Interface

//go:generate moq -out mock_appsv1_deployment_test.go -pkg livelint_test . DeploymentInterface
type DeploymentInterface = appsv1.DeploymentInterface

//go:generate moq -out mock_replicasets_test.go -pkg livelint_test . ReplicaSetInterface
type ReplicaSetInterface = appsv1.ReplicaSetInterface

//go:generate moq -out mock_apiv1_service_test.go -pkg livelint_test . Apiv1ServiceInterface
type Apiv1ServiceInterface = apiv1.ServiceInterface

//go:generate moq -out mock_apiv1_endpoints_test.go -pkg livelint_test . Apiv1EndpointsInterface
type Apiv1EndpointsInterface = apiv1.EndpointsInterface

//go:generate moq -out mock_apiv1_pod_test.go -pkg livelint_test . Apiv1PodInterface
type Apiv1PodInterface = apiv1.PodInterface

//go:generate moq -out mock_apiv1_node_test.go -pkg livelint_test . Apiv1NodeInterface
type Apiv1NodeInterface = apiv1.NodeInterface

//go:generate moq -out mock_netv1_test.go -pkg livelint_test . NetworkingV1Interface
type NetworkingV1Interface = netv1.NetworkingV1Interface

//go:generate moq -out mock_netv1_ingress_test.go -pkg livelint_test . IngressInterface
type IngressInterface = netv1.IngressInterface
