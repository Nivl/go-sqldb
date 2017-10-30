package mocksqldb

import (
	"errors"

	gomock "github.com/golang/mock/gomock"
)

// ExpectTransaction is a helper that expects a transaction
func (mr *MockConnectionMockRecorder) ExpectTransaction(ctrl *gomock.Controller) (*MockTx, *gomock.Call) {
	tx := NewMockTx(ctrl)
	call := mr.Beginx().Return(tx, nil)
	return tx, call
}

// ExpectTransactionError is a helper that expects a transaction to fail
func (mr *MockConnectionMockRecorder) ExpectTransactionError() *gomock.Call {
	return mr.Beginx().Return(nil, errors.New("cound not create transaction"))
}
