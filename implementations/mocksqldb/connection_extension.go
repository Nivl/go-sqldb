package mocksqldb

import (
	"errors"

	gomock "github.com/golang/mock/gomock"
)

// TransactionSuccess is a helper that expects a transaction
func (mr *MockConnectionMockRecorder) TransactionSuccess(ctrl *gomock.Controller) (*MockTx, *gomock.Call) {
	tx := NewMockTx(ctrl)
	call := mr.Beginx().Return(tx, nil)
	return tx, call
}

// TransactionError is a helper that expects a transaction to fail
func (mr *MockConnectionMockRecorder) TransactionError() *gomock.Call {
	return mr.Beginx().Return(nil, errors.New("could not create transaction"))
}
