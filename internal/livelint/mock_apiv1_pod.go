// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package livelint

import (
	"context"
	apiv1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	"k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/rest"
	"sync"
)

// Ensure, that apiv1PodInterfaceMock does implement apiv1PodInterface.
// If this is not the case, regenerate this file with moq.
var _ apiv1PodInterface = &apiv1PodInterfaceMock{}

// apiv1PodInterfaceMock is a mock implementation of apiv1PodInterface.
//
// 	func TestSomethingThatUsesapiv1PodInterface(t *testing.T) {
//
// 		// make and configure a mocked apiv1PodInterface
// 		mockedapiv1PodInterface := &apiv1PodInterfaceMock{
// 			ApplyFunc: func(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (*apiv1.Pod, error) {
// 				panic("mock out the Apply method")
// 			},
// 			ApplyStatusFunc: func(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (*apiv1.Pod, error) {
// 				panic("mock out the ApplyStatus method")
// 			},
// 			BindFunc: func(ctx context.Context, binding *apiv1.Binding, opts metav1.CreateOptions) error {
// 				panic("mock out the Bind method")
// 			},
// 			CreateFunc: func(ctx context.Context, pod *apiv1.Pod, opts metav1.CreateOptions) (*apiv1.Pod, error) {
// 				panic("mock out the Create method")
// 			},
// 			DeleteFunc: func(ctx context.Context, name string, opts metav1.DeleteOptions) error {
// 				panic("mock out the Delete method")
// 			},
// 			DeleteCollectionFunc: func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
// 				panic("mock out the DeleteCollection method")
// 			},
// 			EvictFunc: func(ctx context.Context, eviction *v1beta1.Eviction) error {
// 				panic("mock out the Evict method")
// 			},
// 			EvictV1Func: func(ctx context.Context, eviction *policyv1.Eviction) error {
// 				panic("mock out the EvictV1 method")
// 			},
// 			EvictV1beta1Func: func(ctx context.Context, eviction *v1beta1.Eviction) error {
// 				panic("mock out the EvictV1beta1 method")
// 			},
// 			GetFunc: func(ctx context.Context, name string, opts metav1.GetOptions) (*apiv1.Pod, error) {
// 				panic("mock out the Get method")
// 			},
// 			GetLogsFunc: func(name string, opts *apiv1.PodLogOptions) *rest.Request {
// 				panic("mock out the GetLogs method")
// 			},
// 			ListFunc: func(ctx context.Context, opts metav1.ListOptions) (*apiv1.PodList, error) {
// 				panic("mock out the List method")
// 			},
// 			PatchFunc: func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*apiv1.Pod, error) {
// 				panic("mock out the Patch method")
// 			},
// 			ProxyGetFunc: func(scheme string, name string, port string, path string, params map[string]string) rest.ResponseWrapper {
// 				panic("mock out the ProxyGet method")
// 			},
// 			UpdateFunc: func(ctx context.Context, pod *apiv1.Pod, opts metav1.UpdateOptions) (*apiv1.Pod, error) {
// 				panic("mock out the Update method")
// 			},
// 			UpdateEphemeralContainersFunc: func(ctx context.Context, podName string, pod *apiv1.Pod, opts metav1.UpdateOptions) (*apiv1.Pod, error) {
// 				panic("mock out the UpdateEphemeralContainers method")
// 			},
// 			UpdateStatusFunc: func(ctx context.Context, pod *apiv1.Pod, opts metav1.UpdateOptions) (*apiv1.Pod, error) {
// 				panic("mock out the UpdateStatus method")
// 			},
// 			WatchFunc: func(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
// 				panic("mock out the Watch method")
// 			},
// 		}
//
// 		// use mockedapiv1PodInterface in code that requires apiv1PodInterface
// 		// and then make assertions.
//
// 	}
type apiv1PodInterfaceMock struct {
	// ApplyFunc mocks the Apply method.
	ApplyFunc func(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (*apiv1.Pod, error)

	// ApplyStatusFunc mocks the ApplyStatus method.
	ApplyStatusFunc func(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (*apiv1.Pod, error)

	// BindFunc mocks the Bind method.
	BindFunc func(ctx context.Context, binding *apiv1.Binding, opts metav1.CreateOptions) error

	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, pod *apiv1.Pod, opts metav1.CreateOptions) (*apiv1.Pod, error)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, name string, opts metav1.DeleteOptions) error

	// DeleteCollectionFunc mocks the DeleteCollection method.
	DeleteCollectionFunc func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error

	// EvictFunc mocks the Evict method.
	EvictFunc func(ctx context.Context, eviction *v1beta1.Eviction) error

	// EvictV1Func mocks the EvictV1 method.
	EvictV1Func func(ctx context.Context, eviction *policyv1.Eviction) error

	// EvictV1beta1Func mocks the EvictV1beta1 method.
	EvictV1beta1Func func(ctx context.Context, eviction *v1beta1.Eviction) error

	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, name string, opts metav1.GetOptions) (*apiv1.Pod, error)

	// GetLogsFunc mocks the GetLogs method.
	GetLogsFunc func(name string, opts *apiv1.PodLogOptions) *rest.Request

	// ListFunc mocks the List method.
	ListFunc func(ctx context.Context, opts metav1.ListOptions) (*apiv1.PodList, error)

	// PatchFunc mocks the Patch method.
	PatchFunc func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*apiv1.Pod, error)

	// ProxyGetFunc mocks the ProxyGet method.
	ProxyGetFunc func(scheme string, name string, port string, path string, params map[string]string) rest.ResponseWrapper

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, pod *apiv1.Pod, opts metav1.UpdateOptions) (*apiv1.Pod, error)

	// UpdateEphemeralContainersFunc mocks the UpdateEphemeralContainers method.
	UpdateEphemeralContainersFunc func(ctx context.Context, podName string, pod *apiv1.Pod, opts metav1.UpdateOptions) (*apiv1.Pod, error)

	// UpdateStatusFunc mocks the UpdateStatus method.
	UpdateStatusFunc func(ctx context.Context, pod *apiv1.Pod, opts metav1.UpdateOptions) (*apiv1.Pod, error)

	// WatchFunc mocks the Watch method.
	WatchFunc func(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)

	// calls tracks calls to the methods.
	calls struct {
		// Apply holds details about calls to the Apply method.
		Apply []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Pod is the pod argument value.
			Pod *corev1.PodApplyConfiguration
			// Opts is the opts argument value.
			Opts metav1.ApplyOptions
		}
		// ApplyStatus holds details about calls to the ApplyStatus method.
		ApplyStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Pod is the pod argument value.
			Pod *corev1.PodApplyConfiguration
			// Opts is the opts argument value.
			Opts metav1.ApplyOptions
		}
		// Bind holds details about calls to the Bind method.
		Bind []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Binding is the binding argument value.
			Binding *apiv1.Binding
			// Opts is the opts argument value.
			Opts metav1.CreateOptions
		}
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Pod is the pod argument value.
			Pod *apiv1.Pod
			// Opts is the opts argument value.
			Opts metav1.CreateOptions
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// Opts is the opts argument value.
			Opts metav1.DeleteOptions
		}
		// DeleteCollection holds details about calls to the DeleteCollection method.
		DeleteCollection []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts metav1.DeleteOptions
			// ListOpts is the listOpts argument value.
			ListOpts metav1.ListOptions
		}
		// Evict holds details about calls to the Evict method.
		Evict []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Eviction is the eviction argument value.
			Eviction *v1beta1.Eviction
		}
		// EvictV1 holds details about calls to the EvictV1 method.
		EvictV1 []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Eviction is the eviction argument value.
			Eviction *policyv1.Eviction
		}
		// EvictV1beta1 holds details about calls to the EvictV1beta1 method.
		EvictV1beta1 []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Eviction is the eviction argument value.
			Eviction *v1beta1.Eviction
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// Opts is the opts argument value.
			Opts metav1.GetOptions
		}
		// GetLogs holds details about calls to the GetLogs method.
		GetLogs []struct {
			// Name is the name argument value.
			Name string
			// Opts is the opts argument value.
			Opts *apiv1.PodLogOptions
		}
		// List holds details about calls to the List method.
		List []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts metav1.ListOptions
		}
		// Patch holds details about calls to the Patch method.
		Patch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// Pt is the pt argument value.
			Pt types.PatchType
			// Data is the data argument value.
			Data []byte
			// Opts is the opts argument value.
			Opts metav1.PatchOptions
			// Subresources is the subresources argument value.
			Subresources []string
		}
		// ProxyGet holds details about calls to the ProxyGet method.
		ProxyGet []struct {
			// Scheme is the scheme argument value.
			Scheme string
			// Name is the name argument value.
			Name string
			// Port is the port argument value.
			Port string
			// Path is the path argument value.
			Path string
			// Params is the params argument value.
			Params map[string]string
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Pod is the pod argument value.
			Pod *apiv1.Pod
			// Opts is the opts argument value.
			Opts metav1.UpdateOptions
		}
		// UpdateEphemeralContainers holds details about calls to the UpdateEphemeralContainers method.
		UpdateEphemeralContainers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// PodName is the podName argument value.
			PodName string
			// Pod is the pod argument value.
			Pod *apiv1.Pod
			// Opts is the opts argument value.
			Opts metav1.UpdateOptions
		}
		// UpdateStatus holds details about calls to the UpdateStatus method.
		UpdateStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Pod is the pod argument value.
			Pod *apiv1.Pod
			// Opts is the opts argument value.
			Opts metav1.UpdateOptions
		}
		// Watch holds details about calls to the Watch method.
		Watch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts metav1.ListOptions
		}
	}
	lockApply                     sync.RWMutex
	lockApplyStatus               sync.RWMutex
	lockBind                      sync.RWMutex
	lockCreate                    sync.RWMutex
	lockDelete                    sync.RWMutex
	lockDeleteCollection          sync.RWMutex
	lockEvict                     sync.RWMutex
	lockEvictV1                   sync.RWMutex
	lockEvictV1beta1              sync.RWMutex
	lockGet                       sync.RWMutex
	lockGetLogs                   sync.RWMutex
	lockList                      sync.RWMutex
	lockPatch                     sync.RWMutex
	lockProxyGet                  sync.RWMutex
	lockUpdate                    sync.RWMutex
	lockUpdateEphemeralContainers sync.RWMutex
	lockUpdateStatus              sync.RWMutex
	lockWatch                     sync.RWMutex
}

// Apply calls ApplyFunc.
func (mock *apiv1PodInterfaceMock) Apply(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (*apiv1.Pod, error) {
	if mock.ApplyFunc == nil {
		panic("apiv1PodInterfaceMock.ApplyFunc: method is nil but apiv1PodInterface.Apply was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Pod  *corev1.PodApplyConfiguration
		Opts metav1.ApplyOptions
	}{
		Ctx:  ctx,
		Pod:  pod,
		Opts: opts,
	}
	mock.lockApply.Lock()
	mock.calls.Apply = append(mock.calls.Apply, callInfo)
	mock.lockApply.Unlock()
	return mock.ApplyFunc(ctx, pod, opts)
}

// ApplyCalls gets all the calls that were made to Apply.
// Check the length with:
//     len(mockedapiv1PodInterface.ApplyCalls())
func (mock *apiv1PodInterfaceMock) ApplyCalls() []struct {
	Ctx  context.Context
	Pod  *corev1.PodApplyConfiguration
	Opts metav1.ApplyOptions
} {
	var calls []struct {
		Ctx  context.Context
		Pod  *corev1.PodApplyConfiguration
		Opts metav1.ApplyOptions
	}
	mock.lockApply.RLock()
	calls = mock.calls.Apply
	mock.lockApply.RUnlock()
	return calls
}

// ApplyStatus calls ApplyStatusFunc.
func (mock *apiv1PodInterfaceMock) ApplyStatus(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (*apiv1.Pod, error) {
	if mock.ApplyStatusFunc == nil {
		panic("apiv1PodInterfaceMock.ApplyStatusFunc: method is nil but apiv1PodInterface.ApplyStatus was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Pod  *corev1.PodApplyConfiguration
		Opts metav1.ApplyOptions
	}{
		Ctx:  ctx,
		Pod:  pod,
		Opts: opts,
	}
	mock.lockApplyStatus.Lock()
	mock.calls.ApplyStatus = append(mock.calls.ApplyStatus, callInfo)
	mock.lockApplyStatus.Unlock()
	return mock.ApplyStatusFunc(ctx, pod, opts)
}

// ApplyStatusCalls gets all the calls that were made to ApplyStatus.
// Check the length with:
//     len(mockedapiv1PodInterface.ApplyStatusCalls())
func (mock *apiv1PodInterfaceMock) ApplyStatusCalls() []struct {
	Ctx  context.Context
	Pod  *corev1.PodApplyConfiguration
	Opts metav1.ApplyOptions
} {
	var calls []struct {
		Ctx  context.Context
		Pod  *corev1.PodApplyConfiguration
		Opts metav1.ApplyOptions
	}
	mock.lockApplyStatus.RLock()
	calls = mock.calls.ApplyStatus
	mock.lockApplyStatus.RUnlock()
	return calls
}

// Bind calls BindFunc.
func (mock *apiv1PodInterfaceMock) Bind(ctx context.Context, binding *apiv1.Binding, opts metav1.CreateOptions) error {
	if mock.BindFunc == nil {
		panic("apiv1PodInterfaceMock.BindFunc: method is nil but apiv1PodInterface.Bind was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Binding *apiv1.Binding
		Opts    metav1.CreateOptions
	}{
		Ctx:     ctx,
		Binding: binding,
		Opts:    opts,
	}
	mock.lockBind.Lock()
	mock.calls.Bind = append(mock.calls.Bind, callInfo)
	mock.lockBind.Unlock()
	return mock.BindFunc(ctx, binding, opts)
}

// BindCalls gets all the calls that were made to Bind.
// Check the length with:
//     len(mockedapiv1PodInterface.BindCalls())
func (mock *apiv1PodInterfaceMock) BindCalls() []struct {
	Ctx     context.Context
	Binding *apiv1.Binding
	Opts    metav1.CreateOptions
} {
	var calls []struct {
		Ctx     context.Context
		Binding *apiv1.Binding
		Opts    metav1.CreateOptions
	}
	mock.lockBind.RLock()
	calls = mock.calls.Bind
	mock.lockBind.RUnlock()
	return calls
}

// Create calls CreateFunc.
func (mock *apiv1PodInterfaceMock) Create(ctx context.Context, pod *apiv1.Pod, opts metav1.CreateOptions) (*apiv1.Pod, error) {
	if mock.CreateFunc == nil {
		panic("apiv1PodInterfaceMock.CreateFunc: method is nil but apiv1PodInterface.Create was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Pod  *apiv1.Pod
		Opts metav1.CreateOptions
	}{
		Ctx:  ctx,
		Pod:  pod,
		Opts: opts,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, pod, opts)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedapiv1PodInterface.CreateCalls())
func (mock *apiv1PodInterfaceMock) CreateCalls() []struct {
	Ctx  context.Context
	Pod  *apiv1.Pod
	Opts metav1.CreateOptions
} {
	var calls []struct {
		Ctx  context.Context
		Pod  *apiv1.Pod
		Opts metav1.CreateOptions
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *apiv1PodInterfaceMock) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	if mock.DeleteFunc == nil {
		panic("apiv1PodInterfaceMock.DeleteFunc: method is nil but apiv1PodInterface.Delete was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
		Opts metav1.DeleteOptions
	}{
		Ctx:  ctx,
		Name: name,
		Opts: opts,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(ctx, name, opts)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedapiv1PodInterface.DeleteCalls())
func (mock *apiv1PodInterfaceMock) DeleteCalls() []struct {
	Ctx  context.Context
	Name string
	Opts metav1.DeleteOptions
} {
	var calls []struct {
		Ctx  context.Context
		Name string
		Opts metav1.DeleteOptions
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// DeleteCollection calls DeleteCollectionFunc.
func (mock *apiv1PodInterfaceMock) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	if mock.DeleteCollectionFunc == nil {
		panic("apiv1PodInterfaceMock.DeleteCollectionFunc: method is nil but apiv1PodInterface.DeleteCollection was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Opts     metav1.DeleteOptions
		ListOpts metav1.ListOptions
	}{
		Ctx:      ctx,
		Opts:     opts,
		ListOpts: listOpts,
	}
	mock.lockDeleteCollection.Lock()
	mock.calls.DeleteCollection = append(mock.calls.DeleteCollection, callInfo)
	mock.lockDeleteCollection.Unlock()
	return mock.DeleteCollectionFunc(ctx, opts, listOpts)
}

// DeleteCollectionCalls gets all the calls that were made to DeleteCollection.
// Check the length with:
//     len(mockedapiv1PodInterface.DeleteCollectionCalls())
func (mock *apiv1PodInterfaceMock) DeleteCollectionCalls() []struct {
	Ctx      context.Context
	Opts     metav1.DeleteOptions
	ListOpts metav1.ListOptions
} {
	var calls []struct {
		Ctx      context.Context
		Opts     metav1.DeleteOptions
		ListOpts metav1.ListOptions
	}
	mock.lockDeleteCollection.RLock()
	calls = mock.calls.DeleteCollection
	mock.lockDeleteCollection.RUnlock()
	return calls
}

// Evict calls EvictFunc.
func (mock *apiv1PodInterfaceMock) Evict(ctx context.Context, eviction *v1beta1.Eviction) error {
	if mock.EvictFunc == nil {
		panic("apiv1PodInterfaceMock.EvictFunc: method is nil but apiv1PodInterface.Evict was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Eviction *v1beta1.Eviction
	}{
		Ctx:      ctx,
		Eviction: eviction,
	}
	mock.lockEvict.Lock()
	mock.calls.Evict = append(mock.calls.Evict, callInfo)
	mock.lockEvict.Unlock()
	return mock.EvictFunc(ctx, eviction)
}

// EvictCalls gets all the calls that were made to Evict.
// Check the length with:
//     len(mockedapiv1PodInterface.EvictCalls())
func (mock *apiv1PodInterfaceMock) EvictCalls() []struct {
	Ctx      context.Context
	Eviction *v1beta1.Eviction
} {
	var calls []struct {
		Ctx      context.Context
		Eviction *v1beta1.Eviction
	}
	mock.lockEvict.RLock()
	calls = mock.calls.Evict
	mock.lockEvict.RUnlock()
	return calls
}

// EvictV1 calls EvictV1Func.
func (mock *apiv1PodInterfaceMock) EvictV1(ctx context.Context, eviction *policyv1.Eviction) error {
	if mock.EvictV1Func == nil {
		panic("apiv1PodInterfaceMock.EvictV1Func: method is nil but apiv1PodInterface.EvictV1 was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Eviction *policyv1.Eviction
	}{
		Ctx:      ctx,
		Eviction: eviction,
	}
	mock.lockEvictV1.Lock()
	mock.calls.EvictV1 = append(mock.calls.EvictV1, callInfo)
	mock.lockEvictV1.Unlock()
	return mock.EvictV1Func(ctx, eviction)
}

// EvictV1Calls gets all the calls that were made to EvictV1.
// Check the length with:
//     len(mockedapiv1PodInterface.EvictV1Calls())
func (mock *apiv1PodInterfaceMock) EvictV1Calls() []struct {
	Ctx      context.Context
	Eviction *policyv1.Eviction
} {
	var calls []struct {
		Ctx      context.Context
		Eviction *policyv1.Eviction
	}
	mock.lockEvictV1.RLock()
	calls = mock.calls.EvictV1
	mock.lockEvictV1.RUnlock()
	return calls
}

// EvictV1beta1 calls EvictV1beta1Func.
func (mock *apiv1PodInterfaceMock) EvictV1beta1(ctx context.Context, eviction *v1beta1.Eviction) error {
	if mock.EvictV1beta1Func == nil {
		panic("apiv1PodInterfaceMock.EvictV1beta1Func: method is nil but apiv1PodInterface.EvictV1beta1 was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Eviction *v1beta1.Eviction
	}{
		Ctx:      ctx,
		Eviction: eviction,
	}
	mock.lockEvictV1beta1.Lock()
	mock.calls.EvictV1beta1 = append(mock.calls.EvictV1beta1, callInfo)
	mock.lockEvictV1beta1.Unlock()
	return mock.EvictV1beta1Func(ctx, eviction)
}

// EvictV1beta1Calls gets all the calls that were made to EvictV1beta1.
// Check the length with:
//     len(mockedapiv1PodInterface.EvictV1beta1Calls())
func (mock *apiv1PodInterfaceMock) EvictV1beta1Calls() []struct {
	Ctx      context.Context
	Eviction *v1beta1.Eviction
} {
	var calls []struct {
		Ctx      context.Context
		Eviction *v1beta1.Eviction
	}
	mock.lockEvictV1beta1.RLock()
	calls = mock.calls.EvictV1beta1
	mock.lockEvictV1beta1.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *apiv1PodInterfaceMock) Get(ctx context.Context, name string, opts metav1.GetOptions) (*apiv1.Pod, error) {
	if mock.GetFunc == nil {
		panic("apiv1PodInterfaceMock.GetFunc: method is nil but apiv1PodInterface.Get was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
		Opts metav1.GetOptions
	}{
		Ctx:  ctx,
		Name: name,
		Opts: opts,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(ctx, name, opts)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedapiv1PodInterface.GetCalls())
func (mock *apiv1PodInterfaceMock) GetCalls() []struct {
	Ctx  context.Context
	Name string
	Opts metav1.GetOptions
} {
	var calls []struct {
		Ctx  context.Context
		Name string
		Opts metav1.GetOptions
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// GetLogs calls GetLogsFunc.
func (mock *apiv1PodInterfaceMock) GetLogs(name string, opts *apiv1.PodLogOptions) *rest.Request {
	if mock.GetLogsFunc == nil {
		panic("apiv1PodInterfaceMock.GetLogsFunc: method is nil but apiv1PodInterface.GetLogs was just called")
	}
	callInfo := struct {
		Name string
		Opts *apiv1.PodLogOptions
	}{
		Name: name,
		Opts: opts,
	}
	mock.lockGetLogs.Lock()
	mock.calls.GetLogs = append(mock.calls.GetLogs, callInfo)
	mock.lockGetLogs.Unlock()
	return mock.GetLogsFunc(name, opts)
}

// GetLogsCalls gets all the calls that were made to GetLogs.
// Check the length with:
//     len(mockedapiv1PodInterface.GetLogsCalls())
func (mock *apiv1PodInterfaceMock) GetLogsCalls() []struct {
	Name string
	Opts *apiv1.PodLogOptions
} {
	var calls []struct {
		Name string
		Opts *apiv1.PodLogOptions
	}
	mock.lockGetLogs.RLock()
	calls = mock.calls.GetLogs
	mock.lockGetLogs.RUnlock()
	return calls
}

// List calls ListFunc.
func (mock *apiv1PodInterfaceMock) List(ctx context.Context, opts metav1.ListOptions) (*apiv1.PodList, error) {
	if mock.ListFunc == nil {
		panic("apiv1PodInterfaceMock.ListFunc: method is nil but apiv1PodInterface.List was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Opts metav1.ListOptions
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc(ctx, opts)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//     len(mockedapiv1PodInterface.ListCalls())
func (mock *apiv1PodInterfaceMock) ListCalls() []struct {
	Ctx  context.Context
	Opts metav1.ListOptions
} {
	var calls []struct {
		Ctx  context.Context
		Opts metav1.ListOptions
	}
	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}

// Patch calls PatchFunc.
func (mock *apiv1PodInterfaceMock) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*apiv1.Pod, error) {
	if mock.PatchFunc == nil {
		panic("apiv1PodInterfaceMock.PatchFunc: method is nil but apiv1PodInterface.Patch was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		Name         string
		Pt           types.PatchType
		Data         []byte
		Opts         metav1.PatchOptions
		Subresources []string
	}{
		Ctx:          ctx,
		Name:         name,
		Pt:           pt,
		Data:         data,
		Opts:         opts,
		Subresources: subresources,
	}
	mock.lockPatch.Lock()
	mock.calls.Patch = append(mock.calls.Patch, callInfo)
	mock.lockPatch.Unlock()
	return mock.PatchFunc(ctx, name, pt, data, opts, subresources...)
}

// PatchCalls gets all the calls that were made to Patch.
// Check the length with:
//     len(mockedapiv1PodInterface.PatchCalls())
func (mock *apiv1PodInterfaceMock) PatchCalls() []struct {
	Ctx          context.Context
	Name         string
	Pt           types.PatchType
	Data         []byte
	Opts         metav1.PatchOptions
	Subresources []string
} {
	var calls []struct {
		Ctx          context.Context
		Name         string
		Pt           types.PatchType
		Data         []byte
		Opts         metav1.PatchOptions
		Subresources []string
	}
	mock.lockPatch.RLock()
	calls = mock.calls.Patch
	mock.lockPatch.RUnlock()
	return calls
}

// ProxyGet calls ProxyGetFunc.
func (mock *apiv1PodInterfaceMock) ProxyGet(scheme string, name string, port string, path string, params map[string]string) rest.ResponseWrapper {
	if mock.ProxyGetFunc == nil {
		panic("apiv1PodInterfaceMock.ProxyGetFunc: method is nil but apiv1PodInterface.ProxyGet was just called")
	}
	callInfo := struct {
		Scheme string
		Name   string
		Port   string
		Path   string
		Params map[string]string
	}{
		Scheme: scheme,
		Name:   name,
		Port:   port,
		Path:   path,
		Params: params,
	}
	mock.lockProxyGet.Lock()
	mock.calls.ProxyGet = append(mock.calls.ProxyGet, callInfo)
	mock.lockProxyGet.Unlock()
	return mock.ProxyGetFunc(scheme, name, port, path, params)
}

// ProxyGetCalls gets all the calls that were made to ProxyGet.
// Check the length with:
//     len(mockedapiv1PodInterface.ProxyGetCalls())
func (mock *apiv1PodInterfaceMock) ProxyGetCalls() []struct {
	Scheme string
	Name   string
	Port   string
	Path   string
	Params map[string]string
} {
	var calls []struct {
		Scheme string
		Name   string
		Port   string
		Path   string
		Params map[string]string
	}
	mock.lockProxyGet.RLock()
	calls = mock.calls.ProxyGet
	mock.lockProxyGet.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *apiv1PodInterfaceMock) Update(ctx context.Context, pod *apiv1.Pod, opts metav1.UpdateOptions) (*apiv1.Pod, error) {
	if mock.UpdateFunc == nil {
		panic("apiv1PodInterfaceMock.UpdateFunc: method is nil but apiv1PodInterface.Update was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Pod  *apiv1.Pod
		Opts metav1.UpdateOptions
	}{
		Ctx:  ctx,
		Pod:  pod,
		Opts: opts,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, pod, opts)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedapiv1PodInterface.UpdateCalls())
func (mock *apiv1PodInterfaceMock) UpdateCalls() []struct {
	Ctx  context.Context
	Pod  *apiv1.Pod
	Opts metav1.UpdateOptions
} {
	var calls []struct {
		Ctx  context.Context
		Pod  *apiv1.Pod
		Opts metav1.UpdateOptions
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}

// UpdateEphemeralContainers calls UpdateEphemeralContainersFunc.
func (mock *apiv1PodInterfaceMock) UpdateEphemeralContainers(ctx context.Context, podName string, pod *apiv1.Pod, opts metav1.UpdateOptions) (*apiv1.Pod, error) {
	if mock.UpdateEphemeralContainersFunc == nil {
		panic("apiv1PodInterfaceMock.UpdateEphemeralContainersFunc: method is nil but apiv1PodInterface.UpdateEphemeralContainers was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		PodName string
		Pod     *apiv1.Pod
		Opts    metav1.UpdateOptions
	}{
		Ctx:     ctx,
		PodName: podName,
		Pod:     pod,
		Opts:    opts,
	}
	mock.lockUpdateEphemeralContainers.Lock()
	mock.calls.UpdateEphemeralContainers = append(mock.calls.UpdateEphemeralContainers, callInfo)
	mock.lockUpdateEphemeralContainers.Unlock()
	return mock.UpdateEphemeralContainersFunc(ctx, podName, pod, opts)
}

// UpdateEphemeralContainersCalls gets all the calls that were made to UpdateEphemeralContainers.
// Check the length with:
//     len(mockedapiv1PodInterface.UpdateEphemeralContainersCalls())
func (mock *apiv1PodInterfaceMock) UpdateEphemeralContainersCalls() []struct {
	Ctx     context.Context
	PodName string
	Pod     *apiv1.Pod
	Opts    metav1.UpdateOptions
} {
	var calls []struct {
		Ctx     context.Context
		PodName string
		Pod     *apiv1.Pod
		Opts    metav1.UpdateOptions
	}
	mock.lockUpdateEphemeralContainers.RLock()
	calls = mock.calls.UpdateEphemeralContainers
	mock.lockUpdateEphemeralContainers.RUnlock()
	return calls
}

// UpdateStatus calls UpdateStatusFunc.
func (mock *apiv1PodInterfaceMock) UpdateStatus(ctx context.Context, pod *apiv1.Pod, opts metav1.UpdateOptions) (*apiv1.Pod, error) {
	if mock.UpdateStatusFunc == nil {
		panic("apiv1PodInterfaceMock.UpdateStatusFunc: method is nil but apiv1PodInterface.UpdateStatus was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Pod  *apiv1.Pod
		Opts metav1.UpdateOptions
	}{
		Ctx:  ctx,
		Pod:  pod,
		Opts: opts,
	}
	mock.lockUpdateStatus.Lock()
	mock.calls.UpdateStatus = append(mock.calls.UpdateStatus, callInfo)
	mock.lockUpdateStatus.Unlock()
	return mock.UpdateStatusFunc(ctx, pod, opts)
}

// UpdateStatusCalls gets all the calls that were made to UpdateStatus.
// Check the length with:
//     len(mockedapiv1PodInterface.UpdateStatusCalls())
func (mock *apiv1PodInterfaceMock) UpdateStatusCalls() []struct {
	Ctx  context.Context
	Pod  *apiv1.Pod
	Opts metav1.UpdateOptions
} {
	var calls []struct {
		Ctx  context.Context
		Pod  *apiv1.Pod
		Opts metav1.UpdateOptions
	}
	mock.lockUpdateStatus.RLock()
	calls = mock.calls.UpdateStatus
	mock.lockUpdateStatus.RUnlock()
	return calls
}

// Watch calls WatchFunc.
func (mock *apiv1PodInterfaceMock) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	if mock.WatchFunc == nil {
		panic("apiv1PodInterfaceMock.WatchFunc: method is nil but apiv1PodInterface.Watch was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Opts metav1.ListOptions
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockWatch.Lock()
	mock.calls.Watch = append(mock.calls.Watch, callInfo)
	mock.lockWatch.Unlock()
	return mock.WatchFunc(ctx, opts)
}

// WatchCalls gets all the calls that were made to Watch.
// Check the length with:
//     len(mockedapiv1PodInterface.WatchCalls())
func (mock *apiv1PodInterfaceMock) WatchCalls() []struct {
	Ctx  context.Context
	Opts metav1.ListOptions
} {
	var calls []struct {
		Ctx  context.Context
		Opts metav1.ListOptions
	}
	mock.lockWatch.RLock()
	calls = mock.calls.Watch
	mock.lockWatch.RUnlock()
	return calls
}