// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"sync"
)

// RepeaterSvcMock is a mock implementation of middleware.RepeaterSvc.
//
//     func TestSomethingThatUsesRepeaterSvc(t *testing.T) {
//
//         // make and configure a mocked middleware.RepeaterSvc
//         mockedRepeaterSvc := &RepeaterSvcMock{
//             DoFunc: func(fun func() error, errs ...error) error {
// 	               panic("mock out the Do method")
//             },
//         }
//
//         // use mockedRepeaterSvc in code that requires middleware.RepeaterSvc
//         // and then make assertions.
//
//     }
type RepeaterSvcMock struct {
	// DoFunc mocks the Do method.
	DoFunc func(fun func() error, errs ...error) error

	// calls tracks calls to the methods.
	calls struct {
		// Do holds details about calls to the Do method.
		Do []struct {
			// Fun is the fun argument value.
			Fun func() error
			// Errs is the errs argument value.
			Errs []error
		}
	}
	lockDo sync.RWMutex
}

// Do calls DoFunc.
func (mock *RepeaterSvcMock) Do(fun func() error, errs ...error) error {
	if mock.DoFunc == nil {
		panic("RepeaterSvcMock.DoFunc: method is nil but RepeaterSvc.Do was just called")
	}
	callInfo := struct {
		Fun  func() error
		Errs []error
	}{
		Fun:  fun,
		Errs: errs,
	}
	mock.lockDo.Lock()
	mock.calls.Do = append(mock.calls.Do, callInfo)
	mock.lockDo.Unlock()
	return mock.DoFunc(fun, errs...)
}

// DoCalls gets all the calls that were made to Do.
// Check the length with:
//     len(mockedRepeaterSvc.DoCalls())
func (mock *RepeaterSvcMock) DoCalls() []struct {
	Fun  func() error
	Errs []error
} {
	var calls []struct {
		Fun  func() error
		Errs []error
	}
	mock.lockDo.RLock()
	calls = mock.calls.Do
	mock.lockDo.RUnlock()
	return calls
}
