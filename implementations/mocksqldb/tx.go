package mocksqldb

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTx is a mock of Tx interface
type MockTx struct {
	queryable *MockQueryable

	ctrl     *gomock.Controller
	recorder *MockTxMockRecorder
}

// MockTxMockRecorder is the mock recorder for MockTx
type MockTxMockRecorder struct {
	mock *MockTx

	MockQueryableMockRecorder
}

// NewMockTx creates a new mock instance
func NewMockTx(ctrl *gomock.Controller) *MockTx {
	mock := &MockTx{
		ctrl:      ctrl,
		queryable: NewMockQueryable(ctrl),
	}
	mock.recorder = &MockTxMockRecorder{mock: mock}
	return mock
}

// QEXPECT returns an object that allows the caller to indicate expected use
// for a Queryable
func (m *MockTx) QEXPECT() *MockQueryableMockRecorder {
	return m.queryable.recorder
}

// Q returns a queryable mock
func (m *MockTx) Q() *MockQueryable {
	return m.queryable
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTx) EXPECT() *MockTxMockRecorder {
	return m.recorder
}

// Commit mocks base method
func (m *MockTx) Commit() error {
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockTxMockRecorder) Commit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTx)(nil).Commit))
}

// Rollback mocks base method
func (m *MockTx) Rollback() error {
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback
func (mr *MockTxMockRecorder) Rollback() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTx)(nil).Rollback))
}
