package database

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestConnector(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	obj := &connector{
		openConnection: func(driverName, dataSourceName string) (*sqlx.DB, error) {
			return sqlxDB, nil
		},
	}
	obj.Connect(&Option{
		MasterConnectionString: "",
		SlaveConnectionString:  "",
	})
}

func TestConnector_FailMasterConnection(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	obj := &connector{
		openConnection: func(driverName, dataSourceName string) (*sqlx.DB, error) {
			if dataSourceName == "master" {
				return nil, errors.New("failed")
			}
			return sqlxDB, nil
		},
	}
	_, err = obj.Connect(&Option{
		MasterConnectionString: "master",
		SlaveConnectionString:  "slave",
	})
	require.Error(t, err)
}
func TestConnector_FailSlaveConnection(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	obj := &connector{
		openConnection: func(driverName, dataSourceName string) (*sqlx.DB, error) {
			if dataSourceName == "slave" {
				return nil, errors.New("failed")
			}
			return sqlxDB, nil
		},
	}
	_, err = obj.Connect(&Option{
		MasterConnectionString: "master",
		SlaveConnectionString:  "slave",
	})
	require.Error(t, err)
}

func TestNewConnector(t *testing.T) {
	obj := NewConnector()
	require.NotNil(t, obj)
}
