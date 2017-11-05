package mocksqldb

import (
	sql "database/sql"
	reflect "reflect"

	go_sqldb "github.com/Nivl/go-sqldb"
	gomock "github.com/golang/mock/gomock"
)

// MockConnection is a mock of Connection interface
type MockConnection struct {
	queryable *MockQueryable

	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{
		ctrl:      ctrl,
		queryable: NewMockQueryable(ctrl),
	}
	mock.recorder = &MockConnectionMockRecorder{
		mock: mock,
	}
	return mock
}

// QEXPECT returns an object that allows the caller to indicate expected use
// for a Queryable
func (m *MockConnection) QEXPECT() *MockQueryableMockRecorder {
	return m.queryable.recorder
}

// Q returns a queryable mock
func (m *MockConnection) Q() *MockQueryable {
	return m.queryable
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Beginx mocks base method
func (m *MockConnection) Beginx() (go_sqldb.Tx, error) {
	ret := m.ctrl.Call(m, "Beginx")
	ret0, _ := ret[0].(go_sqldb.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Beginx indicates an expected call of Beginx
func (mr *MockConnectionMockRecorder) Beginx() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Beginx", reflect.TypeOf((*MockConnection)(nil).Beginx))
}

// Close mocks base method
func (m *MockConnection) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// DSN mocks base method
func (m *MockConnection) DSN() string {
	ret := m.ctrl.Call(m, "DSN")
	ret0, _ := ret[0].(string)
	return ret0
}

// DSN indicates an expected call of DSN
func (mr *MockConnectionMockRecorder) DSN() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DSN", reflect.TypeOf((*MockConnection)(nil).DSN))
}

// SQL mocks base method
func (m *MockConnection) SQL() *sql.DB {
	ret := m.ctrl.Call(m, "SQL")
	ret0, _ := ret[0].(*sql.DB)
	return ret0
}

// SQL indicates an expected call of SQL
func (mr *MockConnectionMockRecorder) SQL() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SQL", reflect.TypeOf((*MockConnection)(nil).SQL))
}
