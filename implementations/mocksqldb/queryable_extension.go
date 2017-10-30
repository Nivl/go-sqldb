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

// ExpectGet is a helper that expects a Get
func (mr *MockQueryableMockRecorder) ExpectGet(typ string, runnable func()) *gomock.Call {
	getCall := mr.Get(gomock.Eq(typ), StringType, StringType)
	getCall.Return(nil)
	if runnable != nil {
		getCall.Do(runnable)
	}
	return getCall.Times(1)
}

// ExpectGetID is a helper that expects a Get with a specific ID
func (mr *MockQueryableMockRecorder) ExpectGetID(typ string, uuid string, runnable func()) *gomock.Call {
	getCall := mr.Get(gomock.Eq(typ), StringType, uuid)
	getCall.Return(nil)
	if runnable != nil {
		getCall.Do(runnable)
	}
	return getCall.Times(1)
}

// ExpectGetNotFound is a helper that expects a not found on a Get
func (mr *MockQueryableMockRecorder) ExpectGetNotFound(typ string) *gomock.Call {
	getCall := mr.Get(gomock.Eq(typ), StringType, StringType)
	getCall.Return(sql.ErrNoRows)
	return getCall.Times(1)
}

// ExpectGetIDNotFound is a helper that expects a not found on a Get with a specific ID
func (mr *MockQueryableMockRecorder) ExpectGetIDNotFound(typ string, uuid string) *gomock.Call {
	getCall := mr.Get(gomock.Eq(typ), StringType, uuid)
	getCall.Return(sql.ErrNoRows)
	return getCall.Times(1)
}

// ExpectGetIDError is a helper that expects a connection error on a Get with a specific ID
func (mr *MockQueryableMockRecorder) ExpectGetIDError(typ string, uuid string, err error) *gomock.Call {
	getCall := mr.Get(gomock.Eq(typ), StringType, uuid)
	getCall.Return(err)
	return getCall.Times(1)
}

// ExpectGetNoParams is a helper that expects a Get with no params but the stmt
func (mr *MockQueryableMockRecorder) ExpectGetNoParams(typ string, runnable func()) *gomock.Call {
	call := mr.Get(gomock.Eq(typ), StringType)
	call.Return(nil)
	if runnable != nil {
		call.Do(runnable)
	}
	return call.Times(1)
}

// ExpectGetNoParamsNotFound is a helper that expects a not found on a Get with no params but the stmt
func (mr *MockQueryableMockRecorder) ExpectGetNoParamsNotFound(typ string) *gomock.Call {
	call := mr.Get(gomock.Eq(typ), StringType)
	call.Return(sql.ErrNoRows)
	return call.Times(1)
}

// ExpectGetNoParamsError is a helper that expects a connection error on a Get with no params but the stmt
func (mr *MockQueryableMockRecorder) ExpectGetNoParamsError(typ string, err error) *gomock.Call {
	call := mr.Get(gomock.Eq(typ), StringType)
	call.Return(err)
	return call.Times(1)
}

// ExpectGetError is a helper that expects a connection error on a Get
func (mr *MockQueryableMockRecorder) ExpectGetError(typ string, err error) *gomock.Call {
	getCall := mr.Get(gomock.Eq(typ), StringType, StringType)
	getCall.Return(err)
	return getCall.Times(1)
}

// ExpectSelect is an helper that expects a connection error on a Select
func (mr *MockQueryableMockRecorder) ExpectSelect(typ string, runnable func()) *gomock.Call {
	selectCall := mr.Select(gomock.Eq(typ), StringType, IntType, IntType)
	selectCall.Return(nil)
	if runnable != nil {
		selectCall.Do(runnable)
	}
	return selectCall.Times(1)
}

// ExpectSelectError is an helper that expects a Select
func (mr *MockQueryableMockRecorder) ExpectSelectError(typ string, err error) *gomock.Call {
	selectCall := mr.Select(gomock.Eq(typ), StringType, IntType, IntType)
	selectCall.Return(err)
	return selectCall.Times(1)
}

// ExpectDeletion is a helper that expects a deletion
func (mr *MockQueryableMockRecorder) ExpectDeletion() *gomock.Call {
	return mr.Exec(StringType, StringType).Return(int64(1), nil).Times(1)
}

// ExpectDeletionError is a helper that expects a deletion to fail
func (mr *MockQueryableMockRecorder) ExpectDeletionError(err error) *gomock.Call {
	return mr.Exec(StringType, StringType).Return(int64(0), err).Times(1)
}

// ExpectInsert is a helper that expects an insertion
func (mr *MockQueryableMockRecorder) ExpectInsert(typ string) *gomock.Call {
	return mr.NamedExec(StringType, gomock.Eq(typ)).Return(int64(1), nil).Times(1)
}

// ExpectInsertError is a helper that expects an insert to fail
func (mr *MockQueryableMockRecorder) ExpectInsertError(typ string, err error) *gomock.Call {
	return mr.NamedExec(StringType, gomock.Eq(typ)).Return(int64(1), err).Times(1)
}

// ExpectInsertConflict is a helper that expects a conflict on an insertion
func (mr *MockQueryableMockRecorder) ExpectInsertConflict(typ string, err error) *gomock.Call {
	return mr.NamedExec(StringType, gomock.Eq(typ)).Return(int64(0), err).Times(1)
}

// ExpectUpdate is a helper that expects an update
func (mr *MockQueryableMockRecorder) ExpectUpdate(typ string) *gomock.Call {
	return mr.NamedExec(StringType, gomock.Eq(typ)).Return(int64(1), nil).Times(1)
}

// ExpectUpdateConflict is a helper that expects a conflict on an update
func (mr *MockQueryableMockRecorder) ExpectUpdateConflict(typ string, err error) *gomock.Call {
	return mr.NamedExec(StringType, gomock.Eq(typ)).Return(int64(0), err).Times(1)
}

// ExpectUpdateError is a helper that expects an update to fail
func (mr *MockQueryableMockRecorder) ExpectUpdateError(typ string, err error) *gomock.Call {
	return mr.NamedExec(StringType, gomock.Eq(typ)).Return(int64(0), err).Times(1)
}
