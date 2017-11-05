package mocksqldb_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Nivl/go-sqldb/implementations/mocksqldb"
	gomock "github.com/golang/mock/gomock"
)

type testStruct struct {
	Name  string
	Value int
}

func TestGetSuccess(t *testing.T) {
	t.Run("no runnable", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocksqldb.NewMockQueryable(mockCtrl)
		m.EXPECT().GetSuccess(&testStruct{}, nil)

		output := &testStruct{}
		err := m.Get(output, "select * ...", "param")
		assert.NoError(t, err, "m.Get() should have worked")
	})

	t.Run("with runnable", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocksqldb.NewMockQueryable(mockCtrl)
		m.EXPECT().GetSuccess(&testStruct{}, func(output *testStruct, stmt, params string) {
			output.Name = "New name"
			output.Value = 42
		})

		output := &testStruct{}
		err := m.Get(output, "select * ...", "param")
		assert.NoError(t, err, "m.Get() should have worked")
		assert.Equal(t, "New name", output.Name, "wrong value for output.Name")
		assert.Equal(t, 42, output.Value, "wrong value for output.Value")
	})
}

func TestGetID(t *testing.T) {
	t.Run("no runnable", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocksqldb.NewMockQueryable(mockCtrl)
		m.EXPECT().GetID(&testStruct{}, "uuid", nil)

		output := &testStruct{}
		err := m.Get(output, "select * ...", "uuid")
		assert.NoError(t, err, "m.Get() should have worked")
	})

	t.Run("with runnable", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocksqldb.NewMockQueryable(mockCtrl)
		m.EXPECT().GetID(&testStruct{}, "uuid", func(output *testStruct, stmt, uuid string) {
			output.Name = "New name"
			output.Value = 42
			assert.Equal(t, "uuid", uuid, "wrong value for uuid")
		})

		output := &testStruct{}
		err := m.Get(output, "select * ...", "uuid")
		assert.NoError(t, err, "m.Get() should have worked")
		assert.Equal(t, "New name", output.Name, "wrong value for output.Name")
		assert.Equal(t, 42, output.Value, "wrong value for output.Value")
	})
}

func TestGetNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().GetNotFound(&testStruct{})

	output := &testStruct{}
	err := m.Get(output, "select * ...", "uuid")
	assert.Error(t, err, "m.Get() should have failed")
}

func TestGetIDNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().GetIDNotFound(&testStruct{}, "uuid")

	output := &testStruct{}
	err := m.Get(output, "select * ...", "uuid")
	assert.Error(t, err, "m.Get() should have failed")
}

func TestGetIDError(t *testing.T) {
	expectedErr := errors.New("test error")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().GetIDError(&testStruct{}, "uuid", expectedErr)

	output := &testStruct{}
	err := m.Get(output, "select * ...", "uuid")
	assert.Error(t, err, "m.Get() should have failed")
	assert.Equal(t, expectedErr, err, "wrong error returned")
}

func TestGetNoParams(t *testing.T) {
	t.Run("no runnable", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocksqldb.NewMockQueryable(mockCtrl)
		m.EXPECT().GetNoParams(&testStruct{}, nil)

		output := &testStruct{}
		err := m.Get(output, "select * ...")
		assert.NoError(t, err, "m.Get() should have worked")
	})

	t.Run("with runnable", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocksqldb.NewMockQueryable(mockCtrl)
		m.EXPECT().GetNoParams(&testStruct{}, func(output *testStruct, stmt string) {
			output.Name = "New name"
			output.Value = 42
		})

		output := &testStruct{}
		err := m.Get(output, "select * ...")
		assert.NoError(t, err, "m.Get() should have worked")
		assert.Equal(t, "New name", output.Name, "wrong value for output.Name")
		assert.Equal(t, 42, output.Value, "wrong value for output.Value")
	})
}

func TestGetNoParamsNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().GetNoParamsNotFound(&testStruct{})

	output := &testStruct{}
	err := m.Get(output, "select * ...")
	assert.Error(t, err, "m.Get() should have failed")
}

func TestGetNoParamsError(t *testing.T) {
	expectedErr := errors.New("test error")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().GetNoParamsError(&testStruct{}, expectedErr)

	output := &testStruct{}
	err := m.Get(output, "select * ...")
	assert.Error(t, err, "m.Get() should have failed")
	assert.Equal(t, expectedErr, err, "wrong error returned")
}

func TestGetError(t *testing.T) {
	expectedErr := errors.New("test error")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().GetError(&testStruct{}, expectedErr)

	output := &testStruct{}
	err := m.Get(output, "select * ...", "param")
	assert.Error(t, err, "m.Get() should have failed")
	assert.Equal(t, expectedErr, err, "wrong error returned")
}

func TestSelectSuccess(t *testing.T) {
	t.Run("no runnable", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocksqldb.NewMockQueryable(mockCtrl)
		m.EXPECT().SelectSuccess([]*testStruct{}, nil)

		output := []*testStruct{}
		err := m.Select(output, "select * ...", 1, 2)
		assert.NoError(t, err, "m.Select() should have worked")
	})

	t.Run("with runnable", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocksqldb.NewMockQueryable(mockCtrl)
		m.EXPECT().SelectSuccess([]*testStruct{}, func(output []*testStruct, stmt string, page, perPage int) {})

		var output []*testStruct
		err := m.Select(output, "select * ...", 1, 2)
		require.NoError(t, err, "m.Select() should have worked")
	})
}

func TestSelectError(t *testing.T) {
	expectedErr := errors.New("test error")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().SelectError([]*testStruct{}, expectedErr)

	var output []*testStruct
	err := m.Select(output, "select * ...", 1, 2)
	assert.Error(t, err, "m.Select() should have failed")
	assert.Equal(t, expectedErr, err, "wrong error returned")
}

func TestDeletionSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().DeletionSuccess()

	_, err := m.Exec("delete ...", "id")
	assert.NoError(t, err, "m.Exec() should have worked")
}

func TestDeletionError(t *testing.T) {
	expectedErr := errors.New("test error")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().DeletionError(expectedErr)

	_, err := m.Exec("select * ...", "id")
	assert.Error(t, err, "m.Select() should have failed")
	assert.Equal(t, expectedErr, err, "wrong error returned")
}

func TestInsertSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().InsertSuccess(&testStruct{})

	_, err := m.NamedExec("insert into ..", &testStruct{})
	assert.NoError(t, err, "m.NamedExec() should have worked")
}

func TestInsertError(t *testing.T) {
	expectedErr := errors.New("test error")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().InsertError(&testStruct{}, expectedErr)

	_, err := m.NamedExec("insert into ..", &testStruct{})
	assert.Error(t, err, "m.NamedExec() should have failed")
	assert.Equal(t, expectedErr, err, "wrong error returned")
}

func TestUpdateSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().UpdateSuccess(&testStruct{})

	_, err := m.NamedExec("update ..", &testStruct{})
	assert.NoError(t, err, "m.NamedExec() should have worked")
}

func TestUpdateError(t *testing.T) {
	expectedErr := errors.New("test error")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockQueryable(mockCtrl)
	m.EXPECT().UpdateError(&testStruct{}, expectedErr)

	_, err := m.NamedExec("update ..", &testStruct{})
	assert.Error(t, err, "m.NamedExec() should have failed")
	assert.Equal(t, expectedErr, err, "wrong error returned")
}
