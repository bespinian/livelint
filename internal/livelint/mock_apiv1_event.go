// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package livelint

import (
	"context"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/applyconfigurations/core/v1"
	"sync"
)

// Ensure, that apiv1EventInterfaceMock does implement apiv1EventInterface.
// If this is not the case, regenerate this file with moq.
var _ apiv1EventInterface = &apiv1EventInterfaceMock{}

// apiv1EventInterfaceMock is a mock implementation of apiv1EventInterface.
//
// 	func TestSomethingThatUsesapiv1EventInterface(t *testing.T) {
//
// 		// make and configure a mocked apiv1EventInterface
// 		mockedapiv1EventInterface := &apiv1EventInterfaceMock{
// 			ApplyFunc: func(ctx context.Context, event *v1.EventApplyConfiguration, opts metav1.ApplyOptions) (*apiv1.Event, error) {
// 				panic("mock out the Apply method")
// 			},
// 			CreateFunc: func(ctx context.Context, event *apiv1.Event, opts metav1.CreateOptions) (*apiv1.Event, error) {
// 				panic("mock out the Create method")
// 			},
// 			CreateWithEventNamespaceFunc: func(event *apiv1.Event) (*apiv1.Event, error) {
// 				panic("mock out the CreateWithEventNamespace method")
// 			},
// 			DeleteFunc: func(ctx context.Context, name string, opts metav1.DeleteOptions) error {
// 				panic("mock out the Delete method")
// 			},
// 			DeleteCollectionFunc: func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
// 				panic("mock out the DeleteCollection method")
// 			},
// 			GetFunc: func(ctx context.Context, name string, opts metav1.GetOptions) (*apiv1.Event, error) {
// 				panic("mock out the Get method")
// 			},
// 			GetFieldSelectorFunc: func(involvedObjectName *string, involvedObjectNamespace *string, involvedObjectKind *string, involvedObjectUID *string) fields.Selector {
// 				panic("mock out the GetFieldSelector method")
// 			},
// 			ListFunc: func(ctx context.Context, opts metav1.ListOptions) (*apiv1.EventList, error) {
// 				panic("mock out the List method")
// 			},
// 			PatchFunc: func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*apiv1.Event, error) {
// 				panic("mock out the Patch method")
// 			},
// 			PatchWithEventNamespaceFunc: func(event *apiv1.Event, data []byte) (*apiv1.Event, error) {
// 				panic("mock out the PatchWithEventNamespace method")
// 			},
// 			SearchFunc: func(scheme *runtime.Scheme, objOrRef runtime.Object) (*apiv1.EventList, error) {
// 				panic("mock out the Search method")
// 			},
// 			UpdateFunc: func(ctx context.Context, event *apiv1.Event, opts metav1.UpdateOptions) (*apiv1.Event, error) {
// 				panic("mock out the Update method")
// 			},
// 			UpdateWithEventNamespaceFunc: func(event *apiv1.Event) (*apiv1.Event, error) {
// 				panic("mock out the UpdateWithEventNamespace method")
// 			},
// 			WatchFunc: func(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
// 				panic("mock out the Watch method")
// 			},
// 		}
//
// 		// use mockedapiv1EventInterface in code that requires apiv1EventInterface
// 		// and then make assertions.
//
// 	}
type apiv1EventInterfaceMock struct {
	// ApplyFunc mocks the Apply method.
	ApplyFunc func(ctx context.Context, event *v1.EventApplyConfiguration, opts metav1.ApplyOptions) (*apiv1.Event, error)

	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, event *apiv1.Event, opts metav1.CreateOptions) (*apiv1.Event, error)

	// CreateWithEventNamespaceFunc mocks the CreateWithEventNamespace method.
	CreateWithEventNamespaceFunc func(event *apiv1.Event) (*apiv1.Event, error)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, name string, opts metav1.DeleteOptions) error

	// DeleteCollectionFunc mocks the DeleteCollection method.
	DeleteCollectionFunc func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error

	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, name string, opts metav1.GetOptions) (*apiv1.Event, error)

	// GetFieldSelectorFunc mocks the GetFieldSelector method.
	GetFieldSelectorFunc func(involvedObjectName *string, involvedObjectNamespace *string, involvedObjectKind *string, involvedObjectUID *string) fields.Selector

	// ListFunc mocks the List method.
	ListFunc func(ctx context.Context, opts metav1.ListOptions) (*apiv1.EventList, error)

	// PatchFunc mocks the Patch method.
	PatchFunc func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*apiv1.Event, error)

	// PatchWithEventNamespaceFunc mocks the PatchWithEventNamespace method.
	PatchWithEventNamespaceFunc func(event *apiv1.Event, data []byte) (*apiv1.Event, error)

	// SearchFunc mocks the Search method.
	SearchFunc func(scheme *runtime.Scheme, objOrRef runtime.Object) (*apiv1.EventList, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, event *apiv1.Event, opts metav1.UpdateOptions) (*apiv1.Event, error)

	// UpdateWithEventNamespaceFunc mocks the UpdateWithEventNamespace method.
	UpdateWithEventNamespaceFunc func(event *apiv1.Event) (*apiv1.Event, error)

	// WatchFunc mocks the Watch method.
	WatchFunc func(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)

	// calls tracks calls to the methods.
	calls struct {
		// Apply holds details about calls to the Apply method.
		Apply []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Event is the event argument value.
			Event *v1.EventApplyConfiguration
			// Opts is the opts argument value.
			Opts metav1.ApplyOptions
		}
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Event is the event argument value.
			Event *apiv1.Event
			// Opts is the opts argument value.
			Opts metav1.CreateOptions
		}
		// CreateWithEventNamespace holds details about calls to the CreateWithEventNamespace method.
		CreateWithEventNamespace []struct {
			// Event is the event argument value.
			Event *apiv1.Event
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
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// Opts is the opts argument value.
			Opts metav1.GetOptions
		}
		// GetFieldSelector holds details about calls to the GetFieldSelector method.
		GetFieldSelector []struct {
			// InvolvedObjectName is the involvedObjectName argument value.
			InvolvedObjectName *string
			// InvolvedObjectNamespace is the involvedObjectNamespace argument value.
			InvolvedObjectNamespace *string
			// InvolvedObjectKind is the involvedObjectKind argument value.
			InvolvedObjectKind *string
			// InvolvedObjectUID is the involvedObjectUID argument value.
			InvolvedObjectUID *string
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
		// PatchWithEventNamespace holds details about calls to the PatchWithEventNamespace method.
		PatchWithEventNamespace []struct {
			// Event is the event argument value.
			Event *apiv1.Event
			// Data is the data argument value.
			Data []byte
		}
		// Search holds details about calls to the Search method.
		Search []struct {
			// Scheme is the scheme argument value.
			Scheme *runtime.Scheme
			// ObjOrRef is the objOrRef argument value.
			ObjOrRef runtime.Object
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Event is the event argument value.
			Event *apiv1.Event
			// Opts is the opts argument value.
			Opts metav1.UpdateOptions
		}
		// UpdateWithEventNamespace holds details about calls to the UpdateWithEventNamespace method.
		UpdateWithEventNamespace []struct {
			// Event is the event argument value.
			Event *apiv1.Event
		}
		// Watch holds details about calls to the Watch method.
		Watch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts metav1.ListOptions
		}
	}
	lockApply                    sync.RWMutex
	lockCreate                   sync.RWMutex
	lockCreateWithEventNamespace sync.RWMutex
	lockDelete                   sync.RWMutex
	lockDeleteCollection         sync.RWMutex
	lockGet                      sync.RWMutex
	lockGetFieldSelector         sync.RWMutex
	lockList                     sync.RWMutex
	lockPatch                    sync.RWMutex
	lockPatchWithEventNamespace  sync.RWMutex
	lockSearch                   sync.RWMutex
	lockUpdate                   sync.RWMutex
	lockUpdateWithEventNamespace sync.RWMutex
	lockWatch                    sync.RWMutex
}

// Apply calls ApplyFunc.
func (mock *apiv1EventInterfaceMock) Apply(ctx context.Context, event *v1.EventApplyConfiguration, opts metav1.ApplyOptions) (*apiv1.Event, error) {
	if mock.ApplyFunc == nil {
		panic("apiv1EventInterfaceMock.ApplyFunc: method is nil but apiv1EventInterface.Apply was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Event *v1.EventApplyConfiguration
		Opts  metav1.ApplyOptions
	}{
		Ctx:   ctx,
		Event: event,
		Opts:  opts,
	}
	mock.lockApply.Lock()
	mock.calls.Apply = append(mock.calls.Apply, callInfo)
	mock.lockApply.Unlock()
	return mock.ApplyFunc(ctx, event, opts)
}

// ApplyCalls gets all the calls that were made to Apply.
// Check the length with:
//     len(mockedapiv1EventInterface.ApplyCalls())
func (mock *apiv1EventInterfaceMock) ApplyCalls() []struct {
	Ctx   context.Context
	Event *v1.EventApplyConfiguration
	Opts  metav1.ApplyOptions
} {
	var calls []struct {
		Ctx   context.Context
		Event *v1.EventApplyConfiguration
		Opts  metav1.ApplyOptions
	}
	mock.lockApply.RLock()
	calls = mock.calls.Apply
	mock.lockApply.RUnlock()
	return calls
}

// Create calls CreateFunc.
func (mock *apiv1EventInterfaceMock) Create(ctx context.Context, event *apiv1.Event, opts metav1.CreateOptions) (*apiv1.Event, error) {
	if mock.CreateFunc == nil {
		panic("apiv1EventInterfaceMock.CreateFunc: method is nil but apiv1EventInterface.Create was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Event *apiv1.Event
		Opts  metav1.CreateOptions
	}{
		Ctx:   ctx,
		Event: event,
		Opts:  opts,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, event, opts)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedapiv1EventInterface.CreateCalls())
func (mock *apiv1EventInterfaceMock) CreateCalls() []struct {
	Ctx   context.Context
	Event *apiv1.Event
	Opts  metav1.CreateOptions
} {
	var calls []struct {
		Ctx   context.Context
		Event *apiv1.Event
		Opts  metav1.CreateOptions
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// CreateWithEventNamespace calls CreateWithEventNamespaceFunc.
func (mock *apiv1EventInterfaceMock) CreateWithEventNamespace(event *apiv1.Event) (*apiv1.Event, error) {
	if mock.CreateWithEventNamespaceFunc == nil {
		panic("apiv1EventInterfaceMock.CreateWithEventNamespaceFunc: method is nil but apiv1EventInterface.CreateWithEventNamespace was just called")
	}
	callInfo := struct {
		Event *apiv1.Event
	}{
		Event: event,
	}
	mock.lockCreateWithEventNamespace.Lock()
	mock.calls.CreateWithEventNamespace = append(mock.calls.CreateWithEventNamespace, callInfo)
	mock.lockCreateWithEventNamespace.Unlock()
	return mock.CreateWithEventNamespaceFunc(event)
}

// CreateWithEventNamespaceCalls gets all the calls that were made to CreateWithEventNamespace.
// Check the length with:
//     len(mockedapiv1EventInterface.CreateWithEventNamespaceCalls())
func (mock *apiv1EventInterfaceMock) CreateWithEventNamespaceCalls() []struct {
	Event *apiv1.Event
} {
	var calls []struct {
		Event *apiv1.Event
	}
	mock.lockCreateWithEventNamespace.RLock()
	calls = mock.calls.CreateWithEventNamespace
	mock.lockCreateWithEventNamespace.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *apiv1EventInterfaceMock) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	if mock.DeleteFunc == nil {
		panic("apiv1EventInterfaceMock.DeleteFunc: method is nil but apiv1EventInterface.Delete was just called")
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
//     len(mockedapiv1EventInterface.DeleteCalls())
func (mock *apiv1EventInterfaceMock) DeleteCalls() []struct {
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
func (mock *apiv1EventInterfaceMock) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	if mock.DeleteCollectionFunc == nil {
		panic("apiv1EventInterfaceMock.DeleteCollectionFunc: method is nil but apiv1EventInterface.DeleteCollection was just called")
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
//     len(mockedapiv1EventInterface.DeleteCollectionCalls())
func (mock *apiv1EventInterfaceMock) DeleteCollectionCalls() []struct {
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

// Get calls GetFunc.
func (mock *apiv1EventInterfaceMock) Get(ctx context.Context, name string, opts metav1.GetOptions) (*apiv1.Event, error) {
	if mock.GetFunc == nil {
		panic("apiv1EventInterfaceMock.GetFunc: method is nil but apiv1EventInterface.Get was just called")
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
//     len(mockedapiv1EventInterface.GetCalls())
func (mock *apiv1EventInterfaceMock) GetCalls() []struct {
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

// GetFieldSelector calls GetFieldSelectorFunc.
func (mock *apiv1EventInterfaceMock) GetFieldSelector(involvedObjectName *string, involvedObjectNamespace *string, involvedObjectKind *string, involvedObjectUID *string) fields.Selector {
	if mock.GetFieldSelectorFunc == nil {
		panic("apiv1EventInterfaceMock.GetFieldSelectorFunc: method is nil but apiv1EventInterface.GetFieldSelector was just called")
	}
	callInfo := struct {
		InvolvedObjectName      *string
		InvolvedObjectNamespace *string
		InvolvedObjectKind      *string
		InvolvedObjectUID       *string
	}{
		InvolvedObjectName:      involvedObjectName,
		InvolvedObjectNamespace: involvedObjectNamespace,
		InvolvedObjectKind:      involvedObjectKind,
		InvolvedObjectUID:       involvedObjectUID,
	}
	mock.lockGetFieldSelector.Lock()
	mock.calls.GetFieldSelector = append(mock.calls.GetFieldSelector, callInfo)
	mock.lockGetFieldSelector.Unlock()
	return mock.GetFieldSelectorFunc(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID)
}

// GetFieldSelectorCalls gets all the calls that were made to GetFieldSelector.
// Check the length with:
//     len(mockedapiv1EventInterface.GetFieldSelectorCalls())
func (mock *apiv1EventInterfaceMock) GetFieldSelectorCalls() []struct {
	InvolvedObjectName      *string
	InvolvedObjectNamespace *string
	InvolvedObjectKind      *string
	InvolvedObjectUID       *string
} {
	var calls []struct {
		InvolvedObjectName      *string
		InvolvedObjectNamespace *string
		InvolvedObjectKind      *string
		InvolvedObjectUID       *string
	}
	mock.lockGetFieldSelector.RLock()
	calls = mock.calls.GetFieldSelector
	mock.lockGetFieldSelector.RUnlock()
	return calls
}

// List calls ListFunc.
func (mock *apiv1EventInterfaceMock) List(ctx context.Context, opts metav1.ListOptions) (*apiv1.EventList, error) {
	if mock.ListFunc == nil {
		panic("apiv1EventInterfaceMock.ListFunc: method is nil but apiv1EventInterface.List was just called")
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
//     len(mockedapiv1EventInterface.ListCalls())
func (mock *apiv1EventInterfaceMock) ListCalls() []struct {
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
func (mock *apiv1EventInterfaceMock) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*apiv1.Event, error) {
	if mock.PatchFunc == nil {
		panic("apiv1EventInterfaceMock.PatchFunc: method is nil but apiv1EventInterface.Patch was just called")
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
//     len(mockedapiv1EventInterface.PatchCalls())
func (mock *apiv1EventInterfaceMock) PatchCalls() []struct {
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

// PatchWithEventNamespace calls PatchWithEventNamespaceFunc.
func (mock *apiv1EventInterfaceMock) PatchWithEventNamespace(event *apiv1.Event, data []byte) (*apiv1.Event, error) {
	if mock.PatchWithEventNamespaceFunc == nil {
		panic("apiv1EventInterfaceMock.PatchWithEventNamespaceFunc: method is nil but apiv1EventInterface.PatchWithEventNamespace was just called")
	}
	callInfo := struct {
		Event *apiv1.Event
		Data  []byte
	}{
		Event: event,
		Data:  data,
	}
	mock.lockPatchWithEventNamespace.Lock()
	mock.calls.PatchWithEventNamespace = append(mock.calls.PatchWithEventNamespace, callInfo)
	mock.lockPatchWithEventNamespace.Unlock()
	return mock.PatchWithEventNamespaceFunc(event, data)
}

// PatchWithEventNamespaceCalls gets all the calls that were made to PatchWithEventNamespace.
// Check the length with:
//     len(mockedapiv1EventInterface.PatchWithEventNamespaceCalls())
func (mock *apiv1EventInterfaceMock) PatchWithEventNamespaceCalls() []struct {
	Event *apiv1.Event
	Data  []byte
} {
	var calls []struct {
		Event *apiv1.Event
		Data  []byte
	}
	mock.lockPatchWithEventNamespace.RLock()
	calls = mock.calls.PatchWithEventNamespace
	mock.lockPatchWithEventNamespace.RUnlock()
	return calls
}

// Search calls SearchFunc.
func (mock *apiv1EventInterfaceMock) Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*apiv1.EventList, error) {
	if mock.SearchFunc == nil {
		panic("apiv1EventInterfaceMock.SearchFunc: method is nil but apiv1EventInterface.Search was just called")
	}
	callInfo := struct {
		Scheme   *runtime.Scheme
		ObjOrRef runtime.Object
	}{
		Scheme:   scheme,
		ObjOrRef: objOrRef,
	}
	mock.lockSearch.Lock()
	mock.calls.Search = append(mock.calls.Search, callInfo)
	mock.lockSearch.Unlock()
	return mock.SearchFunc(scheme, objOrRef)
}

// SearchCalls gets all the calls that were made to Search.
// Check the length with:
//     len(mockedapiv1EventInterface.SearchCalls())
func (mock *apiv1EventInterfaceMock) SearchCalls() []struct {
	Scheme   *runtime.Scheme
	ObjOrRef runtime.Object
} {
	var calls []struct {
		Scheme   *runtime.Scheme
		ObjOrRef runtime.Object
	}
	mock.lockSearch.RLock()
	calls = mock.calls.Search
	mock.lockSearch.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *apiv1EventInterfaceMock) Update(ctx context.Context, event *apiv1.Event, opts metav1.UpdateOptions) (*apiv1.Event, error) {
	if mock.UpdateFunc == nil {
		panic("apiv1EventInterfaceMock.UpdateFunc: method is nil but apiv1EventInterface.Update was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Event *apiv1.Event
		Opts  metav1.UpdateOptions
	}{
		Ctx:   ctx,
		Event: event,
		Opts:  opts,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, event, opts)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedapiv1EventInterface.UpdateCalls())
func (mock *apiv1EventInterfaceMock) UpdateCalls() []struct {
	Ctx   context.Context
	Event *apiv1.Event
	Opts  metav1.UpdateOptions
} {
	var calls []struct {
		Ctx   context.Context
		Event *apiv1.Event
		Opts  metav1.UpdateOptions
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}

// UpdateWithEventNamespace calls UpdateWithEventNamespaceFunc.
func (mock *apiv1EventInterfaceMock) UpdateWithEventNamespace(event *apiv1.Event) (*apiv1.Event, error) {
	if mock.UpdateWithEventNamespaceFunc == nil {
		panic("apiv1EventInterfaceMock.UpdateWithEventNamespaceFunc: method is nil but apiv1EventInterface.UpdateWithEventNamespace was just called")
	}
	callInfo := struct {
		Event *apiv1.Event
	}{
		Event: event,
	}
	mock.lockUpdateWithEventNamespace.Lock()
	mock.calls.UpdateWithEventNamespace = append(mock.calls.UpdateWithEventNamespace, callInfo)
	mock.lockUpdateWithEventNamespace.Unlock()
	return mock.UpdateWithEventNamespaceFunc(event)
}

// UpdateWithEventNamespaceCalls gets all the calls that were made to UpdateWithEventNamespace.
// Check the length with:
//     len(mockedapiv1EventInterface.UpdateWithEventNamespaceCalls())
func (mock *apiv1EventInterfaceMock) UpdateWithEventNamespaceCalls() []struct {
	Event *apiv1.Event
} {
	var calls []struct {
		Event *apiv1.Event
	}
	mock.lockUpdateWithEventNamespace.RLock()
	calls = mock.calls.UpdateWithEventNamespace
	mock.lockUpdateWithEventNamespace.RUnlock()
	return calls
}

// Watch calls WatchFunc.
func (mock *apiv1EventInterfaceMock) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	if mock.WatchFunc == nil {
		panic("apiv1EventInterfaceMock.WatchFunc: method is nil but apiv1EventInterface.Watch was just called")
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
//     len(mockedapiv1EventInterface.WatchCalls())
func (mock *apiv1EventInterfaceMock) WatchCalls() []struct {
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
