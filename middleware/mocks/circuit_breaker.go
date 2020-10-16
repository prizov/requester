// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"sync"
)

// CircuitBreakerSvcMock is a mock implementation of middleware.CircuitBreakerSvc.
//
//     func TestSomethingThatUsesCircuitBreakerSvc(t *testing.T) {
//
//         // make and configure a mocked middleware.CircuitBreakerSvc
//         mockedCircuitBreakerSvc := &CircuitBreakerSvcMock{
//             ExecuteFunc: func(req func() (interface{}, error)) (interface{}, error) {
// 	               panic("mock out the Execute method")
//             },
//         }
//
//         // use mockedCircuitBreakerSvc in code that requires middleware.CircuitBreakerSvc
//         // and then make assertions.
//
//     }
type CircuitBreakerSvcMock struct {
	// ExecuteFunc mocks the Execute method.
	ExecuteFunc func(req func() (interface{}, error)) (interface{}, error)

	// calls tracks calls to the methods.
	calls struct {
		// Execute holds details about calls to the Execute method.
		Execute []struct {
			// Req is the req argument value.
			Req func() (interface{}, error)
		}
	}
	lockExecute sync.RWMutex
}

// Execute calls ExecuteFunc.
func (mock *CircuitBreakerSvcMock) Execute(req func() (interface{}, error)) (interface{}, error) {
	if mock.ExecuteFunc == nil {
		panic("CircuitBreakerSvcMock.ExecuteFunc: method is nil but CircuitBreakerSvc.Execute was just called")
	}
	callInfo := struct {
		Req func() (interface{}, error)
	}{
		Req: req,
	}
	mock.lockExecute.Lock()
	mock.calls.Execute = append(mock.calls.Execute, callInfo)
	mock.lockExecute.Unlock()
	return mock.ExecuteFunc(req)
}

// ExecuteCalls gets all the calls that were made to Execute.
// Check the length with:
//     len(mockedCircuitBreakerSvc.ExecuteCalls())
func (mock *CircuitBreakerSvcMock) ExecuteCalls() []struct {
	Req func() (interface{}, error)
} {
	var calls []struct {
		Req func() (interface{}, error)
	}
	mock.lockExecute.RLock()
	calls = mock.calls.Execute
	mock.lockExecute.RUnlock()
	return calls
}
