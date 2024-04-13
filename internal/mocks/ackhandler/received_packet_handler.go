// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/daeuniverse/quic-go/internal/ackhandler (interfaces: ReceivedPacketHandler)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package mockackhandler -destination ackhandler/received_packet_handler.go github.com/daeuniverse/quic-go/internal/ackhandler ReceivedPacketHandler
//

// Package mockackhandler is a generated GoMock package.
package mockackhandler

import (
	reflect "reflect"
	time "time"

	protocol "github.com/daeuniverse/quic-go/internal/protocol"
	wire "github.com/daeuniverse/quic-go/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockReceivedPacketHandler is a mock of ReceivedPacketHandler interface.
type MockReceivedPacketHandler struct {
	ctrl     *gomock.Controller
	recorder *MockReceivedPacketHandlerMockRecorder
}

// MockReceivedPacketHandlerMockRecorder is the mock recorder for MockReceivedPacketHandler.
type MockReceivedPacketHandlerMockRecorder struct {
	mock *MockReceivedPacketHandler
}

// NewMockReceivedPacketHandler creates a new mock instance.
func NewMockReceivedPacketHandler(ctrl *gomock.Controller) *MockReceivedPacketHandler {
	mock := &MockReceivedPacketHandler{ctrl: ctrl}
	mock.recorder = &MockReceivedPacketHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceivedPacketHandler) EXPECT() *MockReceivedPacketHandlerMockRecorder {
	return m.recorder
}

// DropPackets mocks base method.
func (m *MockReceivedPacketHandler) DropPackets(arg0 protocol.EncryptionLevel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DropPackets", arg0)
}

// DropPackets indicates an expected call of DropPackets.
func (mr *MockReceivedPacketHandlerMockRecorder) DropPackets(arg0 any) *MockReceivedPacketHandlerDropPacketsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropPackets", reflect.TypeOf((*MockReceivedPacketHandler)(nil).DropPackets), arg0)
	return &MockReceivedPacketHandlerDropPacketsCall{Call: call}
}

// MockReceivedPacketHandlerDropPacketsCall wrap *gomock.Call
type MockReceivedPacketHandlerDropPacketsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockReceivedPacketHandlerDropPacketsCall) Return() *MockReceivedPacketHandlerDropPacketsCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockReceivedPacketHandlerDropPacketsCall) Do(f func(protocol.EncryptionLevel)) *MockReceivedPacketHandlerDropPacketsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockReceivedPacketHandlerDropPacketsCall) DoAndReturn(f func(protocol.EncryptionLevel)) *MockReceivedPacketHandlerDropPacketsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAckFrame mocks base method.
func (m *MockReceivedPacketHandler) GetAckFrame(arg0 protocol.EncryptionLevel, arg1 bool) *wire.AckFrame {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAckFrame", arg0, arg1)
	ret0, _ := ret[0].(*wire.AckFrame)
	return ret0
}

// GetAckFrame indicates an expected call of GetAckFrame.
func (mr *MockReceivedPacketHandlerMockRecorder) GetAckFrame(arg0, arg1 any) *MockReceivedPacketHandlerGetAckFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAckFrame", reflect.TypeOf((*MockReceivedPacketHandler)(nil).GetAckFrame), arg0, arg1)
	return &MockReceivedPacketHandlerGetAckFrameCall{Call: call}
}

// MockReceivedPacketHandlerGetAckFrameCall wrap *gomock.Call
type MockReceivedPacketHandlerGetAckFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockReceivedPacketHandlerGetAckFrameCall) Return(arg0 *wire.AckFrame) *MockReceivedPacketHandlerGetAckFrameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockReceivedPacketHandlerGetAckFrameCall) Do(f func(protocol.EncryptionLevel, bool) *wire.AckFrame) *MockReceivedPacketHandlerGetAckFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockReceivedPacketHandlerGetAckFrameCall) DoAndReturn(f func(protocol.EncryptionLevel, bool) *wire.AckFrame) *MockReceivedPacketHandlerGetAckFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAlarmTimeout mocks base method.
func (m *MockReceivedPacketHandler) GetAlarmTimeout() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlarmTimeout")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetAlarmTimeout indicates an expected call of GetAlarmTimeout.
func (mr *MockReceivedPacketHandlerMockRecorder) GetAlarmTimeout() *MockReceivedPacketHandlerGetAlarmTimeoutCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlarmTimeout", reflect.TypeOf((*MockReceivedPacketHandler)(nil).GetAlarmTimeout))
	return &MockReceivedPacketHandlerGetAlarmTimeoutCall{Call: call}
}

// MockReceivedPacketHandlerGetAlarmTimeoutCall wrap *gomock.Call
type MockReceivedPacketHandlerGetAlarmTimeoutCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockReceivedPacketHandlerGetAlarmTimeoutCall) Return(arg0 time.Time) *MockReceivedPacketHandlerGetAlarmTimeoutCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockReceivedPacketHandlerGetAlarmTimeoutCall) Do(f func() time.Time) *MockReceivedPacketHandlerGetAlarmTimeoutCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockReceivedPacketHandlerGetAlarmTimeoutCall) DoAndReturn(f func() time.Time) *MockReceivedPacketHandlerGetAlarmTimeoutCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsPotentiallyDuplicate mocks base method.
func (m *MockReceivedPacketHandler) IsPotentiallyDuplicate(arg0 protocol.PacketNumber, arg1 protocol.EncryptionLevel) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPotentiallyDuplicate", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPotentiallyDuplicate indicates an expected call of IsPotentiallyDuplicate.
func (mr *MockReceivedPacketHandlerMockRecorder) IsPotentiallyDuplicate(arg0, arg1 any) *MockReceivedPacketHandlerIsPotentiallyDuplicateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPotentiallyDuplicate", reflect.TypeOf((*MockReceivedPacketHandler)(nil).IsPotentiallyDuplicate), arg0, arg1)
	return &MockReceivedPacketHandlerIsPotentiallyDuplicateCall{Call: call}
}

// MockReceivedPacketHandlerIsPotentiallyDuplicateCall wrap *gomock.Call
type MockReceivedPacketHandlerIsPotentiallyDuplicateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockReceivedPacketHandlerIsPotentiallyDuplicateCall) Return(arg0 bool) *MockReceivedPacketHandlerIsPotentiallyDuplicateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockReceivedPacketHandlerIsPotentiallyDuplicateCall) Do(f func(protocol.PacketNumber, protocol.EncryptionLevel) bool) *MockReceivedPacketHandlerIsPotentiallyDuplicateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockReceivedPacketHandlerIsPotentiallyDuplicateCall) DoAndReturn(f func(protocol.PacketNumber, protocol.EncryptionLevel) bool) *MockReceivedPacketHandlerIsPotentiallyDuplicateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReceivedPacket mocks base method.
func (m *MockReceivedPacketHandler) ReceivedPacket(arg0 protocol.PacketNumber, arg1 protocol.ECN, arg2 protocol.EncryptionLevel, arg3 time.Time, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceivedPacket", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReceivedPacket indicates an expected call of ReceivedPacket.
func (mr *MockReceivedPacketHandlerMockRecorder) ReceivedPacket(arg0, arg1, arg2, arg3, arg4 any) *MockReceivedPacketHandlerReceivedPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedPacket", reflect.TypeOf((*MockReceivedPacketHandler)(nil).ReceivedPacket), arg0, arg1, arg2, arg3, arg4)
	return &MockReceivedPacketHandlerReceivedPacketCall{Call: call}
}

// MockReceivedPacketHandlerReceivedPacketCall wrap *gomock.Call
type MockReceivedPacketHandlerReceivedPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockReceivedPacketHandlerReceivedPacketCall) Return(arg0 error) *MockReceivedPacketHandlerReceivedPacketCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockReceivedPacketHandlerReceivedPacketCall) Do(f func(protocol.PacketNumber, protocol.ECN, protocol.EncryptionLevel, time.Time, bool) error) *MockReceivedPacketHandlerReceivedPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockReceivedPacketHandlerReceivedPacketCall) DoAndReturn(f func(protocol.PacketNumber, protocol.ECN, protocol.EncryptionLevel, time.Time, bool) error) *MockReceivedPacketHandlerReceivedPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
