package testhelper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateGormMockRowFromStructArray(t *testing.T) {
	input := []DummyGormDAO{
		{ID: 1},
		{ID: 2},
	}
	resp := GenerateGormMockRowFromStructArray(input)
	require.NotNil(t, resp)
}
func TestGenerateGormMockRowFromStructArray_EmptyInput(t *testing.T) {
	input := []DummyGormDAO{}
	resp := GenerateGormMockRowFromStructArray(input)
	require.Nil(t, resp)
}
func TestGenerateNativeMockRowFromStructArray(t *testing.T) {
	input := []*DummyDBDAO{
		{ID: 1},
		{ID: 2},
	}
	resp := GenerateNativeMockRowFromStructArray(input)
	require.NotNil(t, resp)
}

type DummyGormDAO struct {
	ID int64 `gorm:"id"`
}
type DummyDBDAO struct {
	ID int64 `db:"id"`
}

func TestGenerateSQLMockDataFromStruct(t *testing.T) {
	var testinput *DummyDBDAO
	resp := GenerateSQLMockDataFromStruct(testinput)
	require.NotNil(t, resp)
}
func TestGenerateGormMockRowFromStructArray_err(t *testing.T) {
	var input *DummyDBDAO
	resp := GenerateGormMockRowFromStructArray(input)
	require.Nil(t, resp)
}
