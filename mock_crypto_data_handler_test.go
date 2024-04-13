// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/daeuniverse/quic-go (interfaces: CryptoDataHandler)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/daeuniverse/quic-go -destination mock_crypto_data_handler_test.go github.com/daeuniverse/quic-go CryptoDataHandler
//

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	handshake "github.com/daeuniverse/quic-go/internal/handshake"
	protocol "github.com/daeuniverse/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockCryptoDataHandler is a mock of CryptoDataHandler interface.
type MockCryptoDataHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoDataHandlerMockRecorder
}

// MockCryptoDataHandlerMockRecorder is the mock recorder for MockCryptoDataHandler.
type MockCryptoDataHandlerMockRecorder struct {
	mock *MockCryptoDataHandler
}

// NewMockCryptoDataHandler creates a new mock instance.
func NewMockCryptoDataHandler(ctrl *gomock.Controller) *MockCryptoDataHandler {
	mock := &MockCryptoDataHandler{ctrl: ctrl}
	mock.recorder = &MockCryptoDataHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoDataHandler) EXPECT() *MockCryptoDataHandlerMockRecorder {
	return m.recorder
}

// HandleMessage mocks base method.
func (m *MockCryptoDataHandler) HandleMessage(arg0 []byte, arg1 protocol.EncryptionLevel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleMessage indicates an expected call of HandleMessage.
func (mr *MockCryptoDataHandlerMockRecorder) HandleMessage(arg0, arg1 any) *MockCryptoDataHandlerHandleMessageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockCryptoDataHandler)(nil).HandleMessage), arg0, arg1)
	return &MockCryptoDataHandlerHandleMessageCall{Call: call}
}

// MockCryptoDataHandlerHandleMessageCall wrap *gomock.Call
type MockCryptoDataHandlerHandleMessageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoDataHandlerHandleMessageCall) Return(arg0 error) *MockCryptoDataHandlerHandleMessageCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoDataHandlerHandleMessageCall) Do(f func([]byte, protocol.EncryptionLevel) error) *MockCryptoDataHandlerHandleMessageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoDataHandlerHandleMessageCall) DoAndReturn(f func([]byte, protocol.EncryptionLevel) error) *MockCryptoDataHandlerHandleMessageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NextEvent mocks base method.
func (m *MockCryptoDataHandler) NextEvent() handshake.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextEvent")
	ret0, _ := ret[0].(handshake.Event)
	return ret0
}

// NextEvent indicates an expected call of NextEvent.
func (mr *MockCryptoDataHandlerMockRecorder) NextEvent() *MockCryptoDataHandlerNextEventCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextEvent", reflect.TypeOf((*MockCryptoDataHandler)(nil).NextEvent))
	return &MockCryptoDataHandlerNextEventCall{Call: call}
}

// MockCryptoDataHandlerNextEventCall wrap *gomock.Call
type MockCryptoDataHandlerNextEventCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoDataHandlerNextEventCall) Return(arg0 handshake.Event) *MockCryptoDataHandlerNextEventCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoDataHandlerNextEventCall) Do(f func() handshake.Event) *MockCryptoDataHandlerNextEventCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoDataHandlerNextEventCall) DoAndReturn(f func() handshake.Event) *MockCryptoDataHandlerNextEventCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
