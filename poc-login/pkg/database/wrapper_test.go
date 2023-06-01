package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	obj := Store{
		Master: sqlxDB,
		Slave:  sqlxDB,
	}
	master := obj.GetMaster()
	require.Equal(t, sqlxDB, master)
	slave := obj.GetSlave()
	require.Equal(t, sqlxDB, slave)
	executor := obj.GetExecContext(master)
	require.Equal(t, executor, master)
	executor2 := obj.GetExecContext(nil)
	require.Equal(t, executor2, master)
	selector := obj.GetSelectContext(slave)
	require.Equal(t, slave, selector)
	selector2 := obj.GetSelectContext(nil)
	require.Equal(t, slave, selector2)
}
