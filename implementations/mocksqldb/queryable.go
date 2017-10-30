package mocksqldb

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockQueryable is a mock of Queryable interface
type MockQueryable struct {
	ctrl     *gomock.Controller
	recorder *MockQueryableMockRecorder
}

// MockQueryableMockRecorder is the mock recorder for MockQueryable
type MockQueryableMockRecorder struct {
	mock *MockQueryable
}

// NewMockQueryable creates a new mock instance
func NewMockQueryable(ctrl *gomock.Controller) *MockQueryable {
	mock := &MockQueryable{ctrl: ctrl}
	mock.recorder = &MockQueryableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQueryable) EXPECT() *MockQueryableMockRecorder {
	return m.recorder
}

// Exec mocks base method
func (m *MockQueryable) Exec(arg0 string, arg1 ...interface{}) (int64, error) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec
func (mr *MockQueryableMockRecorder) Exec(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockQueryable)(nil).Exec), varargs...)
}

// Get mocks base method
func (m *MockQueryable) Get(arg0 interface{}, arg1 string, arg2 ...interface{}) error {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockQueryableMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockQueryable)(nil).Get), varargs...)
}

// NamedExec mocks base method
func (m *MockQueryable) NamedExec(arg0 string, arg1 interface{}) (int64, error) {
	ret := m.ctrl.Call(m, "NamedExec", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedExec indicates an expected call of NamedExec
func (mr *MockQueryableMockRecorder) NamedExec(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedExec", reflect.TypeOf((*MockQueryable)(nil).NamedExec), arg0, arg1)
}

// NamedGet mocks base method
func (m *MockQueryable) NamedGet(arg0 interface{}, arg1 string, arg2 interface{}) error {
	ret := m.ctrl.Call(m, "NamedGet", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// NamedGet indicates an expected call of NamedGet
func (mr *MockQueryableMockRecorder) NamedGet(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedGet", reflect.TypeOf((*MockQueryable)(nil).NamedGet), arg0, arg1, arg2)
}

// NamedSelect mocks base method
func (m *MockQueryable) NamedSelect(arg0 interface{}, arg1 string, arg2 interface{}) error {
	ret := m.ctrl.Call(m, "NamedSelect", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// NamedSelect indicates an expected call of NamedSelect
func (mr *MockQueryableMockRecorder) NamedSelect(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedSelect", reflect.TypeOf((*MockQueryable)(nil).NamedSelect), arg0, arg1, arg2)
}

// Select mocks base method
func (m *MockQueryable) Select(arg0 interface{}, arg1 string, arg2 ...interface{}) error {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Select indicates an expected call of Select
func (mr *MockQueryableMockRecorder) Select(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockQueryable)(nil).Select), varargs...)
}
