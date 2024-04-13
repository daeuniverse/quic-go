// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/daeuniverse/quic-go (interfaces: StreamI)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/daeuniverse/quic-go -destination mock_stream_internal_test.go github.com/daeuniverse/quic-go StreamI
//

// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	reflect "reflect"
	time "time"

	ackhandler "github.com/daeuniverse/quic-go/internal/ackhandler"
	protocol "github.com/daeuniverse/quic-go/internal/protocol"
	qerr "github.com/daeuniverse/quic-go/internal/qerr"
	wire "github.com/daeuniverse/quic-go/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockStreamI is a mock of StreamI interface.
type MockStreamI struct {
	ctrl     *gomock.Controller
	recorder *MockStreamIMockRecorder
}

// MockStreamIMockRecorder is the mock recorder for MockStreamI.
type MockStreamIMockRecorder struct {
	mock *MockStreamI
}

// NewMockStreamI creates a new mock instance.
func NewMockStreamI(ctrl *gomock.Controller) *MockStreamI {
	mock := &MockStreamI{ctrl: ctrl}
	mock.recorder = &MockStreamIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamI) EXPECT() *MockStreamIMockRecorder {
	return m.recorder
}

// CancelRead mocks base method.
func (m *MockStreamI) CancelRead(arg0 qerr.StreamErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelRead", arg0)
}

// CancelRead indicates an expected call of CancelRead.
func (mr *MockStreamIMockRecorder) CancelRead(arg0 any) *MockStreamICancelReadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRead", reflect.TypeOf((*MockStreamI)(nil).CancelRead), arg0)
	return &MockStreamICancelReadCall{Call: call}
}

// MockStreamICancelReadCall wrap *gomock.Call
type MockStreamICancelReadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamICancelReadCall) Return() *MockStreamICancelReadCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamICancelReadCall) Do(f func(qerr.StreamErrorCode)) *MockStreamICancelReadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamICancelReadCall) DoAndReturn(f func(qerr.StreamErrorCode)) *MockStreamICancelReadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CancelWrite mocks base method.
func (m *MockStreamI) CancelWrite(arg0 qerr.StreamErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelWrite", arg0)
}

// CancelWrite indicates an expected call of CancelWrite.
func (mr *MockStreamIMockRecorder) CancelWrite(arg0 any) *MockStreamICancelWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelWrite", reflect.TypeOf((*MockStreamI)(nil).CancelWrite), arg0)
	return &MockStreamICancelWriteCall{Call: call}
}

// MockStreamICancelWriteCall wrap *gomock.Call
type MockStreamICancelWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamICancelWriteCall) Return() *MockStreamICancelWriteCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamICancelWriteCall) Do(f func(qerr.StreamErrorCode)) *MockStreamICancelWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamICancelWriteCall) DoAndReturn(f func(qerr.StreamErrorCode)) *MockStreamICancelWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockStreamI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStreamIMockRecorder) Close() *MockStreamICloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStreamI)(nil).Close))
	return &MockStreamICloseCall{Call: call}
}

// MockStreamICloseCall wrap *gomock.Call
type MockStreamICloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamICloseCall) Return(arg0 error) *MockStreamICloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamICloseCall) Do(f func() error) *MockStreamICloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamICloseCall) DoAndReturn(f func() error) *MockStreamICloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Context mocks base method.
func (m *MockStreamI) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockStreamIMockRecorder) Context() *MockStreamIContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockStreamI)(nil).Context))
	return &MockStreamIContextCall{Call: call}
}

// MockStreamIContextCall wrap *gomock.Call
type MockStreamIContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamIContextCall) Return(arg0 context.Context) *MockStreamIContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamIContextCall) Do(f func() context.Context) *MockStreamIContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamIContextCall) DoAndReturn(f func() context.Context) *MockStreamIContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Read mocks base method.
func (m *MockStreamI) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockStreamIMockRecorder) Read(arg0 any) *MockStreamIReadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockStreamI)(nil).Read), arg0)
	return &MockStreamIReadCall{Call: call}
}

// MockStreamIReadCall wrap *gomock.Call
type MockStreamIReadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamIReadCall) Return(arg0 int, arg1 error) *MockStreamIReadCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamIReadCall) Do(f func([]byte) (int, error)) *MockStreamIReadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamIReadCall) DoAndReturn(f func([]byte) (int, error)) *MockStreamIReadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetDeadline mocks base method.
func (m *MockStreamI) SetDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline.
func (mr *MockStreamIMockRecorder) SetDeadline(arg0 any) *MockStreamISetDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockStreamI)(nil).SetDeadline), arg0)
	return &MockStreamISetDeadlineCall{Call: call}
}

// MockStreamISetDeadlineCall wrap *gomock.Call
type MockStreamISetDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamISetDeadlineCall) Return(arg0 error) *MockStreamISetDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamISetDeadlineCall) Do(f func(time.Time) error) *MockStreamISetDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamISetDeadlineCall) DoAndReturn(f func(time.Time) error) *MockStreamISetDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetReadDeadline mocks base method.
func (m *MockStreamI) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockStreamIMockRecorder) SetReadDeadline(arg0 any) *MockStreamISetReadDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockStreamI)(nil).SetReadDeadline), arg0)
	return &MockStreamISetReadDeadlineCall{Call: call}
}

// MockStreamISetReadDeadlineCall wrap *gomock.Call
type MockStreamISetReadDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamISetReadDeadlineCall) Return(arg0 error) *MockStreamISetReadDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamISetReadDeadlineCall) Do(f func(time.Time) error) *MockStreamISetReadDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamISetReadDeadlineCall) DoAndReturn(f func(time.Time) error) *MockStreamISetReadDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetWriteDeadline mocks base method.
func (m *MockStreamI) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockStreamIMockRecorder) SetWriteDeadline(arg0 any) *MockStreamISetWriteDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockStreamI)(nil).SetWriteDeadline), arg0)
	return &MockStreamISetWriteDeadlineCall{Call: call}
}

// MockStreamISetWriteDeadlineCall wrap *gomock.Call
type MockStreamISetWriteDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamISetWriteDeadlineCall) Return(arg0 error) *MockStreamISetWriteDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamISetWriteDeadlineCall) Do(f func(time.Time) error) *MockStreamISetWriteDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamISetWriteDeadlineCall) DoAndReturn(f func(time.Time) error) *MockStreamISetWriteDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StreamID mocks base method.
func (m *MockStreamI) StreamID() protocol.StreamID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamID")
	ret0, _ := ret[0].(protocol.StreamID)
	return ret0
}

// StreamID indicates an expected call of StreamID.
func (mr *MockStreamIMockRecorder) StreamID() *MockStreamIStreamIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamID", reflect.TypeOf((*MockStreamI)(nil).StreamID))
	return &MockStreamIStreamIDCall{Call: call}
}

// MockStreamIStreamIDCall wrap *gomock.Call
type MockStreamIStreamIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamIStreamIDCall) Return(arg0 protocol.StreamID) *MockStreamIStreamIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamIStreamIDCall) Do(f func() protocol.StreamID) *MockStreamIStreamIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamIStreamIDCall) DoAndReturn(f func() protocol.StreamID) *MockStreamIStreamIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Write mocks base method.
func (m *MockStreamI) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockStreamIMockRecorder) Write(arg0 any) *MockStreamIWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStreamI)(nil).Write), arg0)
	return &MockStreamIWriteCall{Call: call}
}

// MockStreamIWriteCall wrap *gomock.Call
type MockStreamIWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamIWriteCall) Return(arg0 int, arg1 error) *MockStreamIWriteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamIWriteCall) Do(f func([]byte) (int, error)) *MockStreamIWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamIWriteCall) DoAndReturn(f func([]byte) (int, error)) *MockStreamIWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// closeForShutdown mocks base method.
func (m *MockStreamI) closeForShutdown(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "closeForShutdown", arg0)
}

// closeForShutdown indicates an expected call of closeForShutdown.
func (mr *MockStreamIMockRecorder) closeForShutdown(arg0 any) *MockStreamIcloseForShutdownCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeForShutdown", reflect.TypeOf((*MockStreamI)(nil).closeForShutdown), arg0)
	return &MockStreamIcloseForShutdownCall{Call: call}
}

// MockStreamIcloseForShutdownCall wrap *gomock.Call
type MockStreamIcloseForShutdownCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamIcloseForShutdownCall) Return() *MockStreamIcloseForShutdownCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamIcloseForShutdownCall) Do(f func(error)) *MockStreamIcloseForShutdownCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamIcloseForShutdownCall) DoAndReturn(f func(error)) *MockStreamIcloseForShutdownCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// getWindowUpdate mocks base method.
func (m *MockStreamI) getWindowUpdate() protocol.ByteCount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getWindowUpdate")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// getWindowUpdate indicates an expected call of getWindowUpdate.
func (mr *MockStreamIMockRecorder) getWindowUpdate() *MockStreamIgetWindowUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getWindowUpdate", reflect.TypeOf((*MockStreamI)(nil).getWindowUpdate))
	return &MockStreamIgetWindowUpdateCall{Call: call}
}

// MockStreamIgetWindowUpdateCall wrap *gomock.Call
type MockStreamIgetWindowUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamIgetWindowUpdateCall) Return(arg0 protocol.ByteCount) *MockStreamIgetWindowUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamIgetWindowUpdateCall) Do(f func() protocol.ByteCount) *MockStreamIgetWindowUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamIgetWindowUpdateCall) DoAndReturn(f func() protocol.ByteCount) *MockStreamIgetWindowUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// handleResetStreamFrame mocks base method.
func (m *MockStreamI) handleResetStreamFrame(arg0 *wire.ResetStreamFrame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleResetStreamFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleResetStreamFrame indicates an expected call of handleResetStreamFrame.
func (mr *MockStreamIMockRecorder) handleResetStreamFrame(arg0 any) *MockStreamIhandleResetStreamFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleResetStreamFrame", reflect.TypeOf((*MockStreamI)(nil).handleResetStreamFrame), arg0)
	return &MockStreamIhandleResetStreamFrameCall{Call: call}
}

// MockStreamIhandleResetStreamFrameCall wrap *gomock.Call
type MockStreamIhandleResetStreamFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamIhandleResetStreamFrameCall) Return(arg0 error) *MockStreamIhandleResetStreamFrameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamIhandleResetStreamFrameCall) Do(f func(*wire.ResetStreamFrame) error) *MockStreamIhandleResetStreamFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamIhandleResetStreamFrameCall) DoAndReturn(f func(*wire.ResetStreamFrame) error) *MockStreamIhandleResetStreamFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// handleStopSendingFrame mocks base method.
func (m *MockStreamI) handleStopSendingFrame(arg0 *wire.StopSendingFrame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handleStopSendingFrame", arg0)
}

// handleStopSendingFrame indicates an expected call of handleStopSendingFrame.
func (mr *MockStreamIMockRecorder) handleStopSendingFrame(arg0 any) *MockStreamIhandleStopSendingFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleStopSendingFrame", reflect.TypeOf((*MockStreamI)(nil).handleStopSendingFrame), arg0)
	return &MockStreamIhandleStopSendingFrameCall{Call: call}
}

// MockStreamIhandleStopSendingFrameCall wrap *gomock.Call
type MockStreamIhandleStopSendingFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamIhandleStopSendingFrameCall) Return() *MockStreamIhandleStopSendingFrameCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamIhandleStopSendingFrameCall) Do(f func(*wire.StopSendingFrame)) *MockStreamIhandleStopSendingFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamIhandleStopSendingFrameCall) DoAndReturn(f func(*wire.StopSendingFrame)) *MockStreamIhandleStopSendingFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// handleStreamFrame mocks base method.
func (m *MockStreamI) handleStreamFrame(arg0 *wire.StreamFrame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleStreamFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleStreamFrame indicates an expected call of handleStreamFrame.
func (mr *MockStreamIMockRecorder) handleStreamFrame(arg0 any) *MockStreamIhandleStreamFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleStreamFrame", reflect.TypeOf((*MockStreamI)(nil).handleStreamFrame), arg0)
	return &MockStreamIhandleStreamFrameCall{Call: call}
}

// MockStreamIhandleStreamFrameCall wrap *gomock.Call
type MockStreamIhandleStreamFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamIhandleStreamFrameCall) Return(arg0 error) *MockStreamIhandleStreamFrameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamIhandleStreamFrameCall) Do(f func(*wire.StreamFrame) error) *MockStreamIhandleStreamFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamIhandleStreamFrameCall) DoAndReturn(f func(*wire.StreamFrame) error) *MockStreamIhandleStreamFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// hasData mocks base method.
func (m *MockStreamI) hasData() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "hasData")
	ret0, _ := ret[0].(bool)
	return ret0
}

// hasData indicates an expected call of hasData.
func (mr *MockStreamIMockRecorder) hasData() *MockStreamIhasDataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "hasData", reflect.TypeOf((*MockStreamI)(nil).hasData))
	return &MockStreamIhasDataCall{Call: call}
}

// MockStreamIhasDataCall wrap *gomock.Call
type MockStreamIhasDataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamIhasDataCall) Return(arg0 bool) *MockStreamIhasDataCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamIhasDataCall) Do(f func() bool) *MockStreamIhasDataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamIhasDataCall) DoAndReturn(f func() bool) *MockStreamIhasDataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// popStreamFrame mocks base method.
func (m *MockStreamI) popStreamFrame(arg0 protocol.ByteCount, arg1 protocol.Version) (ackhandler.StreamFrame, bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "popStreamFrame", arg0, arg1)
	ret0, _ := ret[0].(ackhandler.StreamFrame)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// popStreamFrame indicates an expected call of popStreamFrame.
func (mr *MockStreamIMockRecorder) popStreamFrame(arg0, arg1 any) *MockStreamIpopStreamFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "popStreamFrame", reflect.TypeOf((*MockStreamI)(nil).popStreamFrame), arg0, arg1)
	return &MockStreamIpopStreamFrameCall{Call: call}
}

// MockStreamIpopStreamFrameCall wrap *gomock.Call
type MockStreamIpopStreamFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamIpopStreamFrameCall) Return(arg0 ackhandler.StreamFrame, arg1, arg2 bool) *MockStreamIpopStreamFrameCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamIpopStreamFrameCall) Do(f func(protocol.ByteCount, protocol.Version) (ackhandler.StreamFrame, bool, bool)) *MockStreamIpopStreamFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamIpopStreamFrameCall) DoAndReturn(f func(protocol.ByteCount, protocol.Version) (ackhandler.StreamFrame, bool, bool)) *MockStreamIpopStreamFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// updateSendWindow mocks base method.
func (m *MockStreamI) updateSendWindow(arg0 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "updateSendWindow", arg0)
}

// updateSendWindow indicates an expected call of updateSendWindow.
func (mr *MockStreamIMockRecorder) updateSendWindow(arg0 any) *MockStreamIupdateSendWindowCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateSendWindow", reflect.TypeOf((*MockStreamI)(nil).updateSendWindow), arg0)
	return &MockStreamIupdateSendWindowCall{Call: call}
}

// MockStreamIupdateSendWindowCall wrap *gomock.Call
type MockStreamIupdateSendWindowCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamIupdateSendWindowCall) Return() *MockStreamIupdateSendWindowCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamIupdateSendWindowCall) Do(f func(protocol.ByteCount)) *MockStreamIupdateSendWindowCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamIupdateSendWindowCall) DoAndReturn(f func(protocol.ByteCount)) *MockStreamIupdateSendWindowCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
