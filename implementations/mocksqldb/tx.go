package mocksqldb

import (
	reflect "reflect"

	sqldb "github.com/Nivl/go-sqldb"
	gomock "github.com/golang/mock/gomock"
)

var _ sqldb.Tx = (*MockTx)(nil)

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

// Exec mocks base method
func (m *MockTx) Exec(arg0 string, arg1 ...interface{}) (int64, error) {
	return m.queryable.Exec(arg0, arg1...)
}

// Get mocks base method
func (m *MockTx) Get(arg0 interface{}, arg1 string, arg2 ...interface{}) error {
	return m.queryable.Get(arg0, arg1, arg2...)
}

// NamedExec mocks base method
func (m *MockTx) NamedExec(arg0 string, arg1 interface{}) (int64, error) {
	return m.queryable.NamedExec(arg0, arg1)
}

// NamedGet indicates an expected call of NamedGet
func (m *MockTx) NamedGet(arg0 interface{}, arg1 string, arg2 interface{}) error {
	return m.queryable.NamedGet(arg0, arg1, arg2)
}

// NamedSelect mocks base method
func (m *MockTx) NamedSelect(arg0 interface{}, arg1 string, arg2 interface{}) error {
	return m.queryable.NamedSelect(arg0, arg1, arg2)
}

// Select mocks base method
func (m *MockTx) Select(arg0 interface{}, arg1 string, arg2 ...interface{}) error {
	return m.queryable.Select(arg0, arg1, arg2...)
}
