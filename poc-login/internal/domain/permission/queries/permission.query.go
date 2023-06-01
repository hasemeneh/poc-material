package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hasemeneh/poc-material/poc-login/internal/constants"
	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/hasemeneh/poc-material/poc-login/pkg/database"
	"github.com/hasemeneh/poc-material/poc-login/pkg/pagespecifier"
	"github.com/jmoiron/sqlx"
)

type resource struct {
	db *database.Store
}
type Resource interface {
	FetchPermissionByRoleID(ctx context.Context, RoleID int64) ([]*models.Permission, error)
	FetchValidPermissionByAuthID(ctx context.Context, authID string) ([]*models.Permission, error)
	InsertPermission(ctx context.Context, dbTx database.DBTX, m *models.Permission) (int64, error)
	UpdatePermission(ctx context.Context, dbTx database.DBTX, m *models.Permission) error
	FetchPermissionForInternal(ctx context.Context, rf pagespecifier.ResultFilter) ([]*models.Permission, int64, error)
	InsertRolePermission(ctx context.Context, dbTx database.DBTX, m *models.InternalRolePermission) (int64, error)
	ChangeStatusRolePermission(ctx context.Context, dbTx database.DBTX, status constants.RolePermissionStatus, RoleID, PermissionID int64) error
	StartTransaction(ctx context.Context) (*sqlx.Tx, error)
}

func NewQuery(db *database.Store) Resource {
	return &resource{
		db: db,
	}
}

func (r *resource) StartTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return r.db.GetMaster().BeginTxx(ctx, nil)

}

const selectValidPermissionByAuthID = `
SELECT
	p.code,
	p.name,
	r.id as role_id
FROM rbac_permission p
INNER JOIN rbac_role_permission  rp ON p.id = rp.permission_id AND p.status = ? AND rp.status = ?
INNER JOIN rbac_role r ON rp.role_id = r.id AND r.status = ?
INNER JOIN rbac_user_access ua ON ua.role_id = r.id AND ua.status = ?
INNER JOIN rbac_user u ON u.id = ua.user_id AND u.status = ?
WHERE u.auth_id = ? 
`

func (r *resource) FetchValidPermissionByAuthID(ctx context.Context, authID string) ([]*models.Permission, error) {
	resp := make([]*models.Permission, 0)
	err := r.db.GetSlave().SelectContext(ctx, &resp, selectValidPermissionByAuthID,
		constants.PermissionStatusActive,
		constants.RolePermissionStatusActive,
		constants.RoleStatusActive,
		constants.UserAccessStatusActive,
		constants.UserStatusActive,
		authID)
	return resp, err
}

const fetchInternalQuery = `
SELECT
	id,
	code,
	status,
	name
FROM rbac_permission
`

func (u *resource) FetchPermissionForInternal(ctx context.Context, rf pagespecifier.ResultFilter) ([]*models.Permission, int64, error) {
	qf := pagespecifier.GetQueryFilterWithParser(rf.Filter)
	args := qf.Args
	whereClause := qf.GetWhereClause()
	orderClause := pagespecifier.GetQuerySortWithParser(rf.Sort, pagespecifier.DefaultQuerySortParser).GetOrderClause()

	var limitClause string
	if rf.Limit > 0 {
		var offset int64
		if rf.Page > 0 {
			offset = (rf.Page - 1) * rf.Limit
		}
		limitClause = fmt.Sprintf("LIMIT %d OFFSET %d", rf.Limit, offset)
	}
	query := fmt.Sprintf(`%s %s %s %s`, fetchInternalQuery, whereClause, orderClause, limitClause)
	p := make([]*models.Permission, 0)
	var total int64

	dbTx, err := u.db.GetSlave().BeginTxx(ctx, nil)
	if err != nil {
		return p, total, err
	}
	defer dbTx.Rollback()

	err = dbTx.SelectContext(ctx, &p, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, 0, nil
		}
		return p, 0, err
	}
	row := dbTx.QueryRow("SELECT FOUND_ROWS()")
	err = row.Scan(&total)
	if err != nil {
		return p, 0, err
	}

	return p, total, nil
}

const selectPermissionByRoleID = `
SELECT
	p.code,
	p.name,
	p.status
FROM rbac_permission p
INNER JOIN rbac_role_permission  rp ON p.id = rp.permission_id AND  rp.role_id = ?
`

func (r *resource) FetchPermissionByRoleID(ctx context.Context, RoleID int64) ([]*models.Permission, error) {
	resp := make([]*models.Permission, 0)
	err := r.db.GetSlave().SelectContext(ctx, &resp, selectPermissionByRoleID,
		RoleID)
	return resp, err
}
