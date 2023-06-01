package database

import "github.com/jmoiron/sqlx"

type Store struct {
	Master *sqlx.DB
	Slave  *sqlx.DB
}

func (s *Store) GetSlave() *sqlx.DB {
	return s.Slave
}
func (s *Store) GetMaster() *sqlx.DB {
	return s.Master
}
func (s *Store) GetExecContext(dbtx DBTX) DBTX {
	if dbtx == nil {
		return s.Master
	}
	return dbtx
}
func (s *Store) GetSelectContext(dbtx DBTX) DBTX {
	if dbtx == nil {
		return s.Slave
	}
	return dbtx
}
