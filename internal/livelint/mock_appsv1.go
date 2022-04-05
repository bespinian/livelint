// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package livelint

import (
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/rest"
	"sync"
)

// Ensure, that appsv1InterfaceMock does implement appsv1Interface.
// If this is not the case, regenerate this file with moq.
var _ appsv1Interface = &appsv1InterfaceMock{}

// appsv1InterfaceMock is a mock implementation of appsv1Interface.
//
// 	func TestSomethingThatUsesappsv1Interface(t *testing.T) {
//
// 		// make and configure a mocked appsv1Interface
// 		mockedappsv1Interface := &appsv1InterfaceMock{
// 			ControllerRevisionsFunc: func(namespace string) appsv1.ControllerRevisionInterface {
// 				panic("mock out the ControllerRevisions method")
// 			},
// 			DaemonSetsFunc: func(namespace string) appsv1.DaemonSetInterface {
// 				panic("mock out the DaemonSets method")
// 			},
// 			DeploymentsFunc: func(namespace string) appsv1.DeploymentInterface {
// 				panic("mock out the Deployments method")
// 			},
// 			RESTClientFunc: func() rest.Interface {
// 				panic("mock out the RESTClient method")
// 			},
// 			ReplicaSetsFunc: func(namespace string) appsv1.ReplicaSetInterface {
// 				panic("mock out the ReplicaSets method")
// 			},
// 			StatefulSetsFunc: func(namespace string) appsv1.StatefulSetInterface {
// 				panic("mock out the StatefulSets method")
// 			},
// 		}
//
// 		// use mockedappsv1Interface in code that requires appsv1Interface
// 		// and then make assertions.
//
// 	}
type appsv1InterfaceMock struct {
	// ControllerRevisionsFunc mocks the ControllerRevisions method.
	ControllerRevisionsFunc func(namespace string) appsv1.ControllerRevisionInterface

	// DaemonSetsFunc mocks the DaemonSets method.
	DaemonSetsFunc func(namespace string) appsv1.DaemonSetInterface

	// DeploymentsFunc mocks the Deployments method.
	DeploymentsFunc func(namespace string) appsv1.DeploymentInterface

	// RESTClientFunc mocks the RESTClient method.
	RESTClientFunc func() rest.Interface

	// ReplicaSetsFunc mocks the ReplicaSets method.
	ReplicaSetsFunc func(namespace string) appsv1.ReplicaSetInterface

	// StatefulSetsFunc mocks the StatefulSets method.
	StatefulSetsFunc func(namespace string) appsv1.StatefulSetInterface

	// calls tracks calls to the methods.
	calls struct {
		// ControllerRevisions holds details about calls to the ControllerRevisions method.
		ControllerRevisions []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// DaemonSets holds details about calls to the DaemonSets method.
		DaemonSets []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// Deployments holds details about calls to the Deployments method.
		Deployments []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// RESTClient holds details about calls to the RESTClient method.
		RESTClient []struct {
		}
		// ReplicaSets holds details about calls to the ReplicaSets method.
		ReplicaSets []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
		// StatefulSets holds details about calls to the StatefulSets method.
		StatefulSets []struct {
			// Namespace is the namespace argument value.
			Namespace string
		}
	}
	lockControllerRevisions sync.RWMutex
	lockDaemonSets          sync.RWMutex
	lockDeployments         sync.RWMutex
	lockRESTClient          sync.RWMutex
	lockReplicaSets         sync.RWMutex
	lockStatefulSets        sync.RWMutex
}

// ControllerRevisions calls ControllerRevisionsFunc.
func (mock *appsv1InterfaceMock) ControllerRevisions(namespace string) appsv1.ControllerRevisionInterface {
	if mock.ControllerRevisionsFunc == nil {
		panic("appsv1InterfaceMock.ControllerRevisionsFunc: method is nil but appsv1Interface.ControllerRevisions was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockControllerRevisions.Lock()
	mock.calls.ControllerRevisions = append(mock.calls.ControllerRevisions, callInfo)
	mock.lockControllerRevisions.Unlock()
	return mock.ControllerRevisionsFunc(namespace)
}

// ControllerRevisionsCalls gets all the calls that were made to ControllerRevisions.
// Check the length with:
//     len(mockedappsv1Interface.ControllerRevisionsCalls())
func (mock *appsv1InterfaceMock) ControllerRevisionsCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockControllerRevisions.RLock()
	calls = mock.calls.ControllerRevisions
	mock.lockControllerRevisions.RUnlock()
	return calls
}

// DaemonSets calls DaemonSetsFunc.
func (mock *appsv1InterfaceMock) DaemonSets(namespace string) appsv1.DaemonSetInterface {
	if mock.DaemonSetsFunc == nil {
		panic("appsv1InterfaceMock.DaemonSetsFunc: method is nil but appsv1Interface.DaemonSets was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockDaemonSets.Lock()
	mock.calls.DaemonSets = append(mock.calls.DaemonSets, callInfo)
	mock.lockDaemonSets.Unlock()
	return mock.DaemonSetsFunc(namespace)
}

// DaemonSetsCalls gets all the calls that were made to DaemonSets.
// Check the length with:
//     len(mockedappsv1Interface.DaemonSetsCalls())
func (mock *appsv1InterfaceMock) DaemonSetsCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockDaemonSets.RLock()
	calls = mock.calls.DaemonSets
	mock.lockDaemonSets.RUnlock()
	return calls
}

// Deployments calls DeploymentsFunc.
func (mock *appsv1InterfaceMock) Deployments(namespace string) appsv1.DeploymentInterface {
	if mock.DeploymentsFunc == nil {
		panic("appsv1InterfaceMock.DeploymentsFunc: method is nil but appsv1Interface.Deployments was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockDeployments.Lock()
	mock.calls.Deployments = append(mock.calls.Deployments, callInfo)
	mock.lockDeployments.Unlock()
	return mock.DeploymentsFunc(namespace)
}

// DeploymentsCalls gets all the calls that were made to Deployments.
// Check the length with:
//     len(mockedappsv1Interface.DeploymentsCalls())
func (mock *appsv1InterfaceMock) DeploymentsCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockDeployments.RLock()
	calls = mock.calls.Deployments
	mock.lockDeployments.RUnlock()
	return calls
}

// RESTClient calls RESTClientFunc.
func (mock *appsv1InterfaceMock) RESTClient() rest.Interface {
	if mock.RESTClientFunc == nil {
		panic("appsv1InterfaceMock.RESTClientFunc: method is nil but appsv1Interface.RESTClient was just called")
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
//     len(mockedappsv1Interface.RESTClientCalls())
func (mock *appsv1InterfaceMock) RESTClientCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRESTClient.RLock()
	calls = mock.calls.RESTClient
	mock.lockRESTClient.RUnlock()
	return calls
}

// ReplicaSets calls ReplicaSetsFunc.
func (mock *appsv1InterfaceMock) ReplicaSets(namespace string) appsv1.ReplicaSetInterface {
	if mock.ReplicaSetsFunc == nil {
		panic("appsv1InterfaceMock.ReplicaSetsFunc: method is nil but appsv1Interface.ReplicaSets was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockReplicaSets.Lock()
	mock.calls.ReplicaSets = append(mock.calls.ReplicaSets, callInfo)
	mock.lockReplicaSets.Unlock()
	return mock.ReplicaSetsFunc(namespace)
}

// ReplicaSetsCalls gets all the calls that were made to ReplicaSets.
// Check the length with:
//     len(mockedappsv1Interface.ReplicaSetsCalls())
func (mock *appsv1InterfaceMock) ReplicaSetsCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockReplicaSets.RLock()
	calls = mock.calls.ReplicaSets
	mock.lockReplicaSets.RUnlock()
	return calls
}

// StatefulSets calls StatefulSetsFunc.
func (mock *appsv1InterfaceMock) StatefulSets(namespace string) appsv1.StatefulSetInterface {
	if mock.StatefulSetsFunc == nil {
		panic("appsv1InterfaceMock.StatefulSetsFunc: method is nil but appsv1Interface.StatefulSets was just called")
	}
	callInfo := struct {
		Namespace string
	}{
		Namespace: namespace,
	}
	mock.lockStatefulSets.Lock()
	mock.calls.StatefulSets = append(mock.calls.StatefulSets, callInfo)
	mock.lockStatefulSets.Unlock()
	return mock.StatefulSetsFunc(namespace)
}

// StatefulSetsCalls gets all the calls that were made to StatefulSets.
// Check the length with:
//     len(mockedappsv1Interface.StatefulSetsCalls())
func (mock *appsv1InterfaceMock) StatefulSetsCalls() []struct {
	Namespace string
} {
	var calls []struct {
		Namespace string
	}
	mock.lockStatefulSets.RLock()
	calls = mock.calls.StatefulSets
	mock.lockStatefulSets.RUnlock()
	return calls
}
