package mocksqldb

import (
	"errors"

	gomock "github.com/golang/mock/gomock"
)

// CommitSuccess is a helper that expects a transaction to Commit
func (mr *MockTxMockRecorder) CommitSuccess() *gomock.Call {
	return mr.Commit().Return(nil)
}

// CommitError is a helper that expects a commit to fail
func (mr *MockTxMockRecorder) CommitError() *gomock.Call {
	return mr.Commit().Return(errors.New("could not commit"))
}

// RollbackSuccess is a helper that expects a transaction to Rollback
func (mr *MockTxMockRecorder) RollbackSuccess() *gomock.Call {
	return mr.Rollback().Return(nil)
}

// RollbackError is a helper that expects a Rollback to fail
func (mr *MockTxMockRecorder) RollbackError() *gomock.Call {
	return mr.Rollback().Return(errors.New("could not Rollback"))
}
