// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package livelint

import (
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"sync"
)

// Ensure, that apiv1InterfaceMock does implement apiv1Interface.
// If this is not the case, regenerate this file with moq.
var _ apiv1Interface = &apiv1InterfaceMock{}

// apiv1InterfaceMock is a mock implementation of apiv1Interface.
//
// 	func TestSomethingThatUsesapiv1Interface(t *testing.T) {
//
// 		// make and configure a mocked apiv1Interface
// 		mockedapiv1Interface := &apiv1InterfaceMock{
// 			ComponentStatusesFunc: func() corev1.ComponentStatusInterface {
// 				panic("mock out the ComponentStatuses method")
// 			},
// 			ConfigMapsFunc: func(namespace string) corev1.ConfigMapInterface {
// 				panic("mock out the ConfigMaps method")
// 			},
// 			EndpointsFunc: func(namespace string) corev1.EndpointsInterface {
// 				panic("mock out the Endpoints method")
// 			},
// 			EventsFunc: func(namespace string) corev1.EventInterface {
// 				panic("mock out the Events method")
// 			},
// 			LimitRangesFunc: func(namespace string) corev1.LimitRangeInterface {
// 				panic("mock out the LimitRanges method")
// 			},
// 			NamespacesFunc: func() corev1.NamespaceInterface {
// 				panic("mock out the Namespaces method")
// 			},
// 			NodesFunc: func() corev1.NodeInterface {
// 				panic("mock out the Nodes method")
// 			},
// 			PersistentVolumeClaimsFunc: func(namespace string) corev1.PersistentVolumeClaimInterface {
// 				panic("mock out the PersistentVolumeClaims method")
// 			},
// 			PersistentVolumesFunc: func() corev1.PersistentVolumeInterface {
// 				panic("mock out the PersistentVolumes method")
// 			},
// 			PodTemplatesFunc: func(namespace string) corev1.PodTemplateInterface {
// 				panic("mock out the PodTemplates method")
// 			},
// 			PodsFunc: func(namespace string) corev1.PodInterface {
// 				panic("mock out the Pods method")
// 			},
// 			RESTClientFunc: func() rest.Interface {
// 				panic("mock out the RESTClient method")
// 			},
// 			ReplicationControllersFunc: func(namespace string) corev1.ReplicationControllerInterface {
// 				panic("mock out the ReplicationControllers method")
// 			},
// 			ResourceQuotasFunc: func(namespace string) corev1.ResourceQuotaInterface {
// 				panic("mock out the ResourceQuotas method")
// 			},
// 			SecretsFunc: func(namespace string) corev1.SecretInterface {
// 				panic("mock out the Secrets method")
// 			},
// 			ServiceAccountsFunc: func(namespace string) corev1.ServiceAccountInterface {
// 				panic("mock out the ServiceAccounts method")
// 			},
// 			ServicesFunc: func(namespace string) corev1.ServiceInterface {
// 				panic("mock out the Services method")
// 			},
// 		}
//
// 		// use mockedapiv1Interface in code that requires apiv1Interface
// 		// and then make assertions.
//
// 	}
type apiv1InterfaceMock struct {
	// ComponentStatusesFunc mocks the ComponentStatuses method.
	ComponentStatusesFunc func() corev1.ComponentStatusInterface

	// ConfigMapsFunc mocks the ConfigMaps method.
	ConfigMapsFunc func(namespace string) corev1.ConfigMapInterface

	// EndpointsFunc mocks the Endpoints method.
	EndpointsFunc func(namespace string) corev1.EndpointsInterface

	// EventsFunc mocks the Events method.
	EventsFunc func(namespace string) corev1.EventInterface

	// LimitRangesFunc mocks the LimitRanges method.
	LimitRangesFunc func(namespace string) corev1.LimitRangeInterface

	// NamespacesFunc mocks the Namespaces method.
	NamespacesFunc func() corev1.NamespaceInterface

	// NodesFunc mocks the Nodes method.
	NodesFunc func() corev1.NodeInterface

	// PersistentVolumeClaimsFunc mocks the PersistentVolumeClaims method.
	PersistentVolumeClaimsFunc func(namespace string) corev1.PersistentVolumeClaimInterface

	// PersistentVolumesFunc mocks the PersistentVolumes method.
	PersistentVolumesFunc func() corev1.PersistentVolumeInterface

	// PodTemplatesFunc mocks the PodTemplates method.
	PodTemplatesFunc func(namespace string) corev1.PodTemplateInterface

	// PodsFunc mocks the Pods method.
	PodsFunc func(namespace string) corev1.PodInterface

	// RESTClientFunc mocks the RESTClient method.
	RESTClientFunc func() rest.Interface

	// ReplicationControllersFunc mocks the ReplicationControllers method.
	ReplicationControllersFunc func(namespace string) corev1.ReplicationControllerInterface

	// ResourceQuotasFunc mocks the ResourceQuotas method.
	ResourceQuotasFunc func(namespace string) corev1.ResourceQuotaInterface

	// SecretsFunc mocks the Secrets method.
	SecretsFunc func(namespace string) corev1.SecretInterface

	// ServiceAccountsFunc mocks the ServiceAccounts method.
	ServiceAccountsFunc func(namespace string) corev1.ServiceAccountInterface

	// ServicesFunc mocks the Services method.
	ServicesFunc func(namespace string) corev1.ServiceInterface

	// calls tracks calls to the methods.
	calls struct {
		// ComponentStatuses holds details about calls to the ComponentStatuses method.
		ComponentStatuses []struct {
		}
		// ConfigMaps holds details about calls to the ConfigMaps method.
		ConfigMaps []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// Endpoints holds details about calls to the Endpoints method.
		Endpoints []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// Events holds details about calls to the Events method.
		Events []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// LimitRanges holds details about calls to the LimitRanges method.
		LimitRanges []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// Namespaces holds details about calls to the Namespaces method.
		Namespaces []struct {
		}
		// Nodes holds details about calls to the Nodes method.
		Nodes []struct {
		}
		// PersistentVolumeClaims holds details about calls to the PersistentVolumeClaims method.
		PersistentVolumeClaims []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// PersistentVolumes holds details about calls to the PersistentVolumes method.
		PersistentVolumes []struct {
		}
		// PodTemplates holds details about calls to the PodTemplates method.
		PodTemplates []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// Pods holds details about calls to the Pods method.
		Pods []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// RESTClient holds details about calls to the RESTClient method.
		RESTClient []struct {
		}
		// ReplicationControllers holds details about calls to the ReplicationControllers method.
		ReplicationControllers []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// ResourceQuotas holds details about calls to the ResourceQuotas method.
		ResourceQuotas []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// Secrets holds details about calls to the Secrets method.
		Secrets []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// ServiceAccounts holds details about calls to the ServiceAccounts method.
		ServiceAccounts []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// Services holds details about calls to the Services method.
		Services []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
	}
	lockComponentStatuses      sync.RWMutex
	lockConfigMaps             sync.RWMutex
	lockEndpoints              sync.RWMutex
	lockEvents                 sync.RWMutex
	lockLimitRanges            sync.RWMutex
	lockNamespaces             sync.RWMutex
	lockNodes                  sync.RWMutex
	lockPersistentVolumeClaims sync.RWMutex
	lockPersistentVolumes      sync.RWMutex
	lockPodTemplates           sync.RWMutex
	lockPods                   sync.RWMutex
	lockRESTClient             sync.RWMutex
	lockReplicationControllers sync.RWMutex
	lockResourceQuotas         sync.RWMutex
	lockSecrets                sync.RWMutex
	lockServiceAccounts        sync.RWMutex
	lockServices               sync.RWMutex
}

// ComponentStatuses calls ComponentStatusesFunc.
func (mock *apiv1InterfaceMock) ComponentStatuses() corev1.ComponentStatusInterface {
	if mock.ComponentStatusesFunc == nil {
		panic("apiv1InterfaceMock.ComponentStatusesFunc: method is nil but apiv1Interface.ComponentStatuses was just called")
	}
	callInfo := struct {
	}{}
	mock.lockComponentStatuses.Lock()
	mock.calls.ComponentStatuses = append(mock.calls.ComponentStatuses, callInfo)
	mock.lockComponentStatuses.Unlock()
	return mock.ComponentStatusesFunc()
}

// ComponentStatusesCalls gets all the calls that were made to ComponentStatuses.
// Check the length with:
//     len(mockedapiv1Interface.ComponentStatusesCalls())
func (mock *apiv1InterfaceMock) ComponentStatusesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockComponentStatuses.RLock()
	calls = mock.calls.ComponentStatuses
	mock.lockComponentStatuses.RUnlock()
	return calls
}

// ConfigMaps calls ConfigMapsFunc.
func (mock *apiv1InterfaceMock) ConfigMaps(namespace string) corev1.ConfigMapInterface {
	if mock.ConfigMapsFunc == nil {
		panic("apiv1InterfaceMock.ConfigMapsFunc: method is nil but apiv1Interface.ConfigMaps was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockConfigMaps.Lock()
	mock.calls.ConfigMaps = append(mock.calls.ConfigMaps, callInfo)
	mock.lockConfigMaps.Unlock()
	return mock.ConfigMapsFunc(namespace)
}

// ConfigMapsCalls gets all the calls that were made to ConfigMaps.
// Check the length with:
//     len(mockedapiv1Interface.ConfigMapsCalls())
func (mock *apiv1InterfaceMock) ConfigMapsCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockConfigMaps.RLock()
	calls = mock.calls.ConfigMaps
	mock.lockConfigMaps.RUnlock()
	return calls
}

// Endpoints calls EndpointsFunc.
func (mock *apiv1InterfaceMock) Endpoints(namespace string) corev1.EndpointsInterface {
	if mock.EndpointsFunc == nil {
		panic("apiv1InterfaceMock.EndpointsFunc: method is nil but apiv1Interface.Endpoints was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockEndpoints.Lock()
	mock.calls.Endpoints = append(mock.calls.Endpoints, callInfo)
	mock.lockEndpoints.Unlock()
	return mock.EndpointsFunc(namespace)
}

// EndpointsCalls gets all the calls that were made to Endpoints.
// Check the length with:
//     len(mockedapiv1Interface.EndpointsCalls())
func (mock *apiv1InterfaceMock) EndpointsCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockEndpoints.RLock()
	calls = mock.calls.Endpoints
	mock.lockEndpoints.RUnlock()
	return calls
}

// Events calls EventsFunc.
func (mock *apiv1InterfaceMock) Events(namespace string) corev1.EventInterface {
	if mock.EventsFunc == nil {
		panic("apiv1InterfaceMock.EventsFunc: method is nil but apiv1Interface.Events was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockEvents.Lock()
	mock.calls.Events = append(mock.calls.Events, callInfo)
	mock.lockEvents.Unlock()
	return mock.EventsFunc(namespace)
}

// EventsCalls gets all the calls that were made to Events.
// Check the length with:
//     len(mockedapiv1Interface.EventsCalls())
func (mock *apiv1InterfaceMock) EventsCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockEvents.RLock()
	calls = mock.calls.Events
	mock.lockEvents.RUnlock()
	return calls
}

// LimitRanges calls LimitRangesFunc.
func (mock *apiv1InterfaceMock) LimitRanges(namespace string) corev1.LimitRangeInterface {
	if mock.LimitRangesFunc == nil {
		panic("apiv1InterfaceMock.LimitRangesFunc: method is nil but apiv1Interface.LimitRanges was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockLimitRanges.Lock()
	mock.calls.LimitRanges = append(mock.calls.LimitRanges, callInfo)
	mock.lockLimitRanges.Unlock()
	return mock.LimitRangesFunc(namespace)
}

// LimitRangesCalls gets all the calls that were made to LimitRanges.
// Check the length with:
//     len(mockedapiv1Interface.LimitRangesCalls())
func (mock *apiv1InterfaceMock) LimitRangesCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockLimitRanges.RLock()
	calls = mock.calls.LimitRanges
	mock.lockLimitRanges.RUnlock()
	return calls
}

// Namespaces calls NamespacesFunc.
func (mock *apiv1InterfaceMock) Namespaces() corev1.NamespaceInterface {
	if mock.NamespacesFunc == nil {
		panic("apiv1InterfaceMock.NamespacesFunc: method is nil but apiv1Interface.Namespaces was just called")
	}
	callInfo := struct {
	}{}
	mock.lockNamespaces.Lock()
	mock.calls.Namespaces = append(mock.calls.Namespaces, callInfo)
	mock.lockNamespaces.Unlock()
	return mock.NamespacesFunc()
}

// NamespacesCalls gets all the calls that were made to Namespaces.
// Check the length with:
//     len(mockedapiv1Interface.NamespacesCalls())
func (mock *apiv1InterfaceMock) NamespacesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockNamespaces.RLock()
	calls = mock.calls.Namespaces
	mock.lockNamespaces.RUnlock()
	return calls
}

// Nodes calls NodesFunc.
func (mock *apiv1InterfaceMock) Nodes() corev1.NodeInterface {
	if mock.NodesFunc == nil {
		panic("apiv1InterfaceMock.NodesFunc: method is nil but apiv1Interface.Nodes was just called")
	}
	callInfo := struct {
	}{}
	mock.lockNodes.Lock()
	mock.calls.Nodes = append(mock.calls.Nodes, callInfo)
	mock.lockNodes.Unlock()
	return mock.NodesFunc()
}

// NodesCalls gets all the calls that were made to Nodes.
// Check the length with:
//     len(mockedapiv1Interface.NodesCalls())
func (mock *apiv1InterfaceMock) NodesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockNodes.RLock()
	calls = mock.calls.Nodes
	mock.lockNodes.RUnlock()
	return calls
}

// PersistentVolumeClaims calls PersistentVolumeClaimsFunc.
func (mock *apiv1InterfaceMock) PersistentVolumeClaims(namespace string) corev1.PersistentVolumeClaimInterface {
	if mock.PersistentVolumeClaimsFunc == nil {
		panic("apiv1InterfaceMock.PersistentVolumeClaimsFunc: method is nil but apiv1Interface.PersistentVolumeClaims was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockPersistentVolumeClaims.Lock()
	mock.calls.PersistentVolumeClaims = append(mock.calls.PersistentVolumeClaims, callInfo)
	mock.lockPersistentVolumeClaims.Unlock()
	return mock.PersistentVolumeClaimsFunc(namespace)
}

// PersistentVolumeClaimsCalls gets all the calls that were made to PersistentVolumeClaims.
// Check the length with:
//     len(mockedapiv1Interface.PersistentVolumeClaimsCalls())
func (mock *apiv1InterfaceMock) PersistentVolumeClaimsCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockPersistentVolumeClaims.RLock()
	calls = mock.calls.PersistentVolumeClaims
	mock.lockPersistentVolumeClaims.RUnlock()
	return calls
}

// PersistentVolumes calls PersistentVolumesFunc.
func (mock *apiv1InterfaceMock) PersistentVolumes() corev1.PersistentVolumeInterface {
	if mock.PersistentVolumesFunc == nil {
		panic("apiv1InterfaceMock.PersistentVolumesFunc: method is nil but apiv1Interface.PersistentVolumes was just called")
	}
	callInfo := struct {
	}{}
	mock.lockPersistentVolumes.Lock()
	mock.calls.PersistentVolumes = append(mock.calls.PersistentVolumes, callInfo)
	mock.lockPersistentVolumes.Unlock()
	return mock.PersistentVolumesFunc()
}

// PersistentVolumesCalls gets all the calls that were made to PersistentVolumes.
// Check the length with:
//     len(mockedapiv1Interface.PersistentVolumesCalls())
func (mock *apiv1InterfaceMock) PersistentVolumesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockPersistentVolumes.RLock()
	calls = mock.calls.PersistentVolumes
	mock.lockPersistentVolumes.RUnlock()
	return calls
}

// PodTemplates calls PodTemplatesFunc.
func (mock *apiv1InterfaceMock) PodTemplates(namespace string) corev1.PodTemplateInterface {
	if mock.PodTemplatesFunc == nil {
		panic("apiv1InterfaceMock.PodTemplatesFunc: method is nil but apiv1Interface.PodTemplates was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockPodTemplates.Lock()
	mock.calls.PodTemplates = append(mock.calls.PodTemplates, callInfo)
	mock.lockPodTemplates.Unlock()
	return mock.PodTemplatesFunc(namespace)
}

// PodTemplatesCalls gets all the calls that were made to PodTemplates.
// Check the length with:
//     len(mockedapiv1Interface.PodTemplatesCalls())
func (mock *apiv1InterfaceMock) PodTemplatesCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockPodTemplates.RLock()
	calls = mock.calls.PodTemplates
	mock.lockPodTemplates.RUnlock()
	return calls
}

// Pods calls PodsFunc.
func (mock *apiv1InterfaceMock) Pods(namespace string) corev1.PodInterface {
	if mock.PodsFunc == nil {
		panic("apiv1InterfaceMock.PodsFunc: method is nil but apiv1Interface.Pods was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockPods.Lock()
	mock.calls.Pods = append(mock.calls.Pods, callInfo)
	mock.lockPods.Unlock()
	return mock.PodsFunc(namespace)
}

// PodsCalls gets all the calls that were made to Pods.
// Check the length with:
//     len(mockedapiv1Interface.PodsCalls())
func (mock *apiv1InterfaceMock) PodsCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockPods.RLock()
	calls = mock.calls.Pods
	mock.lockPods.RUnlock()
	return calls
}

// RESTClient calls RESTClientFunc.
func (mock *apiv1InterfaceMock) RESTClient() rest.Interface {
	if mock.RESTClientFunc == nil {
		panic("apiv1InterfaceMock.RESTClientFunc: method is nil but apiv1Interface.RESTClient was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRESTClient.Lock()
	mock.calls.RESTClient = append(mock.calls.RESTClient, callInfo)
	mock.lockRESTClient.Unlock()
	return mock.RESTClientFunc()
}

// RESTClientCalls gets all the calls that were made to RESTClient.
// Check the length with:
//     len(mockedapiv1Interface.RESTClientCalls())
func (mock *apiv1InterfaceMock) RESTClientCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRESTClient.RLock()
	calls = mock.calls.RESTClient
	mock.lockRESTClient.RUnlock()
	return calls
}

// ReplicationControllers calls ReplicationControllersFunc.
func (mock *apiv1InterfaceMock) ReplicationControllers(namespace string) corev1.ReplicationControllerInterface {
	if mock.ReplicationControllersFunc == nil {
		panic("apiv1InterfaceMock.ReplicationControllersFunc: method is nil but apiv1Interface.ReplicationControllers was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockReplicationControllers.Lock()
	mock.calls.ReplicationControllers = append(mock.calls.ReplicationControllers, callInfo)
	mock.lockReplicationControllers.Unlock()
	return mock.ReplicationControllersFunc(namespace)
}

// ReplicationControllersCalls gets all the calls that were made to ReplicationControllers.
// Check the length with:
//     len(mockedapiv1Interface.ReplicationControllersCalls())
func (mock *apiv1InterfaceMock) ReplicationControllersCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockReplicationControllers.RLock()
	calls = mock.calls.ReplicationControllers
	mock.lockReplicationControllers.RUnlock()
	return calls
}

// ResourceQuotas calls ResourceQuotasFunc.
func (mock *apiv1InterfaceMock) ResourceQuotas(namespace string) corev1.ResourceQuotaInterface {
	if mock.ResourceQuotasFunc == nil {
		panic("apiv1InterfaceMock.ResourceQuotasFunc: method is nil but apiv1Interface.ResourceQuotas was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockResourceQuotas.Lock()
	mock.calls.ResourceQuotas = append(mock.calls.ResourceQuotas, callInfo)
	mock.lockResourceQuotas.Unlock()
	return mock.ResourceQuotasFunc(namespace)
}

// ResourceQuotasCalls gets all the calls that were made to ResourceQuotas.
// Check the length with:
//     len(mockedapiv1Interface.ResourceQuotasCalls())
func (mock *apiv1InterfaceMock) ResourceQuotasCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockResourceQuotas.RLock()
	calls = mock.calls.ResourceQuotas
	mock.lockResourceQuotas.RUnlock()
	return calls
}

// Secrets calls SecretsFunc.
func (mock *apiv1InterfaceMock) Secrets(namespace string) corev1.SecretInterface {
	if mock.SecretsFunc == nil {
		panic("apiv1InterfaceMock.SecretsFunc: method is nil but apiv1Interface.Secrets was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockSecrets.Lock()
	mock.calls.Secrets = append(mock.calls.Secrets, callInfo)
	mock.lockSecrets.Unlock()
	return mock.SecretsFunc(namespace)
}

// SecretsCalls gets all the calls that were made to Secrets.
// Check the length with:
//     len(mockedapiv1Interface.SecretsCalls())
func (mock *apiv1InterfaceMock) SecretsCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockSecrets.RLock()
	calls = mock.calls.Secrets
	mock.lockSecrets.RUnlock()
	return calls
}

// ServiceAccounts calls ServiceAccountsFunc.
func (mock *apiv1InterfaceMock) ServiceAccounts(namespace string) corev1.ServiceAccountInterface {
	if mock.ServiceAccountsFunc == nil {
		panic("apiv1InterfaceMock.ServiceAccountsFunc: method is nil but apiv1Interface.ServiceAccounts was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockServiceAccounts.Lock()
	mock.calls.ServiceAccounts = append(mock.calls.ServiceAccounts, callInfo)
	mock.lockServiceAccounts.Unlock()
	return mock.ServiceAccountsFunc(namespace)
}

// ServiceAccountsCalls gets all the calls that were made to ServiceAccounts.
// Check the length with:
//     len(mockedapiv1Interface.ServiceAccountsCalls())
func (mock *apiv1InterfaceMock) ServiceAccountsCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockServiceAccounts.RLock()
	calls = mock.calls.ServiceAccounts
	mock.lockServiceAccounts.RUnlock()
	return calls
}

// Services calls ServicesFunc.
func (mock *apiv1InterfaceMock) Services(namespace string) corev1.ServiceInterface {
	if mock.ServicesFunc == nil {
		panic("apiv1InterfaceMock.ServicesFunc: method is nil but apiv1Interface.Services was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockServices.Lock()
	mock.calls.Services = append(mock.calls.Services, callInfo)
	mock.lockServices.Unlock()
	return mock.ServicesFunc(namespace)
}

// ServicesCalls gets all the calls that were made to Services.
// Check the length with:
//     len(mockedapiv1Interface.ServicesCalls())
func (mock *apiv1InterfaceMock) ServicesCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockServices.RLock()
	calls = mock.calls.Services
	mock.lockServices.RUnlock()
	return calls
}
