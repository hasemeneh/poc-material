package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/constants"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/database"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/pagespecifier"
)

type resource struct {
	db *database.Store
}
type Resource interface {
	FetchValidRoleByAuthID(ctx context.Context, authID string) ([]*models.Role, error)
	InsertRole(ctx context.Context, dbTx database.DBTX, m *models.InternalRole) (int64, error)
	UpdateRole(ctx context.Context, dbTx database.DBTX, m *models.InternalRole) error
	FetchRoleForInternal(ctx context.Context, rf pagespecifier.ResultFilter) ([]*models.InternalRole, int64, error)
	FetchUserAccessByIDInternal(ctx context.Context, userID int64) ([]*models.UserAccess, error)
	InsertUserAccess(ctx context.Context, dbTx database.DBTX, m *models.UserAccess) (int64, error)
	ChangeStatusUserAccess(ctx context.Context, dbTx database.DBTX, status constants.UserAccessStatus, RoleID, UserID int64) error
}

func NewQuery(db *database.Store) Resource {
	return &resource{
		db: db,
	}
}

const fetchRoleByAuthIDQuery = `
SELECT
	r.id,
	r.name
FROM rbac_role r
INNER JOIN rbac_user_access ua ON ua.role_id = r.id AND r.status = ? AND ua.status = ?
INNER JOIN rbac_user u ON u.id = ua.user_id AND u.status = ?
WHERE u.auth_id = ?
`

func (r *resource) FetchValidRoleByAuthID(ctx context.Context, authID string) ([]*models.Role, error) {
	resp := make([]*models.Role, 0)
	err := r.db.GetSlave().SelectContext(ctx, &resp, fetchRoleByAuthIDQuery,
		constants.RoleStatusActive,
		constants.UserAccessStatusActive,
		constants.UserStatusActive,
		authID)
	return resp, err
}

const fetchInternalQuery = `
SELECT
	id,
	status,
	name
FROM rbac_role
`

func (u *resource) FetchRoleForInternal(ctx context.Context, rf pagespecifier.ResultFilter) ([]*models.InternalRole, int64, error) {
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
	p := make([]*models.InternalRole, 0)
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
