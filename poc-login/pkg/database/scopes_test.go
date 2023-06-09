package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestPaginateMaximum(t *testing.T) {
	db, dbmock, err := sqlmock.New()
	require.NoError(t, err)
	dbmock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow(1.0))
	dbmock.ExpectQuery("SELECT sale_order.id FROM `sale_order` LIMIT 100").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)
	rows, err := gormDB.Table("sale_order").
		Scopes(Paginate(200, 20000)).
		Select(`sale_order.id`).
		Rows()
	require.NoError(t, err)
	defer rows.Close()
	resp := []*dummyStruct{}
	for rows.Next() {
		err := gormDB.ScanRows(rows, &resp)
		require.NoError(t, err)
	}
}
func TestPaginateMinimum(t *testing.T) {
	db, dbmock, err := sqlmock.New()
	require.NoError(t, err)
	dbmock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow(1.0))
	dbmock.ExpectQuery("SELECT sale_order.id FROM `sale_order` LIMIT 10").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)
	rows, err := gormDB.Table("sale_order").
		Scopes(Paginate(0, 5)).
		Select(`sale_order.id`).
		Rows()
	require.NoError(t, err)
	defer rows.Close()
	resp := []*dummyStruct{}
	for rows.Next() {
		err := gormDB.ScanRows(rows, &resp)
		require.NoError(t, err)
	}
}

type dummyStruct struct {
	ID int
}

func TestEqualInt64(t *testing.T) {
	db, dbmock, err := sqlmock.New()
	require.NoError(t, err)
	dbmock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow(1.0))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)
	fn := EqualInt64("field", 0)
	result := fn(gormDB)
	require.NotNil(t, result)
}

func TestEqualString(t *testing.T) {
	db, dbmock, err := sqlmock.New()
	require.NoError(t, err)
	dbmock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow(1.0))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)
	fn := EqualString("field", "a")
	result := fn(gormDB)
	require.NotNil(t, result)
}
