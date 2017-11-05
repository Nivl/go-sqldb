package mocksqldb

import (
	"database/sql"

	matcher "github.com/Nivl/gomock-type-matcher"
	gomock "github.com/golang/mock/gomock"
)

var (
	// StringType represents a string argument
	StringType = matcher.String()
	// IntType represents an int argument
	IntType = matcher.Int()
	// AnyType represents an argument that can be anything
	AnyType = gomock.Any()
)

// GetSuccess is a helper that expects a Get to succeed
func (mr *MockQueryableMockRecorder) GetSuccess(typ interface{}, runnable interface{}) *gomock.Call {
	getCall := mr.Get(matcher.Interface(typ), StringType, StringType)
	getCall.Return(nil)
	if runnable != nil {
		getCall.Do(runnable)
	}
	return getCall.Times(1)
}

// GetID is a helper that expects a Get with a specific ID
func (mr *MockQueryableMockRecorder) GetID(typ interface{}, uuid string, runnable interface{}) *gomock.Call {
	getCall := mr.Get(matcher.Interface(typ), StringType, uuid)
	getCall.Return(nil)
	if runnable != nil {
		getCall.Do(runnable)
	}
	return getCall.Times(1)
}

// GetNotFound is a helper that expects a not found on a Get
func (mr *MockQueryableMockRecorder) GetNotFound(typ interface{}) *gomock.Call {
	getCall := mr.Get(matcher.Interface(typ), StringType, StringType)
	getCall.Return(sql.ErrNoRows)
	return getCall.Times(1)
}

// GetIDNotFound is a helper that expects a not found on a Get with a specific ID
func (mr *MockQueryableMockRecorder) GetIDNotFound(typ interface{}, uuid string) *gomock.Call {
	getCall := mr.Get(matcher.Interface(typ), StringType, uuid)
	getCall.Return(sql.ErrNoRows)
	return getCall.Times(1)
}

// GetIDError is a helper that expects a connection error on a Get with a specific ID
func (mr *MockQueryableMockRecorder) GetIDError(typ interface{}, uuid string, err error) *gomock.Call {
	getCall := mr.Get(matcher.Interface(typ), StringType, uuid)
	getCall.Return(err)
	return getCall.Times(1)
}

// GetNoParams is a helper that expects a Get with no params but the stmt
func (mr *MockQueryableMockRecorder) GetNoParams(typ interface{}, runnable interface{}) *gomock.Call {
	call := mr.Get(matcher.Interface(typ), StringType)
	call.Return(nil)
	if runnable != nil {
		call.Do(runnable)
	}
	return call.Times(1)
}

// GetNoParamsNotFound is a helper that expects a not found on a Get with no params but the stmt
func (mr *MockQueryableMockRecorder) GetNoParamsNotFound(typ interface{}) *gomock.Call {
	call := mr.Get(matcher.Interface(typ), StringType)
	call.Return(sql.ErrNoRows)
	return call.Times(1)
}

// GetNoParamsError is a helper that expects a connection error on a Get with no params but the stmt
func (mr *MockQueryableMockRecorder) GetNoParamsError(typ interface{}, err error) *gomock.Call {
	call := mr.Get(matcher.Interface(typ), StringType)
	call.Return(err)
	return call.Times(1)
}

// GetError is a helper that expects a connection error on a Get
func (mr *MockQueryableMockRecorder) GetError(typ interface{}, err error) *gomock.Call {
	getCall := mr.Get(matcher.Interface(typ), StringType, StringType)
	getCall.Return(err)
	return getCall.Times(1)
}

// SelectSuccess is an helper that expects a Select
func (mr *MockQueryableMockRecorder) SelectSuccess(typ interface{}, runnable interface{}) *gomock.Call {
	selectCall := mr.Select(matcher.Interface(typ), StringType, IntType, IntType)
	selectCall.Return(nil)
	if runnable != nil {
		selectCall.Do(runnable)
	}
	return selectCall.Times(1)
}

// SelectError is an helper that expects an error on a Select
func (mr *MockQueryableMockRecorder) SelectError(typ interface{}, err error) *gomock.Call {
	selectCall := mr.Select(matcher.Interface(typ), StringType, IntType, IntType)
	selectCall.Return(err)
	return selectCall.Times(1)
}

// DeletionSuccess is a helper that expects a deletion to succeed
func (mr *MockQueryableMockRecorder) DeletionSuccess() *gomock.Call {
	return mr.Exec(StringType, StringType).Return(int64(1), nil).Times(1)
}

// DeletionError is a helper that expects a deletion to fail
func (mr *MockQueryableMockRecorder) DeletionError(err error) *gomock.Call {
	return mr.Exec(StringType, StringType).Return(int64(0), err).Times(1)
}

// InsertSuccess is a helper that expects an insertion
func (mr *MockQueryableMockRecorder) InsertSuccess(typ interface{}) *gomock.Call {
	return mr.NamedExec(StringType, matcher.Interface(typ)).Return(int64(1), nil).Times(1)
}

// InsertError is a helper that expects an insert to fail
func (mr *MockQueryableMockRecorder) InsertError(typ interface{}, err error) *gomock.Call {
	return mr.NamedExec(StringType, matcher.Interface(typ)).Return(int64(1), err).Times(1)
}

// UpdateSuccess is a helper that expects an update
func (mr *MockQueryableMockRecorder) UpdateSuccess(typ interface{}) *gomock.Call {
	return mr.NamedExec(StringType, matcher.Interface(typ)).Return(int64(1), nil).Times(1)
}

// UpdateError is a helper that expects an update to fail
func (mr *MockQueryableMockRecorder) UpdateError(typ interface{}, err error) *gomock.Call {
	return mr.NamedExec(StringType, matcher.Interface(typ)).Return(int64(0), err).Times(1)
}
