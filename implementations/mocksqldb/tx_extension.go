package mocksqldb

import (
	"errors"

	gomock "github.com/golang/mock/gomock"
)

// ExpectCommit is a helper that expects a transaction to Commit
func (mr *MockTxMockRecorder) ExpectCommit() *gomock.Call {
	return mr.Commit().Return(nil)
}

// ExpectCommitError is a helper that expects a commit to fail
func (mr *MockTxMockRecorder) ExpectCommitError() *gomock.Call {
	return mr.Commit().Return(errors.New("could not commit"))
}

// ExpectRollback is a helper that expects a transaction to Rollback
func (mr *MockTxMockRecorder) ExpectRollback() *gomock.Call {
	return mr.Rollback().Return(nil)
}

// ExpectRollbackError is a helper that expects a Rollback to fail
func (mr *MockTxMockRecorder) ExpectRollbackError() *gomock.Call {
	return mr.Rollback().Return(errors.New("could not Rollback"))
}
