package database

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type connector struct {
	openConnection OpenConnection
}
type Connector interface {
	Connect(o *Option) (*Store, error)
}

type Option struct {
	MasterConnectionString string
	SlaveConnectionString  string
	MaxIdleConn            int
	MaxConn                int
}

func NewConnector() Connector {
	return &connector{
		openConnection: sqlx.Open,
	}
}

type ExecFunc func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
type OpenConnection func(driverName, dataSourceName string) (*sqlx.DB, error)

const MySQLDatabaseDriver = "mysql"

func (d *connector) Connect(o *Option) (*Store, error) {
	master, err := d.openConnection(MySQLDatabaseDriver, o.MasterConnectionString)
	if err != nil {
		return nil, err
	}
	master.SetMaxOpenConns(o.MaxConn)
	master.SetMaxIdleConns(o.MaxIdleConn)
	slave, err := d.openConnection(MySQLDatabaseDriver, o.SlaveConnectionString)
	if err != nil {
		return nil, err
	}
	slave.SetMaxOpenConns(o.MaxConn)
	slave.SetMaxIdleConns(o.MaxIdleConn)

	return &Store{
		Master: master,
		Slave:  slave,
	}, nil
}
