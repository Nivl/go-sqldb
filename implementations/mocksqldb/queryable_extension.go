package mocksqldb

import (
	"database/sql"

	gomock "github.com/golang/mock/gomock"
)

var (
	// StringType represents a string argument
	StringType = gomock.Eq("string")
	// IntType represents an int argument
	IntType = gomock.Eq("int")
	// AnyType represents an argument that can be anything
	AnyType = gomock.Any()
)

// GetSuccess is a helper that expects a Get to succeed
func (mr *MockQueryableMockRecorder) GetSuccess(typ string, runnable func()) *gomock.Call {
	getCall := mr.Get(gomock.Eq(typ), StringType, StringType)
	getCall.Return(nil)
	if runnable != nil {
		getCall.Do(runnable)
	}
	return getCall.Times(1)
}

// GetID is a helper that expects a Get with a specific ID
func (mr *MockQueryableMockRecorder) GetID(typ string, uuid string, runnable func()) *gomock.Call {
	getCall := mr.Get(gomock.Eq(typ), StringType, uuid)
	getCall.Return(nil)
	if runnable != nil {
		getCall.Do(runnable)
	}
	return getCall.Times(1)
}

// GetNotFound is a helper that expects a not found on a Get
func (mr *MockQueryableMockRecorder) GetNotFound(typ string) *gomock.Call {
	getCall := mr.Get(gomock.Eq(typ), StringType, StringType)
	getCall.Return(sql.ErrNoRows)
	return getCall.Times(1)
}

// GetIDNotFound is a helper that expects a not found on a Get with a specific ID
func (mr *MockQueryableMockRecorder) GetIDNotFound(typ string, uuid string) *gomock.Call {
	getCall := mr.Get(gomock.Eq(typ), StringType, uuid)
	getCall.Return(sql.ErrNoRows)
	return getCall.Times(1)
}

// GetIDError is a helper that expects a connection error on a Get with a specific ID
func (mr *MockQueryableMockRecorder) GetIDError(typ string, uuid string, err error) *gomock.Call {
	getCall := mr.Get(gomock.Eq(typ), StringType, uuid)
	getCall.Return(err)
	return getCall.Times(1)
}

// GetNoParams is a helper that expects a Get with no params but the stmt
func (mr *MockQueryableMockRecorder) GetNoParams(typ string, runnable func()) *gomock.Call {
	call := mr.Get(gomock.Eq(typ), StringType)
	call.Return(nil)
	if runnable != nil {
		call.Do(runnable)
	}
	return call.Times(1)
}

// GetNoParamsNotFound is a helper that expects a not found on a Get with no params but the stmt
func (mr *MockQueryableMockRecorder) GetNoParamsNotFound(typ string) *gomock.Call {
	call := mr.Get(gomock.Eq(typ), StringType)
	call.Return(sql.ErrNoRows)
	return call.Times(1)
}

// GetNoParamsError is a helper that expects a connection error on a Get with no params but the stmt
func (mr *MockQueryableMockRecorder) GetNoParamsError(typ string, err error) *gomock.Call {
	call := mr.Get(gomock.Eq(typ), StringType)
	call.Return(err)
	return call.Times(1)
}

// GetError is a helper that expects a connection error on a Get
func (mr *MockQueryableMockRecorder) GetError(typ string, err error) *gomock.Call {
	getCall := mr.Get(gomock.Eq(typ), StringType, StringType)
	getCall.Return(err)
	return getCall.Times(1)
}

// SelectSuccess is an helper that expects a Select
func (mr *MockQueryableMockRecorder) SelectSuccess(typ string, runnable func()) *gomock.Call {
	selectCall := mr.Select(gomock.Eq(typ), StringType, IntType, IntType)
	selectCall.Return(nil)
	if runnable != nil {
		selectCall.Do(runnable)
	}
	return selectCall.Times(1)
}

// SelectError is an helper that expects an error on a Select
func (mr *MockQueryableMockRecorder) SelectError(typ string, err error) *gomock.Call {
	selectCall := mr.Select(gomock.Eq(typ), StringType, IntType, IntType)
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
func (mr *MockQueryableMockRecorder) InsertSuccess(typ string) *gomock.Call {
	return mr.NamedExec(StringType, gomock.Eq(typ)).Return(int64(1), nil).Times(1)
}

// InsertError is a helper that expects an insert to fail
func (mr *MockQueryableMockRecorder) InsertError(typ string, err error) *gomock.Call {
	return mr.NamedExec(StringType, gomock.Eq(typ)).Return(int64(1), err).Times(1)
}

// UpdateSuccess is a helper that expects an update
func (mr *MockQueryableMockRecorder) UpdateSuccess(typ string) *gomock.Call {
	return mr.NamedExec(StringType, gomock.Eq(typ)).Return(int64(1), nil).Times(1)
}

// UpdateError is a helper that expects an update to fail
func (mr *MockQueryableMockRecorder) UpdateError(typ string, err error) *gomock.Call {
	return mr.NamedExec(StringType, gomock.Eq(typ)).Return(int64(0), err).Times(1)
}
